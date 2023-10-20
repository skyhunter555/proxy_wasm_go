package v3to_vjcal

import (
    "bitbucket.region.vtb.ru/projects/USBP/repos/envoy-accesslogs/internal/vtbjournalclient_accesslogs"
    envoy_service_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
)

//
//
//

// Adapter
type Adapter interface {
    Adapt(src *envoy_service_accesslog_v3.StreamAccessLogsMessage) []*vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog
}

// NewAdapter
func NewAdapter() Adapter {
    return &adapterImpl{}
}

//
//
//

// adapterImpl
type adapterImpl struct {
}

// Adapt
func (a *adapterImpl) Adapt(src *envoy_service_accesslog_v3.StreamAccessLogsMessage) []*vtbjournalclient_accesslogs.VtbJournalExtContextAgrAccessLog {
    return adapt_Message(src)
}
