package v2tov3

import (
    envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
    envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
    envoy_data_accesslog_src "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v2"
    envoy_data_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
    "github.com/golang/protobuf/ptypes/any"
    structpb "github.com/golang/protobuf/ptypes/struct"
    "github.com/golang/protobuf/ptypes/wrappers"
    "github.com/stretchr/testify/assert"
    "google.golang.org/protobuf/types/known/durationpb"
    "google.golang.org/protobuf/types/known/timestamppb"
    "testing"
    "time"
)

func Test_adapt_AccessLogCommon(t *testing.T) {
    type args struct {
        src *envoy_data_accesslog_src.AccessLogCommon
    }
    tests := []struct {
        name string
        args args
        want *envoy_data_accesslog_v3.AccessLogCommon
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
                src: &envoy_data_accesslog_src.AccessLogCommon{
                    SampleRate:                     1,
                    DownstreamRemoteAddress:        &envoy_api_v2_core.Address{},
                    DownstreamLocalAddress:         &envoy_api_v2_core.Address{},
                    TlsProperties:                  &envoy_data_accesslog_src.TLSProperties{},
                    StartTime:                      timestamppb.New(time.Date(1, time.January, 1, 1, 1, 1, 5, time.UTC)),
                    TimeToLastRxByte:               durationpb.New(time.Second * 6),
                    TimeToFirstUpstreamTxByte:      durationpb.New(time.Second * 7),
                    TimeToLastUpstreamTxByte:       durationpb.New(time.Second * 8),
                    TimeToFirstUpstreamRxByte:      durationpb.New(time.Second * 9),
                    TimeToLastUpstreamRxByte:       durationpb.New(time.Second * 10),
                    TimeToFirstDownstreamTxByte:    durationpb.New(time.Second * 11),
                    TimeToLastDownstreamTxByte:     durationpb.New(time.Second * 12),
                    UpstreamRemoteAddress:          &envoy_api_v2_core.Address{},
                    UpstreamLocalAddress:           &envoy_api_v2_core.Address{},
                    UpstreamCluster:                "15",
                    ResponseFlags:                  &envoy_data_accesslog_src.ResponseFlags{},
                    Metadata:                       nil,
                    UpstreamTransportFailureReason: "18",
                    RouteName:                      "19",
                    DownstreamDirectRemoteAddress:  &envoy_api_v2_core.Address{},
                    FilterStateObjects:             map[string]*any.Any{},
                },
            },
            want: &envoy_data_accesslog_v3.AccessLogCommon{
                SampleRate:                     1,
                DownstreamRemoteAddress:        &envoy_config_core_v3.Address{},
                DownstreamLocalAddress:         &envoy_config_core_v3.Address{},
                TlsProperties:                  &envoy_data_accesslog_v3.TLSProperties{},
                StartTime:                      timestamppb.New(time.Date(1, time.January, 1, 1, 1, 1, 5, time.UTC)),
                TimeToLastRxByte:               durationpb.New(time.Second * 6),
                TimeToFirstUpstreamTxByte:      durationpb.New(time.Second * 7),
                TimeToLastUpstreamTxByte:       durationpb.New(time.Second * 8),
                TimeToFirstUpstreamRxByte:      durationpb.New(time.Second * 9),
                TimeToLastUpstreamRxByte:       durationpb.New(time.Second * 10),
                TimeToFirstDownstreamTxByte:    durationpb.New(time.Second * 11),
                TimeToLastDownstreamTxByte:     durationpb.New(time.Second * 12),
                UpstreamRemoteAddress:          &envoy_config_core_v3.Address{},
                UpstreamLocalAddress:           &envoy_config_core_v3.Address{},
                UpstreamCluster:                "15",
                ResponseFlags:                  &envoy_data_accesslog_v3.ResponseFlags{},
                Metadata:                       nil,
                UpstreamTransportFailureReason: "18",
                RouteName:                      "19",
                DownstreamDirectRemoteAddress:  &envoy_config_core_v3.Address{},
                FilterStateObjects:             map[string]*any.Any{},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_AccessLogCommon(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_TLSProperties(t *testing.T) {
    type args struct {
        src *envoy_data_accesslog_src.TLSProperties
    }
    tests := []struct {
        name string
        args args
        want *envoy_data_accesslog_v3.TLSProperties
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
                src: &envoy_data_accesslog_src.TLSProperties{
                    TlsVersion:                 envoy_data_accesslog_src.TLSProperties_TLSv1_3,
                    TlsCipherSuite:             &wrappers.UInt32Value{Value: 2},
                    TlsSniHostname:             "3",
                    LocalCertificateProperties: &envoy_data_accesslog_src.TLSProperties_CertificateProperties{},
                    PeerCertificateProperties:  &envoy_data_accesslog_src.TLSProperties_CertificateProperties{},
                    TlsSessionId:               "6",
                },
            },
            want: &envoy_data_accesslog_v3.TLSProperties{
                TlsVersion:                 envoy_data_accesslog_v3.TLSProperties_TLSv1_3,
                TlsCipherSuite:             &wrappers.UInt32Value{Value: 2},
                TlsSniHostname:             "3",
                LocalCertificateProperties: &envoy_data_accesslog_v3.TLSProperties_CertificateProperties{},
                PeerCertificateProperties:  &envoy_data_accesslog_v3.TLSProperties_CertificateProperties{},
                TlsSessionId:               "6",
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_TLSProperties(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_TLSProperties_CertificateProperties(t *testing.T) {
    type args struct {
        src *envoy_data_accesslog_src.TLSProperties_CertificateProperties
    }
    tests := []struct {
        name string
        args args
        want *envoy_data_accesslog_v3.TLSProperties_CertificateProperties
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
                src: &envoy_data_accesslog_src.TLSProperties_CertificateProperties{
                    SubjectAltName: []*envoy_data_accesslog_src.TLSProperties_CertificateProperties_SubjectAltName{
                        {},
                    },
                    Subject: "123",
                },
            },
            want: &envoy_data_accesslog_v3.TLSProperties_CertificateProperties{
                SubjectAltName: []*envoy_data_accesslog_v3.TLSProperties_CertificateProperties_SubjectAltName{
                    {},
                },
                Subject: "123",
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_TLSProperties_CertificateProperties(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_TLSProperties_CertificateProperties_SubjectAltNames(t *testing.T) {
    type args struct {
        src []*envoy_data_accesslog_src.TLSProperties_CertificateProperties_SubjectAltName
    }
    tests := []struct {
        name string
        args args
        want []*envoy_data_accesslog_v3.TLSProperties_CertificateProperties_SubjectAltName
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
                src: []*envoy_data_accesslog_src.TLSProperties_CertificateProperties_SubjectAltName{
                    {},
                    {},
                },
            },
            want: []*envoy_data_accesslog_v3.TLSProperties_CertificateProperties_SubjectAltName{
                {},
                {},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_TLSProperties_CertificateProperties_SubjectAltNames(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_TLSProperties_CertificateProperties_SubjectAltName(t *testing.T) {
    type args struct {
        src *envoy_data_accesslog_src.TLSProperties_CertificateProperties_SubjectAltName
    }
    tests := []struct {
        name string
        args args
        want *envoy_data_accesslog_v3.TLSProperties_CertificateProperties_SubjectAltName
    }{
        {
            name: "Convert NIL",
            args: args{
                src: nil,
            },
            want: nil,
        },
        {
            name: "Convert OBJECT Uri",
            args: args{
                src: &envoy_data_accesslog_src.TLSProperties_CertificateProperties_SubjectAltName{
                    San: &envoy_data_accesslog_src.TLSProperties_CertificateProperties_SubjectAltName_Uri{
                        Uri: "uri",
                    },
                },
            },
            want: &envoy_data_accesslog_v3.TLSProperties_CertificateProperties_SubjectAltName{
                San: &envoy_data_accesslog_v3.TLSProperties_CertificateProperties_SubjectAltName_Uri{
                    Uri: "uri",
                },
            },
        },
        {
            name: "Convert OBJECT Dns",
            args: args{
                src: &envoy_data_accesslog_src.TLSProperties_CertificateProperties_SubjectAltName{
                    San: &envoy_data_accesslog_src.TLSProperties_CertificateProperties_SubjectAltName_Dns{
                        Dns: "dns",
                    },
                },
            },
            want: &envoy_data_accesslog_v3.TLSProperties_CertificateProperties_SubjectAltName{
                San: &envoy_data_accesslog_v3.TLSProperties_CertificateProperties_SubjectAltName_Dns{
                    Dns: "dns",
                },
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_TLSProperties_CertificateProperties_SubjectAltName(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_ResponseFlags(t *testing.T) {
    type args struct {
        src *envoy_data_accesslog_src.ResponseFlags
    }
    tests := []struct {
        name string
        args args
        want *envoy_data_accesslog_v3.ResponseFlags
    }{
        {
            name: "Convert NIL",
            args: args{
                src: nil,
            },
            want: nil,
        },
        {
            name: "Convert OBJECT FailedLocalHealthcheck",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    FailedLocalHealthcheck: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                FailedLocalHealthcheck: true,
            },
        },
        {
            name: "Convert OBJECT NoHealthyUpstream",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    NoHealthyUpstream: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                NoHealthyUpstream: true,
            },
        },
        {
            name: "Convert OBJECT UpstreamRequestTimeout",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    UpstreamRequestTimeout: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                UpstreamRequestTimeout: true,
            },
        },
        {
            name: "Convert OBJECT LocalReset",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    LocalReset: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                LocalReset: true,
            },
        },
        {
            name: "Convert OBJECT UpstreamRemoteReset",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    UpstreamRemoteReset: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                UpstreamRemoteReset: true,
            },
        },
        {
            name: "Convert OBJECT UpstreamConnectionFailure",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    UpstreamConnectionFailure: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                UpstreamConnectionFailure: true,
            },
        },
        {
            name: "Convert OBJECT UpstreamConnectionTermination",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    UpstreamConnectionTermination: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                UpstreamConnectionTermination: true,
            },
        },
        {
            name: "Convert OBJECT UpstreamOverflow",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    UpstreamOverflow: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                UpstreamOverflow: true,
            },
        },
        {
            name: "Convert OBJECT NoRouteFound",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    NoRouteFound: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                NoRouteFound: true,
            },
        },
        {
            name: "Convert OBJECT DelayInjected",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    DelayInjected: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                DelayInjected: true,
            },
        },
        {
            name: "Convert OBJECT FaultInjected",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    FaultInjected: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                FaultInjected: true,
            },
        },
        {
            name: "Convert OBJECT RateLimited",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    RateLimited: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                RateLimited: true,
            },
        },
        {
            name: "Convert OBJECT UnauthorizedDetails",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    UnauthorizedDetails: &envoy_data_accesslog_src.ResponseFlags_Unauthorized{},
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                UnauthorizedDetails: &envoy_data_accesslog_v3.ResponseFlags_Unauthorized{},
            },
        },
        {
            name: "Convert OBJECT RateLimitServiceError",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    RateLimitServiceError: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                RateLimitServiceError: true,
            },
        },
        {
            name: "Convert OBJECT DownstreamConnectionTermination",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    DownstreamConnectionTermination: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                DownstreamConnectionTermination: true,
            },
        },
        {
            name: "Convert OBJECT UpstreamRetryLimitExceeded",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    UpstreamRetryLimitExceeded: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                UpstreamRetryLimitExceeded: true,
            },
        },
        {
            name: "Convert OBJECT StreamIdleTimeout",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    StreamIdleTimeout: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                StreamIdleTimeout: true,
            },
        },
        {
            name: "Convert OBJECT InvalidEnvoyRequestHeaders",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    InvalidEnvoyRequestHeaders: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                InvalidEnvoyRequestHeaders: true,
            },
        },
        {
            name: "Convert OBJECT DownstreamProtocolError",
            args: args{
                src: &envoy_data_accesslog_src.ResponseFlags{
                    DownstreamProtocolError: true,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags{
                DownstreamProtocolError: true,
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_ResponseFlags(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_ResponseFlags_Unauthorized(t *testing.T) {
    type args struct {
        src *envoy_data_accesslog_src.ResponseFlags_Unauthorized
    }
    tests := []struct {
        name string
        args args
        want *envoy_data_accesslog_v3.ResponseFlags_Unauthorized
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
                src: &envoy_data_accesslog_src.ResponseFlags_Unauthorized{
                    Reason: envoy_data_accesslog_src.ResponseFlags_Unauthorized_EXTERNAL_SERVICE,
                },
            },
            want: &envoy_data_accesslog_v3.ResponseFlags_Unauthorized{
                Reason: envoy_data_accesslog_v3.ResponseFlags_Unauthorized_EXTERNAL_SERVICE,
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_ResponseFlags_Unauthorized(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_Metadata(t *testing.T) {
    type args struct {
        src *envoy_api_v2_core.Metadata
    }
    tests := []struct {
        name string
        args args
        want *envoy_config_core_v3.Metadata
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
                src: &envoy_api_v2_core.Metadata{
                    FilterMetadata: map[string]*structpb.Struct{"1.k": nil},
                },
            },
            want: &envoy_config_core_v3.Metadata{
                FilterMetadata: map[string]*structpb.Struct{"1.k": nil},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_Metadata(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}
