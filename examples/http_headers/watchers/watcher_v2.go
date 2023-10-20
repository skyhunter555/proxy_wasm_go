package watchers

import (
    envoy_service_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
    "github.com/golang/glog"
    "github.com/google/uuid"
    v2tov3 "github.com/tetratelabs/proxy-wasm-go-sdk/examples/http_headers/adapters/v2tov3"
    "io"
)

//
//
//

// AccessLogWatcherV2 Реализация grpc-сервиса для обработки access-логов v2
// Наследник `envoy_service_accesslog_v2.AccessLogServiceServer`
type AccessLogWatcherV2 struct {
    envoy_service_accesslog_v2.UnimplementedAccessLogServiceServer

    Adapter v2tov3.Adapter
    Handler AccessLogHandler
}

//
//
//

// NewWatcherV2 Фабрика для создания grpc-сервисов для прослушивания access-logs сообщений в формате v2
func NewWatcherV2(handler AccessLogHandler) envoy_service_accesslog_v2.AccessLogServiceServer {
    return NewWatcherV2Ex(v2tov3.NewAdapter(), handler)
}

// NewWatcherV2Ex Фабрика для создания grpc-сервисов для прослушивания access-logs сообщений в формате v2
func NewWatcherV2Ex(adapter v2tov3.Adapter, handler AccessLogHandler) envoy_service_accesslog_v2.AccessLogServiceServer {
    return &AccessLogWatcherV2{
        Adapter: adapter,
        Handler: handler,
    }
}

//
//
//

func (alw *AccessLogWatcherV2) StreamAccessLogs(stream envoy_service_accesslog_v2.AccessLogService_StreamAccessLogsServer) error {
    for {
        rquid := uuid.New().String()
        v2, err := stream.Recv()

        //
        glog.V(3).Infof("[%s] Got message V2: '%s'", rquid, v2)
        if err == io.EOF {
            glog.V(3).Infof("[%s] Close connection. End of stream", rquid)
            return nil
        }
        if err != nil {
            glog.Errorf("[%s] Close connection. Got error (%s)", rquid, err)
            return err
        }

        //
        go func(v2 *envoy_service_accesslog_v2.StreamAccessLogsMessage) {
            v3 := alw.Adapter.Adapt(v2)
            glog.V(3).Infof("[%s] Transformed from V2 to V3 successfully", rquid)
            glog.V(4).Infof("[%s] Transformed message: '%s'", rquid, v3)
            alw.Handler.Handle(rquid, v3)
        }(v2)
    }
}
