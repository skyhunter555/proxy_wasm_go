package v3to_vjcal

import (
    "fmt"
    "strings"

    envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
    envoy_data_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
    envoy_service_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"

    "bitbucket.region.vtb.ru/projects/USBP/repos/envoy-accesslogs/internal/vtbjournalclient_accesslogs"
)

const (
    HEADER_X_B3_TRACE_ID           = "x-b3-traceid"
    HEADER_X_FORWARDER_FOR         = "x-forwarded-for"
    HEADER_X_FORWARDER_CLIENT_CERT = "x-forwarded-client-cert"
    HEADER_X_RATELIMIT_LIMIT       = "x-ratelimit-limit"
    HEADER_X_RATELIMIT_REMAINING   = "x-ratelimit-remaining"
    HEADER_X_RATELIMIT_RESET       = "x-ratelimit-reset"
    HEADER_HOST                    = "host"
    HEADER_X_ENVOY_ORIGINAL_PATH   = "x-envoy-original-path"
)

// adapt_Message
func adapt_Message(src *envoy_service_accesslog_v3.StreamAccessLogsMessage) []*vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog {
    dst := make([]*vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog, 0)

    // http entries
    if src_entries := src.GetHttpLogs(); src_entries != nil {
        for _, src_entry := range src_entries.GetLogEntry() {
            dst = append(dst, adapt_HTTP(src_entry))
        }
    }

    // tcp entries
    if src_entries := src.GetTcpLogs(); src_entries != nil {
        for _, src_entry := range src_entries.GetLogEntry() {
            dst = append(dst, adapt_TCP(src_entry))
        }
    }

    return dst
}

// adapt_HTTP
func adapt_HTTP(src *envoy_data_accesslog_v3.HTTPAccessLogEntry) *vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog {
    if src == nil {
        return nil
    }

    //
    dst := &vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog{}
    dst.Protocol = envoy_data_accesslog_v3.HTTPAccessLogEntry_HTTPVersion_name[int32(src.GetProtocolVersion())]

    //
    if src_cmp := src.GetCommonProperties(); src_cmp != nil {
        dst.UpstreamLocalAddress = adapt_Address(src_cmp.GetUpstreamLocalAddress())
        dst.Duration = src_cmp.GetTimeToLastDownstreamTxByte().AsDuration()
        dst.UpstreamTransportFailureReason = src_cmp.GetUpstreamTransportFailureReason()
        dst.DownstreamLocalAddress = adapt_Address(src_cmp.GetDownstreamLocalAddress())
        dst.ResponseFlags = adapt_ResponseFlags(src_cmp.GetResponseFlags())
        dst.StartTime = src_cmp.GetStartTime().AsTime()
        dst.RequestedServerName = src_cmp.GetTlsProperties().GetTlsSniHostname()
        dst.UpstreamHost = adapt_Address(src_cmp.GetUpstreamRemoteAddress())
        dst.UpstreamCluster = src_cmp.GetUpstreamCluster()
        dst.DownstreamRemoteAddress = adapt_Address(src_cmp.GetDownstreamRemoteAddress())
    }

    //
    if src_rq := src.GetRequest(); src_rq != nil {
        dst.TraceId = src_rq.GetRequestHeaders()[HEADER_X_B3_TRACE_ID]
        dst.UserAgent = src_rq.GetUserAgent()
        dst.RequestId = src_rq.GetRequestId()
        dst.XForwardedFor = src_rq.GetRequestHeaders()[HEADER_X_FORWARDER_FOR]
        dst.XForwardedClientCert = src_rq.GetRequestHeaders()[HEADER_X_FORWARDER_CLIENT_CERT]
        dst.XRatelimitLimit = src_rq.GetRequestHeaders()[HEADER_X_RATELIMIT_LIMIT]
        dst.XRatelimitRemaining = src_rq.GetRequestHeaders()[HEADER_X_RATELIMIT_REMAINING]
        dst.XRatelimitReset = src_rq.GetRequestHeaders()[HEADER_X_RATELIMIT_RESET]
        dst.BytesReceived = src_rq.GetRequestBodyBytes()
        dst.Authority = src_rq.GetAuthority()

        dst.Host = src_rq.GetRequestHeaders()[HEADER_HOST]
        dst.Method = envoy_config_core_v3.RequestMethod_name[int32(src_rq.GetRequestMethod())]

        // Получение HTTP PATH, выбирается первый не пустой
        paths := []string{
            src_rq.GetOriginalPath(),
            src_rq.GetRequestHeaders()[HEADER_X_ENVOY_ORIGINAL_PATH],
            src_rq.GetPath(),
        }
        for idx := 0; idx < len(paths) && (dst.Path == "" || dst.Path == "-"); idx++ {
            dst.Path = paths[idx]
        }
    }

    //
    if src_rs := src.GetResponse(); src_rs != nil {
        if src_rs_rc := src_rs.GetResponseCode(); src_rs_rc != nil {
            dst.ResponseCode = src_rs_rc.GetValue()
        }

        dst.BytesSent = src_rs.GetResponseBodyBytes()
    }

    return dst
}

