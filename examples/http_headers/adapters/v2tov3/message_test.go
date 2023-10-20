package v2tov3

import (
    envoy_data_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v2"
    envoy_data_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
    envoy_service_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
    envoy_service_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_adapt_Message(t *testing.T) {
    type args struct {
        src *envoy_service_accesslog_v2.StreamAccessLogsMessage
    }
    tests := []struct {
        name string
        args args
        want *envoy_service_accesslog_v3.StreamAccessLogsMessage
    }{
        {
            name: "Convert NIL",
            args: args{
                src: nil,
            },
            want: nil,
        },
        {
            name: "Convert OBJECT HTTP LogEntries",
            args: args{
                src: &envoy_service_accesslog_v2.StreamAccessLogsMessage{
                    Identifier: &envoy_service_accesslog_v2.StreamAccessLogsMessage_Identifier{},
                    LogEntries: &envoy_service_accesslog_v2.StreamAccessLogsMessage_HttpLogs{
                        HttpLogs: &envoy_service_accesslog_v2.StreamAccessLogsMessage_HTTPAccessLogEntries{
                            LogEntry: []*envoy_data_accesslog_v2.HTTPAccessLogEntry{
                                {},
                                {},
                            },
                        },
                    },
                },
            },
            want: &envoy_service_accesslog_v3.StreamAccessLogsMessage{
                Identifier: &envoy_service_accesslog_v3.StreamAccessLogsMessage_Identifier{},
                LogEntries: &envoy_service_accesslog_v3.StreamAccessLogsMessage_HttpLogs{
                    HttpLogs: &envoy_service_accesslog_v3.StreamAccessLogsMessage_HTTPAccessLogEntries{
                        LogEntry: []*envoy_data_accesslog_v3.HTTPAccessLogEntry{
                            {},
                            {},
                        },
                    },
                },
            },
        },
        {
            name: "Convert OBJECT TCP LogEntries",
            args: args{
                src: &envoy_service_accesslog_v2.StreamAccessLogsMessage{
                    Identifier: &envoy_service_accesslog_v2.StreamAccessLogsMessage_Identifier{},
                    LogEntries: &envoy_service_accesslog_v2.StreamAccessLogsMessage_TcpLogs{
                        TcpLogs: &envoy_service_accesslog_v2.StreamAccessLogsMessage_TCPAccessLogEntries{
                            LogEntry: []*envoy_data_accesslog_v2.TCPAccessLogEntry{
                                {},
                                {},
                            },
                        },
                    },
                },
            },
            want: &envoy_service_accesslog_v3.StreamAccessLogsMessage{
                Identifier: &envoy_service_accesslog_v3.StreamAccessLogsMessage_Identifier{},
                LogEntries: &envoy_service_accesslog_v3.StreamAccessLogsMessage_TcpLogs{
                    TcpLogs: &envoy_service_accesslog_v3.StreamAccessLogsMessage_TCPAccessLogEntries{
                        LogEntry: []*envoy_data_accesslog_v3.TCPAccessLogEntry{
                            {},
                            {},
                        },
                    },
                },
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_Message(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}
