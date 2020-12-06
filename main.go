package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	. "gokit/services"
	"net/http"
)

func main () {
	user := UserService{}
	endp := GenUserEndpoint(user)

	serverHandler := httptransport.NewServer(endp,DecodeUserRequest,EncodeUserRequest)

	router := mymux.NewRouter()
	//r.Handle(`/user/{uid:\d+}`,serverHandler)
	router.Methods("GET","DELETE").Path(`/user/{uid:\d+}`).Handler(serverHandler)
	http.ListenAndServe(":8080",router)

}