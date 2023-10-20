package v2tov3

import (
    "testing"

    envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
    envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
    envoy_service_accesslog_src "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
    envoy_service_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
    "github.com/stretchr/testify/assert"
)

func Test_adapt_Identifier(t *testing.T) {
    type args struct {
        src *envoy_service_accesslog_src.StreamAccessLogsMessage_Identifier
    }
    tests := []struct {
        name string
        args args
        want *envoy_service_accesslog_v3.StreamAccessLogsMessage_Identifier
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
                src: &envoy_service_accesslog_src.StreamAccessLogsMessage_Identifier{
                    Node:    &envoy_api_v2_core.Node{},
                    LogName: "1",
                },
            },
            want: &envoy_service_accesslog_v3.StreamAccessLogsMessage_Identifier{
                Node:    &envoy_config_core_v3.Node{},
                LogName: "1",
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_Identifier(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_Node(t *testing.T) {
    type args struct {
        src *envoy_api_v2_core.Node
    }
    tests := []struct {
        name string
        args args
        want *envoy_config_core_v3.Node
    }{
        {
            name: "Convert NIL",
            args: args{
                src: nil,
            },
            want: nil,
        },
        {
            name: "Convert OBJECT UserAgentBuildVersion",
            args: args{
                src: &envoy_api_v2_core.Node{
                    Id:                   "1",
                    Cluster:              "2",
                    Metadata:             nil,
                    Locality:             &envoy_api_v2_core.Locality{},
                    BuildVersion:         "5",
                    UserAgentName:        "6",
                    UserAgentVersionType: &envoy_api_v2_core.Node_UserAgentBuildVersion{UserAgentBuildVersion: &envoy_api_v2_core.BuildVersion{}},
                    Extensions:           []*envoy_api_v2_core.Extension{},
                    ClientFeatures:       []string{},
                    ListeningAddresses:   []*envoy_api_v2_core.Address{},
                },
            },
            want: &envoy_config_core_v3.Node{
                Id:                   "1",
                Cluster:              "2",
                Metadata:             nil,
                Locality:             &envoy_config_core_v3.Locality{},
                UserAgentName:        "6",
                UserAgentVersionType: &envoy_config_core_v3.Node_UserAgentBuildVersion{UserAgentBuildVersion: &envoy_config_core_v3.BuildVersion{}},
                Extensions:           []*envoy_config_core_v3.Extension{},
                ClientFeatures:       []string{},
                ListeningAddresses:   []*envoy_config_core_v3.Address{},
            },
        },
        {
            name: "Convert OBJECT UserAgentVersion",
            args: args{
                src: &envoy_api_v2_core.Node{
                    Id:                   "1",
                    Cluster:              "2",
                    Metadata:             nil,
                    Locality:             &envoy_api_v2_core.Locality{},
                    BuildVersion:         "5",
                    UserAgentName:        "6",
                    UserAgentVersionType: &envoy_api_v2_core.Node_UserAgentVersion{UserAgentVersion: "123"},
                    Extensions:           []*envoy_api_v2_core.Extension{},
                    ClientFeatures:       []string{},
                    ListeningAddresses:   []*envoy_api_v2_core.Address{},
                },
            },
            want: &envoy_config_core_v3.Node{
                Id:                   "1",
                Cluster:              "2",
                Metadata:             nil,
                Locality:             &envoy_config_core_v3.Locality{},
                UserAgentName:        "6",
                UserAgentVersionType: &envoy_config_core_v3.Node_UserAgentVersion{UserAgentVersion: "123"},
                Extensions:           []*envoy_config_core_v3.Extension{},
                ClientFeatures:       []string{},
                ListeningAddresses:   []*envoy_config_core_v3.Address{},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_Node(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_Locality(t *testing.T) {
    type args struct {
        src *envoy_api_v2_core.Locality
    }
    tests := []struct {
        name string
        args args
        want *envoy_config_core_v3.Locality
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
                src: &envoy_api_v2_core.Locality{
                    Region:  "1",
                    Zone:    "2",
                    SubZone: "3",
                },
            },
            want: &envoy_config_core_v3.Locality{
                Region:  "1",
                Zone:    "2",
                SubZone: "3",
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_Locality(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_Extensions(t *testing.T) {
    type args struct {
        src []*envoy_api_v2_core.Extension
    }
    tests := []struct {
        name string
        args args
        want []*envoy_config_core_v3.Extension
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
                src: []*envoy_api_v2_core.Extension{
                    {},
                    {},
                },
            },
            want: []*envoy_config_core_v3.Extension{
                {},
                {},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_Extensions(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_Extension(t *testing.T) {
    type args struct {
        src *envoy_api_v2_core.Extension
    }
    tests := []struct {
        name string
        args args
        want *envoy_config_core_v3.Extension
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
                src: &envoy_api_v2_core.Extension{
                    Name:           "1",
                    Category:       "2",
                    TypeDescriptor: "3",
                    Version:        &envoy_api_v2_core.BuildVersion{},
                    Disabled:       false,
                },
            },
            want: &envoy_config_core_v3.Extension{
                Name:           "1",
                Category:       "2",
                TypeDescriptor: "3",
                Version:        &envoy_config_core_v3.BuildVersion{},
                Disabled:       false,
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_Extension(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}
