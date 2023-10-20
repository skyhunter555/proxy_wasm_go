package v2tov3

import (
    envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
    envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
    envoy_data_accesslog_src "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v2"
    envoy_data_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
)

//
// Log Entries
// #################################################################################################################

//
// Common
// #################################################################################################################

func adapt_AccessLogCommon(src *envoy_data_accesslog_src.AccessLogCommon) *envoy_data_accesslog_v3.AccessLogCommon {
    if src == nil {
        return nil
    }

    return &envoy_data_accesslog_v3.AccessLogCommon{
        SampleRate:                     src.GetSampleRate(),
        DownstreamRemoteAddress:        adapt_Address(src.GetDownstreamRemoteAddress()),
        DownstreamLocalAddress:         adapt_Address(src.GetDownstreamLocalAddress()),
        TlsProperties:                  adapt_TLSProperties(src.GetTlsProperties()),
        StartTime:                      src.GetStartTime(),
        TimeToLastRxByte:               src.GetTimeToLastRxByte(),
        TimeToFirstUpstreamTxByte:      src.GetTimeToFirstUpstreamTxByte(),
        TimeToLastUpstreamTxByte:       src.GetTimeToLastUpstreamTxByte(),
        TimeToFirstUpstreamRxByte:      src.GetTimeToFirstUpstreamRxByte(),
        TimeToLastUpstreamRxByte:       src.GetTimeToLastUpstreamRxByte(),
        TimeToFirstDownstreamTxByte:    src.GetTimeToFirstDownstreamTxByte(),
        TimeToLastDownstreamTxByte:     src.GetTimeToLastDownstreamTxByte(),
        UpstreamRemoteAddress:          adapt_Address(src.GetUpstreamRemoteAddress()),
        UpstreamLocalAddress:           adapt_Address(src.GetUpstreamLocalAddress()),
        UpstreamCluster:                src.GetUpstreamCluster(),
        ResponseFlags:                  adapt_ResponseFlags(src.GetResponseFlags()),
        Metadata:                       adapt_Metadata(src.GetMetadata()),
        UpstreamTransportFailureReason: src.GetUpstreamTransportFailureReason(),
        RouteName:                      src.GetRouteName(),
        DownstreamDirectRemoteAddress:  adapt_Address(src.GetDownstreamDirectRemoteAddress()),
        FilterStateObjects:             src.GetFilterStateObjects(),
    }
}

func adapt_TLSProperties(src *envoy_data_accesslog_src.TLSProperties) *envoy_data_accesslog_v3.TLSProperties {
    if src == nil {
        return nil
    }

    return &envoy_data_accesslog_v3.TLSProperties{
        TlsVersion:                 envoy_data_accesslog_v3.TLSProperties_TLSVersion(src.GetTlsVersion()),
        TlsCipherSuite:             src.GetTlsCipherSuite(),
        TlsSniHostname:             src.GetTlsSniHostname(),
        LocalCertificateProperties: adapt_TLSProperties_CertificateProperties(src.GetLocalCertificateProperties()),
        PeerCertificateProperties:  adapt_TLSProperties_CertificateProperties(src.GetPeerCertificateProperties()),
        TlsSessionId:               src.GetTlsSessionId(),
    }
}

func adapt_TLSProperties_CertificateProperties(src *envoy_data_accesslog_src.TLSProperties_CertificateProperties) *envoy_data_accesslog_v3.TLSProperties_CertificateProperties {
    if src == nil {
        return nil
    }

    return &envoy_data_accesslog_v3.TLSProperties_CertificateProperties{
        SubjectAltName: adapt_TLSProperties_CertificateProperties_SubjectAltNames(src.GetSubjectAltName()),
        Subject:        src.GetSubject(),
    }
}

func adapt_TLSProperties_CertificateProperties_SubjectAltNames(src_array []*envoy_data_accesslog_src.TLSProperties_CertificateProperties_SubjectAltName) []*envoy_data_accesslog_v3.TLSProperties_CertificateProperties_SubjectAltName {
    if src_array == nil {
        return nil
    }

    dst_array := make([]*envoy_data_accesslog_v3.TLSProperties_CertificateProperties_SubjectAltName, 0)
    for _, src := range src_array {
        dst_array = append(dst_array, adapt_TLSProperties_CertificateProperties_SubjectAltName(src))
    }

    return dst_array
}

