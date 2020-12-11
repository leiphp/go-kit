package services

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"gokit/initialize"
	"gokit/utils"
	"golang.org/x/time/rate"
	"os"
	"strconv"
)

type UserRequest struct {
	Uid int `json:"uid"`
	Method string
}

type UserResponse struct {
	Result string 	`json:"result"`
}

//假如限流功能的中间件
func RateLimit(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				//return nil, errors.New("too many requests")
				return nil, utils.NewMyError(429, "too many requests")
			}
			return next(ctx, request)
		}
	}
}

func GenUserEndpoint( service UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var logger log.Logger
		{
			logger = log.NewLogfmtLogger(os.Stdout)
			logger = log.WithPrefix(logger,"gokit","1.0")
			logger = log.With(logger,"time",log.DefaultTimestampUTC)
			logger = log.With(logger,"caller",log.DefaultCaller)
		}
		r := request.(UserRequest)
		var result string
		if r.Method == "GET" {
			result = service.GetName(r.Uid)+strconv.Itoa(initialize.ServicePort)
			logger.Log("method",r.Method,"event","get user","userId",r.Uid)
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
