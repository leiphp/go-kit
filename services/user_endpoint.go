package services

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"gokit/initialize"
	"gokit/utils"
	"golang.org/x/time/rate"
	"strconv"
)

type UserRequest struct {
	Uid int `json:"uid"`
	Method string
	Token  string
}

type UserResponse struct {
	Result string 	`json:"result"`
}

//token验证中间件
func CheckTokenMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			r := request.(UserRequest)
			uc := UserClaim{}
			getToken,err := jwt.ParseWithClaims(r.Token,&uc,func(token *jwt.Token) (i interface{},e error){
				return []byte(secKey), nil
			})
			if getToken != nil && getToken.Valid {//验证通过
				newCtx := context.WithValue(ctx,"LoginUser",getToken.Claims.(*UserClaim).Uname)
				return next(newCtx, request)
			}else {
				return nil, utils.NewMyError(403, "error token")
			}
		}
	}
}

//日志中间件
func UserServiceLogMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			r := request.(UserRequest)
			logger.Log("method",r.Method,"event","get user","userId",r.Uid)
			return next(ctx, request)
		}
	}
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
		r := request.(UserRequest)
		fmt.Println("当前登录用户名是：",ctx.Value("LoginUser"))
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
