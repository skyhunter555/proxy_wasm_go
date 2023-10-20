package v2tov3

import (
    envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
    envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
    envoy_service_accesslog_src "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
    envoy_service_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
)

//
// Identifier
// #################################################################################################################

// adapt_Identifier
func adapt_Identifier(src *envoy_service_accesslog_src.StreamAccessLogsMessage_Identifier) *envoy_service_accesslog_v3.StreamAccessLogsMessage_Identifier {
    if src == nil {
        return nil
    }

    return &envoy_service_accesslog_v3.StreamAccessLogsMessage_Identifier{
        Node:    adapt_Node(src.GetNode()),
        LogName: src.GetLogName(),
    }
}

// adapt_Node
func adapt_Node(src *envoy_api_v2_core.Node) *envoy_config_core_v3.Node {
    if src == nil {
        return nil
    }

    //
    dst := &envoy_config_core_v3.Node{
        Id:                   src.GetId(),
        Cluster:              src.GetCluster(),
        Metadata:             src.GetMetadata(),
        Locality:             adapt_Locality(src.GetLocality()),
        UserAgentName:        src.GetUserAgentName(),
        UserAgentVersionType: nil,
        Extensions:           adapt_Extensions(src.GetExtensions()),
        ClientFeatures:       src.GetClientFeatures(),
        ListeningAddresses:   adapt_Addresses(src.GetListeningAddresses()),
    }

    //
    if src_userAgentBuildVersion := src.GetUserAgentBuildVersion(); src_userAgentBuildVersion != nil {
        dst.UserAgentVersionType = &envoy_config_core_v3.Node_UserAgentBuildVersion{
            UserAgentBuildVersion: adapt_BuildVersion(src_userAgentBuildVersion),
        }
    }

    if src_userAgentVersion := src.GetUserAgentVersion(); src_userAgentVersion != "" {
        dst.UserAgentVersionType = &envoy_config_core_v3.Node_UserAgentVersion{
            UserAgentVersion: src_userAgentVersion,
        }
    }

    return dst
}

// adapt_Locality
func adapt_Locality(src *envoy_api_v2_core.Locality) *envoy_config_core_v3.Locality {
    if src == nil {
        return nil
    }

    return &envoy_config_core_v3.Locality{
        Region:  src.GetRegion(),
        Zone:    src.GetZone(),
        SubZone: src.GetSubZone(),
    }
}

// adapt_Extensions
func adapt_Extensions(src_array []*envoy_api_v2_core.Extension) []*envoy_config_core_v3.Extension {
    if src_array == nil {
        return nil
    }

    dst_array := make([]*envoy_config_core_v3.Extension, 0)
    for _, src := range src_array {
        dst_array = append(dst_array, adapt_Extension(src))
    }

    return dst_array
}

// adapt_Extension
func adapt_Extension(src *envoy_api_v2_core.Extension) *envoy_config_core_v3.Extension {
    if src == nil {
        return nil
    }

    return &envoy_config_core_v3.Extension{
        Name:           src.GetName(),
        Category:       src.GetCategory(),
        TypeDescriptor: src.GetTypeDescriptor(),
        Version:        adapt_BuildVersion(src.GetVersion()),
        Disabled:       src.GetDisabled(),
    }
}
