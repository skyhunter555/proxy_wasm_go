package watchers

import (
    "context"
    "fmt"
    envoy_service_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
    envoy_service_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/suite"
    "google.golang.org/grpc"
    "google.golang.org/grpc/interop/grpc_testing"
    "io"
    "net"
    "sync"
    "testing"
    "time"
)

//
// Mocks
//

// AdapterV2Mock
type AdapterV2Mock struct {
    mock.Mock
}

func (av2m *AdapterV2Mock) Adapt(src *envoy_service_accesslog_v2.StreamAccessLogsMessage) *envoy_service_accesslog_v3.StreamAccessLogsMessage {
    args := av2m.MethodCalled("Adapt", src)
    return args.Get(0).(*envoy_service_accesslog_v3.StreamAccessLogsMessage)
}

// AccessLogHandlerV2Mock
type AccessLogHandlerV2Mock struct {
    mock.Mock
}

func (alhv2m *AccessLogHandlerV2Mock) Handle(rquid string, message *envoy_service_accesslog_v3.StreamAccessLogsMessage) {
    alhv2m.MethodCalled("Handle", rquid, message)
}

//
//
//

func Test001_NewWatcherV2(t *testing.T) {
    watcher := NewWatcherV2(nil)
    assert.NotNil(t, watcher, "Watcher must not be nil")
}

func Test002_NewWatcherV2Ex(t *testing.T) {
    watcher := NewWatcherV2Ex(nil, nil)
    assert.NotNil(t, watcher, "Watcher must not be nil")
}

//
//
//

func Test_WatcherV2Suite(t *testing.T) {
    ws := &WatcherV2Suite{port: 10100}
    suite.Run(t, ws)
}

type WatcherV2Suite struct {
    suite.Suite

    port  int
    mutex sync.RWMutex
}

func (ws *WatcherV2Suite) GetPort() int {
    ws.mutex.Lock()
    defer ws.mutex.Unlock()

    freeport := ws.port
    ws.port = ws.port + 1
    return freeport
}

func (ws *WatcherV2Suite) StartServerOn(port int, watcher envoy_service_accesslog_v2.AccessLogServiceServer) *grpc.Server {
    listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
    ws.Require().Nil(err, "Listener must be created")
    ws.Require().NotNil(listener, "Listener must not be nil")

    //
    server := grpc.NewServer()
    envoy_service_accesslog_v2.RegisterAccessLogServiceServer(server, watcher)
    go func() {
        err := server.Serve(listener)
        ws.Require().Nil(err, "Grpc server SERVE method must end without err: %s", err)
        listener.Close()
    }()

    //
    return server
}

func (ws *WatcherV2Suite) CreateClientFor(port int) (envoy_service_accesslog_v2.AccessLogServiceClient, *grpc.ClientConn) {
    connection, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
    ws.Require().Nil(err, "Connection error must be nil")
    ws.Require().NotNil(connection, "Connection must be created")
    return envoy_service_accesslog_v2.NewAccessLogServiceClient(connection), connection
}

func (ws *WatcherV2Suite) MakeConnection(port int, watcher envoy_service_accesslog_v2.AccessLogServiceServer) (*grpc.Server, envoy_service_accesslog_v2.AccessLogServiceClient, *grpc.ClientConn) {
    server := ws.StartServerOn(port, watcher)
    client, client_connection := ws.CreateClientFor(port)
    return server, client, client_connection
}

func (ws *WatcherV2Suite) Test100_AccessLogWatcherV2_Receive_EndOfStream_Should_Be_Ok_And_Close_Connection() {
    port := ws.GetPort()

    //
    adapter := &AdapterV2Mock{}
    handler := &AccessLogHandlerV2Mock{}
    watcher := NewWatcherV2Ex(adapter, handler)

    //
    server, client, _ := ws.MakeConnection(port, watcher)
    defer server.Stop()

    //
    ctx, _ := context.WithCancel(context.Background())
    stream, err := client.StreamAccessLogs(ctx)
    ws.Require().Nil(err, "Stream error must be nil")
    ws.Require().NotNil(stream, "Stream must be created")

    //
    response, err := stream.CloseAndRecv()
    ws.Require().Equal(io.EOF, err)
    ws.Require().Nil(response, "Close stream response must be nil")
}

func (ws *WatcherV2Suite) Test101_AccessLogWatcherV2_Receive_Message_Should_Be_Ok_And_Message_Must_Be_Sent_To_Handler() {
    port := ws.GetPort()

    //
    adapter := &AdapterV2Mock{}
    adapter.On("Adapt", mock.Anything).Return(&envoy_service_accesslog_v3.StreamAccessLogsMessage{}).Once()

    handler := &AccessLogHandlerV2Mock{}
    handler.On("Handle", mock.Anything, mock.Anything).Return().Once()

    watcher := NewWatcherV2Ex(adapter, handler)

    //
    server, client, _ := ws.MakeConnection(port, watcher)
    defer server.Stop()

    //
    ctx, _ := context.WithCancel(context.Background())
    stream, err := client.StreamAccessLogs(ctx)
    ws.Require().Nil(err, "Stream error must be nil")
    ws.Require().NotNil(stream, "Stream must be created")

    //
    empty := grpc_testing.Empty{}
    err = stream.SendMsg(&empty)
    ws.Require().Nil(err, "SendMsg err must be nil")

    //
    time.Sleep(time.Millisecond * 100)
    adapter.AssertExpectations(ws.T())
    handler.AssertExpectations(ws.T())

    //
    response, err := stream.CloseAndRecv()
    ws.Require().Equal(io.EOF, err)
    ws.Require().Nil(response, "CloseAndRecv response must be nil")
}
