package v3to_vjcal

import (
    "strings"

    v3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
)

// См. https://github.com/envoyproxy/envoy/blob/main/source/common/stream_info/utility.h#L21
const (
    RF_FAILED_LOCAL_HEALTH_CHECK            = "LH"    // Local service failed health check request in addition to 503 response code.
    RF_NO_HEALTHY_UPSTREAM                  = "UH"    // No healthy upstream hosts in upstream cluster in addition to 503 response code.
    RF_UPSTREAM_REQUEST_TIMEOUT             = "UT"    // Upstream request timeout in addition to 504 response code.
    RF_LOCAL_RESET                          = "LR"    // Connection local reset in addition to 503 response code.
    RF_UPSTREAM_REMOTE_RESET                = "UR"    // Upstream remote reset in addition to 503 response code.
    RF_UPSTREAM_CONNECTION_FAILURE          = "UF"    // Upstream connection failure in addition to 503 response code.
    RF_UPSTREAM_CONNECTION_TERMINATION      = "UC"    // Upstream connection termination in addition to 503 response code.
    RF_UPSTREAM_OVERFLOW                    = "UO"    // Upstream overflow (circuit breaking) in addition to 503 response code.
    RF_NO_ROUTE_FOUND                       = "NR"    // No route configured for a given request in addition to 404 response code, or no matching filter chain for a downstream connection.
    RF_DELAY_INJECTED                       = "DI"    // The request processing was delayed for a period specified via fault injection.
    RF_FAULT_INJECTED                       = "FI"    // The request was aborted with a response code specified via fault injection.
    RF_RATE_LIMITED                         = "RL"    // The request was ratelimited locally by the HTTP rate limit filter in addition to 429 response code.
    RF_UNAUTHORIZED_EXTERNAL_SERVICE        = "UAEX"  // The request was denied by the external authorization service.
    RF_RATELIMIT_SERVICE_ERROR              = "RLSE"  // The request was rejected because there was an error in rate limit service.
    RF_DOWNSTREAM_CONNECTION_TERMINATION    = "DC"    // Downstream connection termination.
    RF_UPSTREAM_RETRY_LIMIT_EXCEEDED        = "URX"   // The request was rejected because the upstream retry limit (HTTP) or maximum connect attempts (TCP) was reached.
    RF_STREAM_IDLE_TIMEOUT                  = "SI"    // Stream idle timeout in addition to 408 response code.
    RF_INVALID_ENVOY_REQUEST_HEADERS        = "IH"    // The request was rejected because it set an invalid value for a strictly-checked header in addition to 400 response code.
    RF_DOWNSTREAM_PROTOCOL_ERROR            = "DPE"   // The downstream request had an HTTP protocol error.
    RF_UPSTREAM_MAX_STREAM_DURATION_REACHED = "UMSDR" // The upstream request reached to max stream duration.
    RF_RESPONSE_FROM_CACHE_FILTER           = "RFCF"  // ???
    RF_NO_FILTER_CONFIG_FOUND               = "NFCF"  // ???
    RF_DURATION_TIMEOUT                     = "DT"    // When a request or connection exceeded max_connection_duration or max_downstream_connection_duration
    RF_UPSTREAM_PROTOCOL_ERROR              = "UPE"   // The upstream response had an HTTP protocol error.
    RF_NO_CLUSTER_FOUND                     = "NC"    // Upstream cluster not found.
    RF_OVERLOAD_MANAGER                     = "OM"    // Overload Manager terminated the request
)

// responseFlagFunc
type responseFlagFunc func(rf *v3.ResponseFlags) bool

type responseFlagHandler struct {
    Flag string
    Func responseFlagFunc
}

func rh(Flag string, Func responseFlagFunc) *responseFlagHandler {
    return &responseFlagHandler{
        Flag: Flag,
        Func: Func,
    }
}