// adapt_TCP
func adapt_TCP(src *envoy_data_accesslog_v3.TCPAccessLogEntry) *vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog {
    if src == nil {
        return nil
    }

    //
    dst := &vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog{}
    dst.Protocol = "TCP"

    //
    if src_cmp := src.GetCommonProperties(); src_cmp != nil {
        dst.UpstreamLocalAddress = adapt_Address(src_cmp.GetUpstreamLocalAddress())
        dst.Duration = src_cmp.GetTimeToLastDownstreamTxByte().AsDuration()
        dst.UpstreamTransportFailureReason = src_cmp.GetUpstreamTransportFailureReason()
        dst.DownstreamLocalAddress = adapt_Address(src_cmp.GetDownstreamLocalAddress())
        dst.ResponseFlags = adapt_ResponseFlags(src_cmp.GetResponseFlags())
        dst.StartTime = src_cmp.GetStartTime().AsTime()
        dst.UpstreamHost = adapt_Address(src_cmp.GetUpstreamRemoteAddress())
        dst.UpstreamCluster = src_cmp.GetUpstreamCluster()
        dst.DownstreamRemoteAddress = adapt_Address(src_cmp.GetDownstreamRemoteAddress())

        //
        if src_cmp_tlsp := src_cmp.GetTlsProperties(); src_cmp_tlsp != nil {
            dst.RequestedServerName = src_cmp_tlsp.GetTlsSniHostname()
        }
    }

    //
    if src_cnp := src.GetConnectionProperties(); src_cnp != nil {
        dst.BytesReceived = src_cnp.GetReceivedBytes()
        dst.BytesSent = src_cnp.GetSentBytes()
    }

    return dst
}

// adapt_Address
func adapt_Address(src *envoy_config_core_v3.Address) string {
    if src == nil {
        return ""
    }

    // socket
    if src_socket := src.GetSocketAddress(); src_socket != nil {
        protocol := strings.ToLower(envoy_config_core_v3.SocketAddress_Protocol_name[int32(src_socket.GetProtocol())])
        address := src_socket.GetAddress()

        port := ""
        if src_port_value := src_socket.GetPortValue(); src_port_value != 0 {
            port = fmt.Sprintf(":%d", src_port_value)
        }
        if src_named_port := src_socket.GetNamedPort(); src_named_port != "" {
            port = fmt.Sprintf(":[%s]", src_named_port)
        }

        return fmt.Sprintf("%s://%s%s", protocol, address, port)
    }

    // envoy internal
    if src_envoy_internal := src.GetEnvoyInternalAddress(); src_envoy_internal != nil {
        if src_server_listener_name := src_envoy_internal.GetServerListenerName(); src_server_listener_name != "" {
            return src_server_listener_name
        }
    }

    // pipe
    if src_pipe := src.GetPipe(); src_pipe != nil {
        return fmt.Sprintf("unix://%s", src_pipe.Path)
    }

    return ""
}
