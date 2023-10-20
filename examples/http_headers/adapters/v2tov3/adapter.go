package v2tov3

import (
    envoy_service_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
    envoy_service_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
)

//
//
//

// Adapter
type Adapter interface {
    Adapt(src *envoy_service_accesslog_v2.StreamAccessLogsMessage) *envoy_service_accesslog_v3.StreamAccessLogsMessage
}

// NewAdapter
func NewAdapter() Adapter {
    return &adapterImpl{}
}

//
//
//

// AdapterImpl
type adapterImpl struct {
}

func (a *adapterImpl) Adapt(src *envoy_service_accesslog_v2.StreamAccessLogsMessage) *envoy_service_accesslog_v3.StreamAccessLogsMessage {
    return adapt_Message(src)
}
