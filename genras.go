package main

import (
	"gokit/utils"
	"log"
)

//生成私钥和公钥时，先注释main文件中的main方法，两个main方法会报错
func main2() {
	err := utils.GenRSAPubAndPri(1024,"./pem")
	if err !=nil {
		log.Fatal(err)
	}
}