var (
    // allResponseFlagFunctions
    allResponseFlagHandlers = []*responseFlagHandler{
        {Flag: RF_FAILED_LOCAL_HEALTH_CHECK, Func: func(rf *v3.ResponseFlags) bool { return rf.GetFailedLocalHealthcheck() }},
        {Flag: RF_NO_HEALTHY_UPSTREAM, Func: func(rf *v3.ResponseFlags) bool { return rf.GetNoHealthyUpstream() }},
        {Flag: RF_UPSTREAM_REQUEST_TIMEOUT, Func: func(rf *v3.ResponseFlags) bool { return rf.GetUpstreamRequestTimeout() }},
        {Flag: RF_LOCAL_RESET, Func: func(rf *v3.ResponseFlags) bool { return rf.GetLocalReset() }},
        {Flag: RF_UPSTREAM_REMOTE_RESET, Func: func(rf *v3.ResponseFlags) bool { return rf.GetUpstreamRemoteReset() }},
        {Flag: RF_UPSTREAM_CONNECTION_FAILURE, Func: func(rf *v3.ResponseFlags) bool { return rf.GetUpstreamConnectionFailure() }},
        {Flag: RF_UPSTREAM_CONNECTION_TERMINATION, Func: func(rf *v3.ResponseFlags) bool { return rf.GetUpstreamConnectionTermination() }},
        {Flag: RF_UPSTREAM_OVERFLOW, Func: func(rf *v3.ResponseFlags) bool { return rf.GetUpstreamOverflow() }},
        {Flag: RF_NO_ROUTE_FOUND, Func: func(rf *v3.ResponseFlags) bool { return rf.GetNoRouteFound() }},
        {Flag: RF_DELAY_INJECTED, Func: func(rf *v3.ResponseFlags) bool { return rf.GetDelayInjected() }},
        {Flag: RF_FAULT_INJECTED, Func: func(rf *v3.ResponseFlags) bool { return rf.GetFaultInjected() }},
        {Flag: RF_RATE_LIMITED, Func: func(rf *v3.ResponseFlags) bool { return rf.GetRateLimited() }},
        {Flag: RF_UNAUTHORIZED_EXTERNAL_SERVICE, Func: func(rf *v3.ResponseFlags) bool { return rf.GetUnauthorizedDetails() != nil && rf.GetUnauthorizedDetails().GetReason() == v3.ResponseFlags_Unauthorized_EXTERNAL_SERVICE }},
        {Flag: RF_RATELIMIT_SERVICE_ERROR, Func: func(rf *v3.ResponseFlags) bool { return rf.GetRateLimitServiceError() }},
        {Flag: RF_DOWNSTREAM_CONNECTION_TERMINATION, Func: func(rf *v3.ResponseFlags) bool { return rf.GetDownstreamConnectionTermination() }},
        {Flag: RF_UPSTREAM_RETRY_LIMIT_EXCEEDED, Func: func(rf *v3.ResponseFlags) bool { return rf.GetUpstreamRetryLimitExceeded() }},
        {Flag: RF_STREAM_IDLE_TIMEOUT, Func: func(rf *v3.ResponseFlags) bool { return rf.GetStreamIdleTimeout() }},
        {Flag: RF_INVALID_ENVOY_REQUEST_HEADERS, Func: func(rf *v3.ResponseFlags) bool { return rf.GetInvalidEnvoyRequestHeaders() }},
        {Flag: RF_DOWNSTREAM_PROTOCOL_ERROR, Func: func(rf *v3.ResponseFlags) bool { return rf.GetDownstreamProtocolError() }},
        {Flag: RF_UPSTREAM_MAX_STREAM_DURATION_REACHED, Func: func(rf *v3.ResponseFlags) bool { return rf.GetUpstreamMaxStreamDurationReached() }},
        {Flag: RF_RESPONSE_FROM_CACHE_FILTER, Func: func(rf *v3.ResponseFlags) bool { return rf.GetResponseFromCacheFilter() }},
        {Flag: RF_NO_FILTER_CONFIG_FOUND, Func: func(rf *v3.ResponseFlags) bool { return rf.GetNoFilterConfigFound() }},
        {Flag: RF_DURATION_TIMEOUT, Func: func(rf *v3.ResponseFlags) bool { return rf.GetDurationTimeout() }},
        {Flag: RF_UPSTREAM_PROTOCOL_ERROR, Func: func(rf *v3.ResponseFlags) bool { return rf.GetUpstreamProtocolError() }},
        {Flag: RF_NO_CLUSTER_FOUND, Func: func(rf *v3.ResponseFlags) bool { return rf.GetNoClusterFound() }},
        {Flag: RF_OVERLOAD_MANAGER, Func: func(rf *v3.ResponseFlags) bool { return rf.GetOverloadManager() }},
    }
)

// adapt_ResponseFlags
func adapt_ResponseFlags(src *v3.ResponseFlags) string {
    if src == nil {
        return ""
    }

    dst_flags := []string{}
    for _, handler := range allResponseFlagHandlers {
        if handler.Func(src) {
            dst_flags = append(dst_flags, handler.Flag)
        }
    }

    return strings.Join(dst_flags, ",")
}
