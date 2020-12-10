package services

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"gokit/initialize"
	"strconv"
)

type UserRequest struct {
	Uid int `json:"uid"`
	Method string
}

type UserResponse struct {
	Result string 	`json:"result"`
}

func GenUserEndpoint( service UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		var result string
		if r.Method == "GET" {
			result = service.GetName(r.Uid)+strconv.Itoa(initialize.ServicePort)
		}else if r.Method == "DELETE" {
			err := service.DelUser(r.Uid)
			if err != nil {
				result = err.Error()
			}else {
				result = fmt.Sprintf("userId为%d的用户删除成功",r.Uid)
			}
		}

		return UserResponse{Result:result}, nil
	}
}
