package grpc

import (
	"context"
	fushinRPC "github.com/JJApplication/fushin/grpc"
	"github.com/JJApplication/fushin/log"
	"google.golang.org/grpc"
	"math"
	"net"
	"time"
)

type Server struct {
	fushinRPC.UnimplementedFushinGRPCServiceServer
	logger      *log.Logger
	grpcServer  *grpc.Server
	serviceName string // 注册服务名
}

func NewServer(logger *log.Logger, serviceName string) *Server {
	s := &Server{
		serviceName: serviceName,
		logger:      logger,
	}
	s.grpcServer = grpc.NewServer(
		grpc.MaxRecvMsgSize(math.MaxInt32),
		grpc.MaxSendMsgSize(math.MaxInt32),
		grpc.MaxConcurrentStreams(math.MaxInt32),
	)
	fushinRPC.RegisterFushinGRPCServiceServer(s.grpcServer, s)

	return s
}

// ServeUDS 以UDS启动GRPC服务器
func (s *Server) ServeUDS(address string) error {
	s.logger.InfoF("Starting gRPC server on UDS: %s", address)
	addr, err := net.ResolveUnixAddr("unix", address)
	if err != nil {
		return err
	}
	listener, err := net.ListenUnix("unix", addr)
	if err != nil {
		return err
	}
	return s.grpcServer.Serve(listener)
}

// Serve 启动gRPC服务器
func (s *Server) Serve(lis net.Listener) error {
	s.logger.InfoF("Starting gRPC server: %s", lis.Addr().String())
	return s.grpcServer.Serve(lis)
}

// GracefulStop 优雅停止gRPC服务器
func (s *Server) GracefulStop() {
	s.grpcServer.GracefulStop()
}

func (s *Server) HealthCheck(ctx context.Context, req *fushinRPC.FushinRequest) (*fushinRPC.FushinResponse, error) {
	return &fushinRPC.FushinResponse{
		Service:   s.serviceName,
		Receiver:  "",
		Timestamp: time.Now().Unix(),
	}, nil
}
