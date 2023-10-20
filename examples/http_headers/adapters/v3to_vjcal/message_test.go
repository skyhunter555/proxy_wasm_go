package v3to_vjcal

import (
    "testing"
    "time"

    envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
    v3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
    envoy_service_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
    "github.com/golang/protobuf/ptypes/any"
    "github.com/stretchr/testify/assert"
    "google.golang.org/protobuf/types/known/durationpb"
    "google.golang.org/protobuf/types/known/timestamppb"
    "google.golang.org/protobuf/types/known/wrapperspb"

    "bitbucket.region.vtb.ru/projects/USBP/repos/envoy-accesslogs/internal/vtbjournalclient_accesslogs"
)

//
// Helpers
//

func getPipeAddressFor(path string) *envoy_config_core_v3.Address {
    return &envoy_config_core_v3.Address{
        Address: &envoy_config_core_v3.Address_Pipe{
            Pipe: &envoy_config_core_v3.Pipe{
                Path: path,
                Mode: 0,
            },
        },
    }
}

//
//
//

func Test_adapt_Message(t *testing.T) {
    type args struct {
        src *envoy_service_accesslog_v3.StreamAccessLogsMessage
    }
    tests := []struct {
        name string
        args args
        want []*vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog
    }{
        {
            name: "Test NIL",
            args: args{
                src: nil,
            },
            want: []*vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog{},
        },
        {
            name: "Test ARRAY HTTP",
            args: args{
                src: &envoy_service_accesslog_v3.StreamAccessLogsMessage{
                    Identifier: &envoy_service_accesslog_v3.StreamAccessLogsMessage_Identifier{},
                    LogEntries: &envoy_service_accesslog_v3.StreamAccessLogsMessage_HttpLogs{
                        HttpLogs: &envoy_service_accesslog_v3.StreamAccessLogsMessage_HTTPAccessLogEntries{
                            LogEntry: []*v3.HTTPAccessLogEntry{
                                {},
                                {},
                            },
                        },
                    },
                },
            },
            want: []*vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog{
                {Protocol: "PROTOCOL_UNSPECIFIED"},
                {Protocol: "PROTOCOL_UNSPECIFIED"},
            },
        },
        {
            name: "Test ARRAY TCP",
            args: args{
                src: &envoy_service_accesslog_v3.StreamAccessLogsMessage{
                    Identifier: &envoy_service_accesslog_v3.StreamAccessLogsMessage_Identifier{},
                    LogEntries: &envoy_service_accesslog_v3.StreamAccessLogsMessage_TcpLogs{
                        TcpLogs: &envoy_service_accesslog_v3.StreamAccessLogsMessage_TCPAccessLogEntries{
                            LogEntry: []*v3.TCPAccessLogEntry{
                                {},
                                {},
                            },
                        },
                    },
                },
            },
            want: []*vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog{
                {Protocol: "TCP"},
                {Protocol: "TCP"},
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

func Test_adapt_HTTP(t *testing.T) {
    type args struct {
        src *v3.HTTPAccessLogEntry
    }
    tests := []struct {
        name string
        args args
        want *vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog
    }{
        {
            name: "Test NIL",
            args: args{
                src: nil,
            },
            want: nil,
        },
        {
            name: "Test OBJECT",
            args: args{
                src: &v3.HTTPAccessLogEntry{
                    CommonProperties: &v3.AccessLogCommon{
                        SampleRate:                     1,
                        DownstreamRemoteAddress:        getPipeAddressFor("2"),
                        DownstreamLocalAddress:         getPipeAddressFor("3"),
                        TlsProperties:                  &v3.TLSProperties{TlsSniHostname: "4"},
                        StartTime:                      timestamppb.New(time.Date(1, time.January, 1, 1, 1, 5, 0, time.UTC)),
                        TimeToLastRxByte:               durationpb.New(time.Second * 6),
                        TimeToFirstUpstreamTxByte:      durationpb.New(time.Second * 7),
                        TimeToLastUpstreamTxByte:       durationpb.New(time.Second * 8),
                        TimeToFirstUpstreamRxByte:      durationpb.New(time.Second * 9),
                        TimeToLastUpstreamRxByte:       durationpb.New(time.Second * 10),
                        TimeToFirstDownstreamTxByte:    durationpb.New(time.Second * 11),
                        TimeToLastDownstreamTxByte:     durationpb.New(time.Second * 12),
                        UpstreamRemoteAddress:          getPipeAddressFor("13"),
                        UpstreamLocalAddress:           getPipeAddressFor("14"),
                        UpstreamCluster:                "15",
                        ResponseFlags:                  &v3.ResponseFlags{FailedLocalHealthcheck: true},
                        Metadata:                       nil,
                        UpstreamTransportFailureReason: "18",
                        RouteName:                      "19",
                        DownstreamDirectRemoteAddress:  getPipeAddressFor("20"),
                        FilterStateObjects:             map[string]*any.Any{"21.k": {}},
                    },
                    ProtocolVersion: v3.HTTPAccessLogEntry_HTTP3,
                    Request: &v3.HTTPRequestProperties{
                        RequestMethod:       envoy_config_core_v3.RequestMethod_GET,
                        Scheme:              "23",
                        Authority:           "24",
                        Port:                wrapperspb.UInt32(25),
                        Path:                "26",
                        UserAgent:           "27",
                        Referer:             "28",
                        ForwardedFor:        "29",
                        RequestId:           "30",
                        OriginalPath:        "31",
                        RequestHeadersBytes: 32,
                        RequestBodyBytes:    33,
                        RequestHeaders: map[string]string{
                            "34.1.k":                  "34.1.v",
                            "x-b3-traceid":            "34.2.v",
                            "x-forwarded-for":         "34.3.v",
                            "x-forwarded-client-cert": "34.4.v",
                            "x-ratelimit-limit":       "34.5.v",
                            "x-ratelimit-remaining":   "34.6.v",
                            "x-ratelimit-reset":       "34.7.v",
                            "host":                    "34.8.v",
                        },
                    },
                    Response: &v3.HTTPResponseProperties{
                        ResponseCode:         wrapperspb.UInt32(35),
                        ResponseHeadersBytes: 36,
                        ResponseBodyBytes:    37,
                        ResponseHeaders:      map[string]string{"38.k": "38.v"},
                        ResponseTrailers:     map[string]string{"38.k": "38.v"},
                        ResponseCodeDetails:  "39",
                    },
                },
            },
            want: &vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog{
                TraceId:                        "34.2.v",
                Protocol:                       "HTTP3",
                UpstreamLocalAddress:           "unix://14",
                Duration:                       time.Second * 12,
                UpstreamTransportFailureReason: "18",
                DownstreamLocalAddress:         "unix://3",
                UserAgent:                      "27",
                ResponseCode:                   35,
                ResponseFlags:                  "LH",
                StartTime:                      time.Date(1, time.January, 1, 1, 1, 5, 0, time.UTC),
                Method:                         "GET",
                RequestId:                      "30",
                UpstreamHost:                   "unix://13",
                XForwardedFor:                  "34.3.v",
                XForwardedClientCert:           "34.4.v",
                XRatelimitLimit:                "34.5.v",
                XRatelimitRemaining:            "34.6.v",
                XRatelimitReset:                "34.7.v",
                RequestedServerName:            "4",
                BytesReceived:                  33,
                BytesSent:                      37,
                UpstreamCluster:                "15",
                DownstreamRemoteAddress:        "unix://2",
                Authority:                      "24",
                Host:                           "34.8.v",
                Path:                           "31",
            },
        },
        {
            name: "Test OBJECT. Path exists: OriginPath, x-envoy-origin-path, PATH. OriginPath should be chosen",
            args: args{
                src: &v3.HTTPAccessLogEntry{
                    CommonProperties: &v3.AccessLogCommon{},
                    ProtocolVersion:  v3.HTTPAccessLogEntry_HTTP3,
                    Request: &v3.HTTPRequestProperties{
                        Path:         "26",
                        OriginalPath: "31",
                        RequestHeaders: map[string]string{
                            "x-envoy-origin-path": "34.1.v",
                        },
                    },
                    Response: &v3.HTTPResponseProperties{},
                },
            },
            want: &vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog{
                StartTime: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
                Protocol:  "HTTP3",
                Method:    "METHOD_UNSPECIFIED",
                Path:      "31",
            },
        },
        {
            name: "Test OBJECT. Path exists: x-envoy-origin-path, PATH. x-envoy-origin-path should be chosen",
            args: args{
                src: &v3.HTTPAccessLogEntry{
                    CommonProperties: &v3.AccessLogCommon{},
                    ProtocolVersion:  v3.HTTPAccessLogEntry_HTTP3,
                    Request: &v3.HTTPRequestProperties{
                        Path: "26",
                        RequestHeaders: map[string]string{
                            "x-envoy-original-path": "34.1.v",
                        },
                    },
                    Response: &v3.HTTPResponseProperties{},
                },
            },
            want: &vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog{
                StartTime: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
                Protocol:  "HTTP3",
                Method:    "METHOD_UNSPECIFIED",
                Path:      "34.1.v",
            },
        },
        {
            name: "Test OBJECT. Path exists: PATH. PATH should be chosen",
            args: args{
                src: &v3.HTTPAccessLogEntry{
                    CommonProperties: &v3.AccessLogCommon{},
                    ProtocolVersion:  v3.HTTPAccessLogEntry_HTTP3,
                    Request: &v3.HTTPRequestProperties{
                        Path: "26",
                    },
                    Response: &v3.HTTPResponseProperties{},
                },
            },
            want: &vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog{
                StartTime: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
                Protocol:  "HTTP3",
                Method:    "METHOD_UNSPECIFIED",
                Path:      "26",
            },
        },
        {
            name: "Test OBJECT. No path exists",
            args: args{
                src: &v3.HTTPAccessLogEntry{
                    CommonProperties: &v3.AccessLogCommon{},
                    ProtocolVersion:  v3.HTTPAccessLogEntry_HTTP3,
                    Request:          &v3.HTTPRequestProperties{},
                    Response:         &v3.HTTPResponseProperties{},
                },
            },
            want: &vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog{
                StartTime: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
                Protocol:  "HTTP3",
                Method:    "METHOD_UNSPECIFIED",
                Path:      "",
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_HTTP(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_TCP(t *testing.T) {
    type args struct {
        src *v3.TCPAccessLogEntry
    }
    tests := []struct {
        name string
        args args
        want *vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog
    }{
        {
            name: "Test NIL",
            args: args{
                src: nil,
            },
            want: nil,
        },
        {
            name: "Test OBJECT",
            args: args{
                src: &v3.TCPAccessLogEntry{
                    CommonProperties: &v3.AccessLogCommon{
                        SampleRate:                     1,
                        DownstreamRemoteAddress:        getPipeAddressFor("2"),
                        DownstreamLocalAddress:         getPipeAddressFor("3"),
                        TlsProperties:                  &v3.TLSProperties{TlsSniHostname: "4"},
                        StartTime:                      timestamppb.New(time.Date(1, time.January, 1, 1, 1, 5, 0, time.UTC)),
                        TimeToLastRxByte:               durationpb.New(time.Second * 6),
                        TimeToFirstUpstreamTxByte:      durationpb.New(time.Second * 7),
                        TimeToLastUpstreamTxByte:       durationpb.New(time.Second * 8),
                        TimeToFirstUpstreamRxByte:      durationpb.New(time.Second * 9),
                        TimeToLastUpstreamRxByte:       durationpb.New(time.Second * 10),
                        TimeToFirstDownstreamTxByte:    durationpb.New(time.Second * 11),
                        TimeToLastDownstreamTxByte:     durationpb.New(time.Second * 12),
                        UpstreamRemoteAddress:          getPipeAddressFor("13"),
                        UpstreamLocalAddress:           getPipeAddressFor("14"),
                        UpstreamCluster:                "15",
                        ResponseFlags:                  &v3.ResponseFlags{FailedLocalHealthcheck: true},
                        Metadata:                       nil,
                        UpstreamTransportFailureReason: "18",
                        RouteName:                      "19",
                        DownstreamDirectRemoteAddress:  getPipeAddressFor("20"),
                        FilterStateObjects:             map[string]*any.Any{"21.k": {}},
                    },
                    ConnectionProperties: &v3.ConnectionProperties{
                        ReceivedBytes: 22,
                        SentBytes:     23,
                    },
                },
            },
            want: &vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog{
                TraceId:                        "",
                Protocol:                       "TCP",
                UpstreamLocalAddress:           "unix://14",
                Duration:                       time.Second * 12,
                UpstreamTransportFailureReason: "18",
                DownstreamLocalAddress:         "unix://3",
                UserAgent:                      "",
                ResponseCode:                   0,
                ResponseFlags:                  "LH",
                StartTime:                      time.Date(1, time.January, 1, 1, 1, 5, 0, time.UTC),
                Method:                         "",
                RequestId:                      "",
                UpstreamHost:                   "unix://13",
                XForwardedFor:                  "",
                XForwardedClientCert:           "",
                XRatelimitLimit:                "",
                XRatelimitRemaining:            "",
                XRatelimitReset:                "",
                RequestedServerName:            "4",
                BytesReceived:                  22,
                BytesSent:                      23,
                UpstreamCluster:                "15",
                DownstreamRemoteAddress:        "unix://2",
                Authority:                      "",
                Host:                           "",
                Path:                           "",
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_TCP(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_Address(t *testing.T) {
    type args struct {
        src *envoy_config_core_v3.Address
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "Test NIL",
            args: args{
                src: nil,
            },
            want: "",
        },
        {
            name: "Test NIL",
            args: args{
                src: &envoy_config_core_v3.Address{},
            },
            want: "",
        },
        {
            name: "Test OBJECT Socket TCP PortValue",
            args: args{
                src: &envoy_config_core_v3.Address{
                    Address: &envoy_config_core_v3.Address_SocketAddress{
                        SocketAddress: &envoy_config_core_v3.SocketAddress{
                            Protocol:      envoy_config_core_v3.SocketAddress_TCP,
                            Address:       "asd",
                            PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{PortValue: 123},
                            ResolverName:  "qwe",
                            Ipv4Compat:    true,
                        },
                    },
                },
            },
            want: "tcp://asd:123",
        },
        {
            name: "Test OBJECT Socket UDP NamedPort",
            args: args{
                src: &envoy_config_core_v3.Address{
                    Address: &envoy_config_core_v3.Address_SocketAddress{
                        SocketAddress: &envoy_config_core_v3.SocketAddress{
                            Protocol:      envoy_config_core_v3.SocketAddress_UDP,
                            Address:       "asd",
                            PortSpecifier: &envoy_config_core_v3.SocketAddress_NamedPort{NamedPort: "zxc"},
                            ResolverName:  "qwe",
                            Ipv4Compat:    true,
                        },
                    },
                },
            },
            want: "udp://asd:[zxc]",
        },
        {
            name: "Test OBJECT EnvoyListenerAddress ServerListenerName",
            args: args{
                src: &envoy_config_core_v3.Address{
                    Address: &envoy_config_core_v3.Address_EnvoyInternalAddress{
                        EnvoyInternalAddress: &envoy_config_core_v3.EnvoyInternalAddress{
                            AddressNameSpecifier: &envoy_config_core_v3.EnvoyInternalAddress_ServerListenerName{
                                ServerListenerName: "qwe123",
                            },
                        },
                    },
                },
            },
            want: "qwe123",
        },
        {
            name: "Test OBJECT Pipe",
            args: args{
                src: &envoy_config_core_v3.Address{
                    Address: &envoy_config_core_v3.Address_Pipe{
                        Pipe: &envoy_config_core_v3.Pipe{
                            Path: "/path/to/file",
                            Mode: 0,
                        },
                    },
                },
            },
            want: "unix:///path/to/file",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_Address(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}
