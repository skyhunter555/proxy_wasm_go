package watchers

import (
    "github.com/google/uuid"
    "io"

    envoy_service_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
    "github.com/golang/glog"
)

//
//
//

// AccessLogWatcherV3 Реализация grpc-сервиса для обработки access-логов v3
// Наследник `envoy_service_accesslog_v3.AccessLogServiceServer`
type AccessLogWatcherV3 struct {
    envoy_service_accesslog_v3.UnimplementedAccessLogServiceServer

    Handler AccessLogHandler
}

//
//
//

// NewWatcherV3 Фабрика для создания grpc-сервисов для прослушивания access-logs сообщений в формате v3
func NewWatcherV3(handler AccessLogHandler) envoy_service_accesslog_v3.AccessLogServiceServer {
    return &AccessLogWatcherV3{
        Handler: handler,
    }
}

//
//
//

func (alw *AccessLogWatcherV3) StreamAccessLogs(stream envoy_service_accesslog_v3.AccessLogService_StreamAccessLogsServer) error {
    for {
        rquid := uuid.New().String()
        v3, err := stream.Recv()

        //
        glog.V(3).Infof("[%s] Got message V3 '%s'", rquid, v3)
        if err == io.EOF {
            glog.V(3).Infof("[%s] End of stream. Close connection", rquid)
            return nil
        }
        if err != nil {
            glog.Errorf("[%s] Got error. Close connection (%s)", rquid, err)
            return err
        }

        //
        go alw.Handler.Handle(rquid, v3)
    }
}

