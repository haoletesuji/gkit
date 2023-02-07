package pkg

import (
	"context"

	userpb "grpcexp/proto"
)

func (srv *GrpcServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{
		Exist: true,
		User: &userpb.User{
			Id:   1,
			Name: "Andy",
		},
		UserId: 3,
	}, nil
}

func (srv *GrpcServer) GetUserIdBySession(ctx context.Context, req *userpb.GetUserIdBySessionRequest) (*userpb.GetUserIdBySessionResponse, error) {
	res := userpb.GetUserIdBySessionResponse{
		UserId: 3,
		Sid:    req.GetSid(),
		Email:  "xxx@gmail.com",
	}
	return &res, nil
}
