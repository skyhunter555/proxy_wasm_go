package v2tov3

import (
    envoy_service_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
    envoy_service_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
)

//
// Message
// #################################################################################################################

// adapt_Message
func adapt_Message(src *envoy_service_accesslog_v2.StreamAccessLogsMessage) *envoy_service_accesslog_v3.StreamAccessLogsMessage {
    if src == nil {
        return nil
    }

    dst := &envoy_service_accesslog_v3.StreamAccessLogsMessage{}
    dst.Identifier = adapt_Identifier(src.GetIdentifier())

    if src_entries := src.GetHttpLogs(); src_entries != nil {
        dst.LogEntries = &envoy_service_accesslog_v3.StreamAccessLogsMessage_HttpLogs{
            HttpLogs: &envoy_service_accesslog_v3.StreamAccessLogsMessage_HTTPAccessLogEntries{
                LogEntry: adapt_HTTPAccessLogEntries(src_entries.LogEntry),
            },
        }
    }

    if src_entries := src.GetTcpLogs(); src_entries != nil {
        dst.LogEntries = &envoy_service_accesslog_v3.StreamAccessLogsMessage_TcpLogs{
            TcpLogs: &envoy_service_accesslog_v3.StreamAccessLogsMessage_TCPAccessLogEntries{
                LogEntry: adapt_TcpAccessLogEntries(src_entries.LogEntry),
            },
        }
    }

    return dst
}
