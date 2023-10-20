package v2tov3

import (
    envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
    envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
    envoy_data_accesslog_src "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v2"
    envoy_data_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
    "github.com/golang/protobuf/ptypes/wrappers"
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_adapt_HTTPAccessLogEntries(t *testing.T) {
    type args struct {
        src []*envoy_data_accesslog_src.HTTPAccessLogEntry
    }
    tests := []struct {
        name string
        args args
        want []*envoy_data_accesslog_v3.HTTPAccessLogEntry
    }{
        {
            name: "Convert NIL",
            args: args{
                src: nil,
            },
            want: nil,
        },
        {
            name: "Convert OBJECT",
            args: args{
                src: []*envoy_data_accesslog_src.HTTPAccessLogEntry{
                    {},
                    {},
                },
            },
            want: []*envoy_data_accesslog_v3.HTTPAccessLogEntry{
                {},
                {},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_HTTPAccessLogEntries(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_HTTPAccessLogEntry(t *testing.T) {
    type args struct {
        src *envoy_data_accesslog_src.HTTPAccessLogEntry
    }
    tests := []struct {
        name string
        args args
        want *envoy_data_accesslog_v3.HTTPAccessLogEntry
    }{
        {
            name: "Convert NIL",
            args: args{
                src: nil,
            },
            want: nil,
        },
        {
            name: "Convert OBJECT",
            args: args{
                src: &envoy_data_accesslog_src.HTTPAccessLogEntry{
                    CommonProperties: &envoy_data_accesslog_src.AccessLogCommon{},
                    ProtocolVersion:  envoy_data_accesslog_src.HTTPAccessLogEntry_HTTP10,
                    Request:          &envoy_data_accesslog_src.HTTPRequestProperties{},
                    Response:         &envoy_data_accesslog_src.HTTPResponseProperties{},
                },
            },
            want: &envoy_data_accesslog_v3.HTTPAccessLogEntry{
                CommonProperties: &envoy_data_accesslog_v3.AccessLogCommon{},
                ProtocolVersion:  envoy_data_accesslog_v3.HTTPAccessLogEntry_HTTP10,
                Request:          &envoy_data_accesslog_v3.HTTPRequestProperties{},
                Response:         &envoy_data_accesslog_v3.HTTPResponseProperties{},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_HTTPAccessLogEntry(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_HTTPRequestProperties(t *testing.T) {
    type args struct {
        src *envoy_data_accesslog_src.HTTPRequestProperties
    }
    tests := []struct {
        name string
        args args
        want *envoy_data_accesslog_v3.HTTPRequestProperties
    }{
        {
            name: "Convert NIL",
            args: args{
                src: nil,
            },
            want: nil,
        },
        {
            name: "Convert OBJECT",
            args: args{
                src: &envoy_data_accesslog_src.HTTPRequestProperties{
                    RequestMethod:       envoy_api_v2_core.RequestMethod_GET,
                    Scheme:              "2",
                    Authority:           "3",
                    Port:                &wrappers.UInt32Value{Value: 4},
                    Path:                "5",
                    UserAgent:           "6",
                    Referer:             "7",
                    ForwardedFor:        "8",
                    RequestId:           "9",
                    OriginalPath:        "10",
                    RequestHeadersBytes: 11,
                    RequestBodyBytes:    12,
                    RequestHeaders:      map[string]string{"13.k": "13.v"},
                },
            },
            want: &envoy_data_accesslog_v3.HTTPRequestProperties{
                RequestMethod:       envoy_config_core_v3.RequestMethod_GET,
                Scheme:              "2",
                Authority:           "3",
                Port:                &wrappers.UInt32Value{Value: 4},
                Path:                "5",
                UserAgent:           "6",
                Referer:             "7",
                ForwardedFor:        "8",
                RequestId:           "9",
                OriginalPath:        "10",
                RequestHeadersBytes: 11,
                RequestBodyBytes:    12,
                RequestHeaders:      map[string]string{"13.k": "13.v"},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_HTTPRequestProperties(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_HTTPResponseProperties(t *testing.T) {
    type args struct {
        src *envoy_data_accesslog_src.HTTPResponseProperties
    }
    tests := []struct {
        name string
        args args
        want *envoy_data_accesslog_v3.HTTPResponseProperties
    }{
        {
            name: "Convert NIL",
            args: args{
                src: nil,
            },
            want: nil,
        },
        {
            name: "Convert OBJECT",
            args: args{
                src: &envoy_data_accesslog_src.HTTPResponseProperties{
                    ResponseCode:         &wrappers.UInt32Value{Value: 1},
                    ResponseHeadersBytes: 2,
                    ResponseBodyBytes:    3,
                    ResponseHeaders:      map[string]string{"3.k": "3.v"},
                    ResponseTrailers:     map[string]string{"4.k": "4.v"},
                    ResponseCodeDetails:  "5",
                },
            },
            want: &envoy_data_accesslog_v3.HTTPResponseProperties{
                ResponseCode:         &wrappers.UInt32Value{Value: 1},
                ResponseHeadersBytes: 2,
                ResponseBodyBytes:    3,
                ResponseHeaders:      map[string]string{"3.k": "3.v"},
                ResponseTrailers:     map[string]string{"4.k": "4.v"},
                ResponseCodeDetails:  "5",
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_HTTPResponseProperties(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}
