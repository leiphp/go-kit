package services

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	Uid int `json:"uid"`
}

type UserResponse struct {
	Result string 	`json:"result"`
}

func GenUserEndpoint( service UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		result := service.GetName(r.Uid)
		return UserResponse{Result:result}, nil
	}
}