func adapt_TLSProperties_CertificateProperties_SubjectAltName(src *envoy_data_accesslog_src.TLSProperties_CertificateProperties_SubjectAltName) *envoy_data_accesslog_v3.TLSProperties_CertificateProperties_SubjectAltName {
    if src == nil {
        return nil
    }

    if src_uri := src.GetUri(); src_uri != "" {
        return &envoy_data_accesslog_v3.TLSProperties_CertificateProperties_SubjectAltName{
            San: &envoy_data_accesslog_v3.TLSProperties_CertificateProperties_SubjectAltName_Uri{
                Uri: src_uri,
            },
        }
    }

    if src_dns := src.GetDns(); src_dns != "" {
        return &envoy_data_accesslog_v3.TLSProperties_CertificateProperties_SubjectAltName{
            San: &envoy_data_accesslog_v3.TLSProperties_CertificateProperties_SubjectAltName_Dns{
                Dns: src_dns,
            },
        }
    }

    return &envoy_data_accesslog_v3.TLSProperties_CertificateProperties_SubjectAltName{}
}

// adapt_ResponseFlags
func adapt_ResponseFlags(src *envoy_data_accesslog_src.ResponseFlags) *envoy_data_accesslog_v3.ResponseFlags {
    if src == nil {
        return nil
    }

    return &envoy_data_accesslog_v3.ResponseFlags{
        FailedLocalHealthcheck:           src.GetFailedLocalHealthcheck(),
        NoHealthyUpstream:                src.GetNoHealthyUpstream(),
        UpstreamRequestTimeout:           src.GetUpstreamRequestTimeout(),
        LocalReset:                       src.GetLocalReset(),
        UpstreamRemoteReset:              src.GetUpstreamRemoteReset(),
        UpstreamConnectionFailure:        src.GetUpstreamConnectionFailure(),
        UpstreamConnectionTermination:    src.GetUpstreamConnectionTermination(),
        UpstreamOverflow:                 src.GetUpstreamOverflow(),
        NoRouteFound:                     src.GetNoRouteFound(),
        DelayInjected:                    src.GetDelayInjected(),
        FaultInjected:                    src.GetFaultInjected(),
        RateLimited:                      src.GetRateLimited(),
        UnauthorizedDetails:              adapt_ResponseFlags_Unauthorized(src.GetUnauthorizedDetails()),
        RateLimitServiceError:            src.GetRateLimitServiceError(),
        DownstreamConnectionTermination:  src.GetDownstreamConnectionTermination(),
        UpstreamRetryLimitExceeded:       src.GetUpstreamRetryLimitExceeded(),
        StreamIdleTimeout:                src.GetStreamIdleTimeout(),
        InvalidEnvoyRequestHeaders:       src.GetInvalidEnvoyRequestHeaders(),
        DownstreamProtocolError:          src.GetDownstreamProtocolError(),
        UpstreamMaxStreamDurationReached: false,
        ResponseFromCacheFilter:          false,
        NoFilterConfigFound:              false,
        DurationTimeout:                  false,
    }
}

// adapt_ResponseFlags_Unauthorized
func adapt_ResponseFlags_Unauthorized(src *envoy_data_accesslog_src.ResponseFlags_Unauthorized) *envoy_data_accesslog_v3.ResponseFlags_Unauthorized {
    if src == nil {
        return nil
    }

    return &envoy_data_accesslog_v3.ResponseFlags_Unauthorized{
        Reason: envoy_data_accesslog_v3.ResponseFlags_Unauthorized_Reason(src.GetReason()),
    }
}

// adapt_Metadata
func adapt_Metadata(src *envoy_api_v2_core.Metadata) *envoy_config_core_v3.Metadata {
    if src == nil {
        return nil
    }

    return &envoy_config_core_v3.Metadata{
        FilterMetadata: src.GetFilterMetadata(),
    }
}
