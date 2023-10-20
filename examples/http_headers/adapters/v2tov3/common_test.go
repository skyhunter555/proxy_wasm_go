package v2tov3

import (
    envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
    envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
    envoy_type "github.com/envoyproxy/go-control-plane/envoy/type"
    envoy_type_v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_adapt_SemanticVersion(t *testing.T) {
    type args struct {
        src *envoy_type.SemanticVersion
    }
    tests := []struct {
        name string
        args args
        want *envoy_type_v3.SemanticVersion
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
                src: &envoy_type.SemanticVersion{
                    MajorNumber: 1,
                    MinorNumber: 2,
                    Patch:       3,
                },
            },
            want: &envoy_type_v3.SemanticVersion{
                MajorNumber: 1,
                MinorNumber: 2,
                Patch:       3,
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_SemanticVersion(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_BuildVersion(t *testing.T) {
    type args struct {
        src *envoy_api_v2_core.BuildVersion
    }
    tests := []struct {
        name string
        args args
        want *envoy_config_core_v3.BuildVersion
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
                src: &envoy_api_v2_core.BuildVersion{
                    Version:  &envoy_type.SemanticVersion{},
                    Metadata: nil,
                },
            },
            want: &envoy_config_core_v3.BuildVersion{
                Version:  &envoy_type_v3.SemanticVersion{},
                Metadata: nil,
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_BuildVersion(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_Addresses(t *testing.T) {
    type args struct {
        src []*envoy_api_v2_core.Address
    }
    tests := []struct {
        name string
        args args
        want []*envoy_config_core_v3.Address
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
                src: []*envoy_api_v2_core.Address{
                    {},
                    {},
                },
            },
            want: []*envoy_config_core_v3.Address{
                {},
                {},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_Addresses(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}

func Test_adapt_Address(t *testing.T) {
    type args struct {
        src *envoy_api_v2_core.Address
    }
    tests := []struct {
        name string
        args args
        want *envoy_config_core_v3.Address
    }{
        {
            name: "Convert NIL",
            args: args{
                src: nil,
            },
            want: nil,
        },
        {
            name: "Convert Socket NIL",
            args: args{
                src: &envoy_api_v2_core.Address{
                    Address: &envoy_api_v2_core.Address_SocketAddress{},
                },
            },
            want: &envoy_config_core_v3.Address{},
        },
        {
            name: "Convert Socket PortValue OBJECT",
            args: args{
                src: &envoy_api_v2_core.Address{
                    Address: &envoy_api_v2_core.Address_SocketAddress{
                        SocketAddress: &envoy_api_v2_core.SocketAddress{
                            Protocol:      envoy_api_v2_core.SocketAddress_TCP,
                            Address:       "1",
                            PortSpecifier: &envoy_api_v2_core.SocketAddress_PortValue{PortValue: 2},
                            ResolverName:  "3",
                            Ipv4Compat:    false,
                        },
                    },
                },
            },
            want: &envoy_config_core_v3.Address{
                Address: &envoy_config_core_v3.Address_SocketAddress{
                    SocketAddress: &envoy_config_core_v3.SocketAddress{
                        Protocol:      envoy_config_core_v3.SocketAddress_TCP,
                        Address:       "1",
                        PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{PortValue: 2},
                        ResolverName:  "3",
                        Ipv4Compat:    false,
                    },
                },
            },
        },
        {
            name: "Convert Socket PortValue OBJECT",
            args: args{
                src: &envoy_api_v2_core.Address{
                    Address: &envoy_api_v2_core.Address_SocketAddress{
                        SocketAddress: &envoy_api_v2_core.SocketAddress{
                            Protocol:      envoy_api_v2_core.SocketAddress_UDP,
                            Address:       "1",
                            PortSpecifier: &envoy_api_v2_core.SocketAddress_NamedPort{NamedPort: "2"},
                            ResolverName:  "3",
                            Ipv4Compat:    true,
                        },
                    },
                },
            },
            want: &envoy_config_core_v3.Address{
                Address: &envoy_config_core_v3.Address_SocketAddress{
                    SocketAddress: &envoy_config_core_v3.SocketAddress{
                        Protocol:      envoy_config_core_v3.SocketAddress_UDP,
                        Address:       "1",
                        PortSpecifier: &envoy_config_core_v3.SocketAddress_NamedPort{NamedPort: "2"},
                        ResolverName:  "3",
                        Ipv4Compat:    true,
                    },
                },
            },
        },
        {
            name: "Convert Pipe NIL",
            args: args{
                src: &envoy_api_v2_core.Address{
                    Address: &envoy_api_v2_core.Address_Pipe{},
                },
            },
            want: &envoy_config_core_v3.Address{},
        },
        {
            name: "Convert Pipe OBJECT",
            args: args{
                src: &envoy_api_v2_core.Address{
                    Address: &envoy_api_v2_core.Address_Pipe{
                        Pipe: &envoy_api_v2_core.Pipe{
                            Path: "123",
                            Mode: 321,
                        },
                    },
                },
            },
            want: &envoy_config_core_v3.Address{
                Address: &envoy_config_core_v3.Address_Pipe{
                    Pipe: &envoy_config_core_v3.Pipe{
                        Path: "123",
                        Mode: 321,
                    },
                },
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_Address(tt.args.src)
            assert.Equal(t, tt.want, got)
        })
    }
}
