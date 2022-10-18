package gkit

type GrpcServer interface {
	Register()
	Run()
	GracefulStop()
}