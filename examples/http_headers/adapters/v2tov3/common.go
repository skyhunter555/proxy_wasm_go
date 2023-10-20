package v2tov3

import (
    envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
    envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
    envoy_type "github.com/envoyproxy/go-control-plane/envoy/type"
    envoy_type_v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
)

//
// Common
// #################################################################################################################

// adapt_SemanticVersion
func adapt_SemanticVersion(src *envoy_type.SemanticVersion) *envoy_type_v3.SemanticVersion {
    if src == nil {
        return nil
    }

    return &envoy_type_v3.SemanticVersion{
        MajorNumber: src.GetMajorNumber(),
        MinorNumber: src.GetMinorNumber(),
        Patch:       src.GetPatch(),
    }
}

// adapt_BuildVersion
func adapt_BuildVersion(src *envoy_api_v2_core.BuildVersion) *envoy_config_core_v3.BuildVersion {
    if src == nil {
        return nil
    }

    return &envoy_config_core_v3.BuildVersion{
        Version:  adapt_SemanticVersion(src.GetVersion()),
        Metadata: src.GetMetadata(),
    }
}

// adapt_Addresses
func adapt_Addresses(src_array []*envoy_api_v2_core.Address) []*envoy_config_core_v3.Address {
    if src_array == nil {
        return nil
    }

    dst_array := make([]*envoy_config_core_v3.Address, 0)
    for _, src := range src_array {
        dst_array = append(dst_array, adapt_Address(src))
    }

    return dst_array
}

// adapt_Address
func adapt_Address(src *envoy_api_v2_core.Address) *envoy_config_core_v3.Address {
    if src == nil {
        return nil
    }

    // Socket
    if src_socket := src.GetSocketAddress(); src_socket != nil {
        return &envoy_config_core_v3.Address{
            Address: &envoy_config_core_v3.Address_SocketAddress{
                SocketAddress: adapt_SocketAddress(src_socket),
            },
        }
    }

    // EnvoyInternalAddress
    // No source data, nothing todo;

    // Pipe
    if src_pipe := src.GetPipe(); src_pipe != nil {
        return &envoy_config_core_v3.Address{
            Address: &envoy_config_core_v3.Address_Pipe{
                Pipe: adapt_Pipe(src_pipe),
            },
        }
    }

    return &envoy_config_core_v3.Address{}
}

// adapt_SocketAddress
func adapt_SocketAddress(src *envoy_api_v2_core.SocketAddress) *envoy_config_core_v3.SocketAddress {
    if src == nil {
        return nil
    }

    dst := &envoy_config_core_v3.SocketAddress{
        Protocol:      envoy_config_core_v3.SocketAddress_Protocol(src.Protocol),
        Address:       src.GetAddress(),
        PortSpecifier: nil,
        ResolverName:  src.GetResolverName(),
        Ipv4Compat:    src.GetIpv4Compat(),
    }

    if src_portValue, ok := src.PortSpecifier.(*envoy_api_v2_core.SocketAddress_PortValue); ok {
        dst.PortSpecifier = &envoy_config_core_v3.SocketAddress_PortValue{
            PortValue: src_portValue.PortValue,
        }
    }
    if src_namedPort, ok := src.PortSpecifier.(*envoy_api_v2_core.SocketAddress_NamedPort); ok {
        dst.PortSpecifier = &envoy_config_core_v3.SocketAddress_NamedPort{
            NamedPort: src_namedPort.NamedPort,
        }
    }

    return dst
}

// adapt_Pipe
func adapt_Pipe(src *envoy_api_v2_core.Pipe) *envoy_config_core_v3.Pipe {
    if src == nil {
        return nil
    }

    return &envoy_config_core_v3.Pipe{
        Path: src.Path,
        Mode: src.Mode,
    }
}
