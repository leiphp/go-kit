package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	. "gokit/services"
	"net/http"
)

func main () {
	user := UserService{}
	endp := GenUserEndpoint(user)

	serverHandler := httptransport.NewServer(endp,DecodeUserRequest,EncodeUserRequest)
	http.ListenAndServe(":8080",serverHandler)

}