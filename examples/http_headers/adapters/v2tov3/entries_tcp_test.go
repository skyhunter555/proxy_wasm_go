package v2tov3

import (
    envoy_data_accesslog_src "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v2"
    envoy_data_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_adapt_TcpAccessLogEntries(t *testing.T) {
    type args struct {
        src []*envoy_data_accesslog_src.TCPAccessLogEntry
    }
    tests := []struct {
        name string
        args args
        want []*envoy_data_accesslog_v3.TCPAccessLogEntry
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
                src: []*envoy_data_accesslog_src.TCPAccessLogEntry{
                    {},
                    {},
                },
            },
            want: []*envoy_data_accesslog_v3.TCPAccessLogEntry{
                {},
                {},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_TcpAccessLogEntries(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_TcpAccessLogEntry(t *testing.T) {
    type args struct {
        src *envoy_data_accesslog_src.TCPAccessLogEntry
    }
    tests := []struct {
        name string
        args args
        want *envoy_data_accesslog_v3.TCPAccessLogEntry
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
                src: &envoy_data_accesslog_src.TCPAccessLogEntry{
                    CommonProperties:     &envoy_data_accesslog_src.AccessLogCommon{},
                    ConnectionProperties: &envoy_data_accesslog_src.ConnectionProperties{},
                },
            },
            want: &envoy_data_accesslog_v3.TCPAccessLogEntry{
                CommonProperties:     &envoy_data_accesslog_v3.AccessLogCommon{},
                ConnectionProperties: &envoy_data_accesslog_v3.ConnectionProperties{},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_TcpAccessLogEntry(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_ConnectionProperties(t *testing.T) {
    type args struct {
        src *envoy_data_accesslog_src.ConnectionProperties
    }
    tests := []struct {
        name string
        args args
        want *envoy_data_accesslog_v3.ConnectionProperties
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
                src: &envoy_data_accesslog_src.ConnectionProperties{
                    ReceivedBytes: 1,
                    SentBytes:     2,
                },
            },
            want: &envoy_data_accesslog_v3.ConnectionProperties{
                ReceivedBytes: 1,
                SentBytes:     2,
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_ConnectionProperties(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}
