package main

import (
	"flag"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"gokit/initialize"
	. "gokit/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main () {
	//服务参数获取
	name := flag.String("name","","服务名称")
	port := flag.Int("p",0,"服务端口")
	flag.Parse()
	if *name == "" {
		log.Fatal("请指定服务名")
	}
	if *port == 0 {
		log.Fatal("请指定端口")
	}
	initialize.SetServiceNameAndPort(*name, *port) //设置服务名和端口

	user := UserService{}
	endp := GenUserEndpoint(user)

	serverHandler := httptransport.NewServer(endp,DecodeUserRequest,EncodeUserRequest)

	router := mymux.NewRouter()
	//r.Handle(`/user/{uid:\d+}`,serverHandler)
	router.Methods("GET","DELETE").Path(`/user/{uid:\d+}`).Handler(serverHandler)
	router.Methods("GET").Path("/health").HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
		write.Header().Set("Content-Type", "application/json")
		write.Write([]byte(`{"status":"ok"}`))
	})

	errChan := make(chan error)
	go func() {
		//注册consul服务
		initialize.RegisterServer()
		err := http.ListenAndServe(":"+strconv.Itoa(*port),router)
		if err != nil {
			log.Println(err)
			errChan <- err
		}
	}()

	go func() {
		sig_c := make(chan os.Signal)
		signal.Notify(sig_c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s",<-sig_c)
	}()

	//如果没有异常错误，errChan将永久阻塞
	getErr := <- errChan
	initialize.UnregisterServer()
	log.Println(getErr)
}