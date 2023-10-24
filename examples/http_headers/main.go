// Copyright 2020-2021 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// There are four types of these interfaces which you are supposed to implement in order to extend your network proxies.
// They are VMContext, PluginContext, TcpContext and HttpContext, and their relationship can be described as the following diagram:
//
//	                        Wasm Virtual Machine(VM)
//	                   (corresponds to VM configuration)
//	┌────────────────────────────────────────────────────────────────────────────┐
//	│                                                      TcpContext            │
//	│                                                  ╱ (Each Tcp stream)       │
//	│                                                 ╱                          │
//	│                      1: N                      ╱ 1: N                      │
//	│       VMContext  ──────────  PluginContext                                 │
//	│  (VM configuration)     (Plugin configuration) ╲ 1: N                      │
//	│                                                 ╲                          │
//	│                                                  ╲   HttpContext           │
//	│                                                   (Each Http stream)       │
//	└────────────────────────────────────────────────────────────────────────────┘
//
package main

import (
	"bytes"
	"fmt"
	"github.com/tidwall/gjson"
	"net/http"
	"strings"
	"time"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

const (
	upstreamURL = "http://example.com"
)

type vmContext struct {
	// Embed the default VM context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultVMContext
}

// Override types.DefaultVMContext.
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext

	// headerName and headerValue are the header to be added to response. They are configured via
	// plugin configuration during OnPluginStart.
	headerName  string
	headerValue string
}

// NewHttpContext используется для создания HttpContext для каждого Http-потока.
// Возвращаем nil, чтобы указать, что этот PluginContext не для HttpContext.
func (p *pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpHeaders{
		contextID:   contextID,
		headerName:  p.headerName,
		headerValue: p.headerValue,
	}
}

func (p *pluginContext) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {
	proxywasm.LogDebug("loading plugin config")
	data, err := proxywasm.GetPluginConfiguration()
	if data == nil {
		return types.OnPluginStartStatusOK
	}

	if err != nil {
		proxywasm.LogCriticalf("error reading plugin configuration: %v", err)
		return types.OnPluginStartStatusFailed
	}

	if !gjson.Valid(string(data)) {
		proxywasm.LogCritical(`invalid configuration format; expected {"header": "<header name>", "value": "<header value>"}`)
		return types.OnPluginStartStatusFailed
	}

	p.headerName = strings.TrimSpace(gjson.Get(string(data), "header").Str)
	p.headerValue = strings.TrimSpace(gjson.Get(string(data), "value").Str)

	if p.headerName == "" || p.headerValue == "" {
		proxywasm.LogCritical(`invalid configuration format; expected {"header": "<header name>", "value": "<header value>"}`)
		return types.OnPluginStartStatusFailed
	}

	proxywasm.LogInfof("header from config: %s = %s", p.headerName, p.headerValue)

	return types.OnPluginStartStatusOK
}

type httpHeaders struct {
	// Вставьте здесь контекст http по умолчанию, так что нам не нужно переопределять все методы.
	types.DefaultHttpContext
	contextID   uint32
	headerName  string
	headerValue string
}

// OnHttpRequestHeaders вызывается, когда приходят заголовки запроса.
// Возвращаем types.ActionPause, если вы хотите прекратить отправку заголовков исходящему потоку.
func (ctx *httpHeaders) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	//Можно добавлять свои заголовки
	return types.ActionContinue
}

// OnHttpResponseHeaders вызывается, когда приходят заголовки ответа.
// Возврат типов.ActionPause, если вы хотите остановить отправку заголовков во входящий поток.
func (ctx *httpHeaders) OnHttpResponseHeaders(_ int, _ bool) types.Action {
	return types.ActionContinue
}

// OnHttpStreamDone вызывается до того, как хост удалит этот контекст.
// Вы можете получить информацию HTTP-запроса/ответа (например, заголовки и т. д.) во время этого вызова.
// Это можно использовать для реализации функций ведения журнала.
func (ctx *httpHeaders) OnHttpStreamDone() {

	requestHeaders, err := proxywasm.GetHttpRequestHeaders()
	if err != nil {
		//proxywasm.LogCriticalf("failed to get request headers: %v", err)
	} else {
		for _, h := range requestHeaders {
			proxywasm.LogInfof("request header --> %s: %s", h[0], h[1])
		}
	}

	// Add the header passed by arguments
	if ctx.headerName != "" {
		if err := proxywasm.AddHttpResponseHeader(ctx.headerName, ctx.headerValue); err != nil {
			proxywasm.LogCriticalf("failed to set response headers: %v", err)
		}
	}

	// Get and log the headers
	responseHeaders, err := proxywasm.GetHttpResponseHeaders()
	if err != nil {
		proxywasm.LogCriticalf("failed to get response headers: %v", err)
	}

	jwtToken := getEventByResponse(requestHeaders, "authorization")
	clientCert := getEventByResponse(requestHeaders, "x-forwarded-client-cert")

	// Обработка статуса в ответе
	for _, h := range responseHeaders {
		if h[0] != ":status" {
			proxywasm.LogInfof("response header <-- %s: %s", h[0], h[1])
			continue
		}

		//Некорректный jwt-токен
		if h[1] == "401" {
			proxywasm.LogInfof("Audit event AUTHENTICATION_JWT_FAIL: %s", getEventByResponse(responseHeaders, "www-authenticate"))
		}

		//Авторизация не пройдена
		if h[1] == "403" {
			if jwtToken != "" {
				proxywasm.LogInfof("Audit event AUTHORIZATION_JWT_FAIL: %s", jwtToken)
			} else {
				proxywasm.LogInfof("Audit event AUTHORIZATION_CERT_FAIL: %s", clientCert)
			}
		}

		if h[1] == "200" || h[1] == "201" {
			if jwtToken != "" {
				proxywasm.LogInfof("Audit event AUTHORIZATION_JWT_SUCCESS: %s", jwtToken)
			} else {
				proxywasm.LogInfof("Audit event AUTHORIZATION_CERT_SUCCESS: %s", clientCert)
			}
		}

		// Create an HTTP client with a timeout
		client := http.Client{
			Timeout: time.Second * 5,
		}

		//rquid := uuid.New().String()

		// JSON body
		body := []byte("rquid")

		// Create a HTTP post request
		resp, err := client.Post(upstreamURL, "application/json", bytes.NewBuffer(body))
		if err != nil {
			proxywasm.LogCriticalf("Failed to send HTTP request:", err)
		}

		defer resp.Body.Close()

		// Process the response
		fmt.Println("Received HTTP response with status code:", resp.StatusCode)

		//handler := watchers.FakeAccessLogHandler{}
		// watcher := watchers.NewWatcherV2(&handler)

		// proxywasm.LogInfof("Audit event watcher: %s", watcher)

		//ctx, _ := context.WithCancel(context.Background())
		//stream, err := client.StreamAccessLogs(ctx)

	}
}

func getEventByResponse(headers [][2]string, key string) string {
	for _, h := range headers {
		if h[0] == key {
			return h[1]
		}
	}
	return ""
}
