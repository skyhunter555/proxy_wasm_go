package v3to_vjcal

import (
    "fmt"
    "regexp"
    "testing"

    v3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
    "github.com/stretchr/testify/assert"
)

func Test_adapt_ResponseFlags(t *testing.T) {
    type args struct {
        src *v3.ResponseFlags
    }
    tests := []struct {
        name string
        args args
        want []string
    }{
        {
            name: "Test NIL",
            args: args{
                src: nil,
            },
            want: []string{},
        },
        {
            name: "Test OBJECT",
            args: args{
                src: &v3.ResponseFlags{
                    FailedLocalHealthcheck:           true,
                    NoHealthyUpstream:                true,
                    UpstreamRequestTimeout:           true,
                    LocalReset:                       true,
                    UpstreamRemoteReset:              true,
                    UpstreamConnectionFailure:        true,
                    UpstreamConnectionTermination:    true,
                    UpstreamOverflow:                 true,
                    NoRouteFound:                     true,
                    DelayInjected:                    true,
                    FaultInjected:                    true,
                    RateLimited:                      true,
                    UnauthorizedDetails:              &v3.ResponseFlags_Unauthorized{Reason: v3.ResponseFlags_Unauthorized_EXTERNAL_SERVICE},
                    RateLimitServiceError:            true,
                    DownstreamConnectionTermination:  true,
                    UpstreamRetryLimitExceeded:       true,
                    StreamIdleTimeout:                true,
                    InvalidEnvoyRequestHeaders:       true,
                    DownstreamProtocolError:          true,
                    UpstreamMaxStreamDurationReached: true,
                    ResponseFromCacheFilter:          true,
                    NoFilterConfigFound:              true,
                    DurationTimeout:                  true,
                    UpstreamProtocolError:            true,
                    NoClusterFound:                   true,
                    OverloadManager:                  true,
                },
            },
            want: []string{
                RF_FAILED_LOCAL_HEALTH_CHECK,
                RF_NO_HEALTHY_UPSTREAM,
                RF_UPSTREAM_REQUEST_TIMEOUT,
                RF_LOCAL_RESET,
                RF_UPSTREAM_REMOTE_RESET,
                RF_UPSTREAM_CONNECTION_FAILURE,
                RF_UPSTREAM_CONNECTION_TERMINATION,
                RF_UPSTREAM_OVERFLOW,
                RF_NO_ROUTE_FOUND,
                RF_DELAY_INJECTED,
                RF_FAULT_INJECTED,
                RF_RATE_LIMITED,
                RF_UNAUTHORIZED_EXTERNAL_SERVICE,
                RF_RATELIMIT_SERVICE_ERROR,
                RF_DOWNSTREAM_CONNECTION_TERMINATION,
                RF_UPSTREAM_RETRY_LIMIT_EXCEEDED,
                RF_STREAM_IDLE_TIMEOUT,
                RF_INVALID_ENVOY_REQUEST_HEADERS,
                RF_DOWNSTREAM_PROTOCOL_ERROR,
                RF_UPSTREAM_MAX_STREAM_DURATION_REACHED,
                RF_RESPONSE_FROM_CACHE_FILTER,
                RF_NO_FILTER_CONFIG_FOUND,
                RF_DURATION_TIMEOUT,
                RF_UPSTREAM_PROTOCOL_ERROR,
                RF_NO_CLUSTER_FOUND,
                RF_OVERLOAD_MANAGER,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := adapt_ResponseFlags(tt.args.src)
            for _, flag := range tt.want {
                matched, err := regexp.MatchString(fmt.Sprintf("(^|,)%s($|,)", flag), got)
                assert.Nil(t, err, "Error must be nil")
                assert.Truef(t, matched, "Flag '%s' must be exists in response: '%s'", flag, got)
            }
        })
    }
}
