package initialize

import (
	consulapi "github.com/hashicorp/consul/api"
	"log"
)

var (
	ConsulClient	*consulapi.Client
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
}

//注册consul
func RegisterServer() {
	reg := consulapi.AgentServiceRegistration{}
	reg.ID = "gokit01"
	reg.Name = "gokitservice"
	reg.Address = "192.168.1.103" //localhost:8080对于的局域网地址192.168.1.103:8080
	reg.Port = 8080
	reg.Tags = []string{"primary"}

	check := consulapi.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = "http://192.168.1.103:8080/health"

	reg.Check = &check

	err := ConsulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}

}

//反注册consul服务
func UnregisterServer() {
	ConsulClient.Agent().ServiceDeregister("gokit01")
}