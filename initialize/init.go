package initialize

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/google/uuid"
	"log"
)

var (
	ConsulClient	*consulapi.Client
	ServiceID 		string
	ServiceName 	string
	ServicePort 	int
)

//初始化
func init() {
	config := consulapi.DefaultConfig()
	config.Address = "192.168.1.104:8500" //虚拟机consul服务地址
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	ConsulClient = client
	ServiceID = "gokit:"+uuid.New().String()
}

//注册consul
func RegisterServer() {
	reg := consulapi.AgentServiceRegistration{}
	reg.ID = ServiceID	//"gokit01"
	reg.Name = ServiceName	//"gokitservice"
	reg.Address = "192.168.1.103" //localhost:8080对应的局域网地址192.168.1.103:8080
	reg.Port = ServicePort	//8080
	reg.Tags = []string{"primary"}

	check := consulapi.AgentServiceCheck{}
	check.Interval = "5s"
	//check.HTTP = "http://192.168.1.103:8080/health"
	check.HTTP = fmt.Sprintf("http://%s:%d/health",reg.Address,ServicePort)

	reg.Check = &check

	err := ConsulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}

}

//设置服务
func SetServiceNameAndPort(name string, port int){
	ServiceName = name
	ServicePort = port
}

//反注册consul服务
func UnregisterServer() {
	ConsulClient.Agent().ServiceDeregister(ServiceID)
}