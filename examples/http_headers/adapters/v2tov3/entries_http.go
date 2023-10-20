package v2tov3

import (
    envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
    envoy_data_accesslog_src "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v2"
    envoy_data_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
)

//
// Log Entries
// #################################################################################################################

//
// Http Logs
// #################################################################################################################

// adapt_HTTPAccessLogEntries
func adapt_HTTPAccessLogEntries(src_array []*envoy_data_accesslog_src.HTTPAccessLogEntry) []*envoy_data_accesslog_v3.HTTPAccessLogEntry {
    if src_array == nil {
        return nil
    }

    dst_array := make([]*envoy_data_accesslog_v3.HTTPAccessLogEntry, 0)
    for _, src := range src_array {
        dst_array = append(dst_array, adapt_HTTPAccessLogEntry(src))
    }

    return dst_array
}

// adapt_HTTPAccessLogEntry
func adapt_HTTPAccessLogEntry(src *envoy_data_accesslog_src.HTTPAccessLogEntry) *envoy_data_accesslog_v3.HTTPAccessLogEntry {
    if src == nil {
        return nil
    }

    return &envoy_data_accesslog_v3.HTTPAccessLogEntry{
        CommonProperties: adapt_AccessLogCommon(src.CommonProperties),
        ProtocolVersion:  envoy_data_accesslog_v3.HTTPAccessLogEntry_HTTPVersion(src.ProtocolVersion),
        Request:          adapt_HTTPRequestProperties(src.GetRequest()),
        Response:         adapt_HTTPResponseProperties(src.GetResponse()),
    }
}

// adapt_HTTPRequestProperties
func adapt_HTTPRequestProperties(src *envoy_data_accesslog_src.HTTPRequestProperties) *envoy_data_accesslog_v3.HTTPRequestProperties {
    if src == nil {
        return nil
    }

    return &envoy_data_accesslog_v3.HTTPRequestProperties{
        RequestMethod:       envoy_config_core_v3.RequestMethod(src.GetRequestMethod()),
        Scheme:              src.GetScheme(),
        Authority:           src.GetAuthority(),
        Port:                src.GetPort(),
        Path:                src.GetPath(),
        UserAgent:           src.GetUserAgent(),
        Referer:             src.GetReferer(),
        ForwardedFor:        src.GetForwardedFor(),
        RequestId:           src.GetRequestId(),
        OriginalPath:        src.GetOriginalPath(),
        RequestHeadersBytes: src.GetRequestHeadersBytes(),
        RequestBodyBytes:    src.GetRequestBodyBytes(),
        RequestHeaders:      src.GetRequestHeaders(),
    }
}

// adapt_HTTPResponseProperties
func adapt_HTTPResponseProperties(src *envoy_data_accesslog_src.HTTPResponseProperties) *envoy_data_accesslog_v3.HTTPResponseProperties {
    if src == nil {
        return nil
    }

    return &envoy_data_accesslog_v3.HTTPResponseProperties{
        ResponseCode:         src.GetResponseCode(),
        ResponseHeadersBytes: src.GetResponseHeadersBytes(),
        ResponseBodyBytes:    src.GetResponseBodyBytes(),
        ResponseHeaders:      src.GetResponseHeaders(),
        ResponseTrailers:     src.GetResponseTrailers(),
        ResponseCodeDetails:  src.GetResponseCodeDetails(),
    }
}
