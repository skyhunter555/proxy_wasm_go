package watchers

import (
    "github.com/golang/glog"

    envoy_service_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
)

//
//
//

// AccessLogHandler Обработчик access-logs сообщений
type AccessLogHandler interface {
    Handle(rquid string, message *envoy_service_accesslog_v3.StreamAccessLogsMessage)
}

//
//
//

// FakeAccessLogHandler
type FakeAccessLogHandler struct {

}

func (f *FakeAccessLogHandler) Handle(rquid string, message *envoy_service_accesslog_v3.StreamAccessLogsMessage) {
    glog.V(3).Infof("[%s] Handle message: %v",rquid,  message)
}

