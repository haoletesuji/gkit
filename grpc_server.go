package common

type GrpcServer interface {
	Register()
	Run()
	GracefulStop()
}