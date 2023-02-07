package pkg

import (
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	gkit "gkit"
	"google.golang.org/grpc"
	userpb "grpcexp/proto"
	"net"
)

type GrpcServer struct {
	grpcPort string
	logger   gkit.GrpcLogger
	s        *grpc.Server
}

func NewGrpcServer(logger gkit.GrpcLogger) gkit.GrpcServer {
	srv := &GrpcServer{
		grpcPort: "50001",
		logger:   logger,
	}
	srv.s = gkit.InitializeGrpcServer(srv.logger)
	return srv
}

func (srv *GrpcServer) Register() {
	userpb.RegisterUserServiceServer(srv.s, srv)
	grpc_prometheus.Register(srv.s)
}

func (srv *GrpcServer) Run() {
	go func() {
		addr := "0.0.0.0:" + srv.grpcPort
		srv.logger.Infoln("grpc server listening on  ", addr)
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			srv.logger.Fatal(err)
		}
		if err := srv.s.Serve(lis); err != nil {
			srv.logger.Fatal(err)
		}
	}()
}

func (srv *GrpcServer) GracefulShutdown() {
	srv.s.GracefulStop()
}
