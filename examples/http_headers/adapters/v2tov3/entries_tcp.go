package v2tov3

import (
    envoy_data_accesslog_src "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v2"
    envoy_data_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
)

//
// Log Entries
// #################################################################################################################

//
// Tcp Logs
// #################################################################################################################

// adapt_TcpAccessLogEntries
func adapt_TcpAccessLogEntries(src_array []*envoy_data_accesslog_src.TCPAccessLogEntry) []*envoy_data_accesslog_v3.TCPAccessLogEntry {
    if src_array == nil {
        return nil
    }

    dst_array := make([]*envoy_data_accesslog_v3.TCPAccessLogEntry, 0)
    for _, src := range src_array {
        dst_array = append(dst_array, adapt_TcpAccessLogEntry(src))
    }
    return dst_array
}

// adapt_TcpAccessLogEntry
func adapt_TcpAccessLogEntry(src *envoy_data_accesslog_src.TCPAccessLogEntry) *envoy_data_accesslog_v3.TCPAccessLogEntry {
    if src == nil {
        return nil
    }

    return &envoy_data_accesslog_v3.TCPAccessLogEntry{
        CommonProperties:     adapt_AccessLogCommon(src.CommonProperties),
        ConnectionProperties: adapt_ConnectionProperties(src.ConnectionProperties),
    }
}

// adapt_ConnectionProperties
func adapt_ConnectionProperties(src *envoy_data_accesslog_src.ConnectionProperties) *envoy_data_accesslog_v3.ConnectionProperties {
    if src == nil {
        return nil
    }

    return &envoy_data_accesslog_v3.ConnectionProperties{
        ReceivedBytes: src.GetReceivedBytes(),
        SentBytes:     src.GetSentBytes(),
    }
}
