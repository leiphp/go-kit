package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func GenRSAPubAndPri(bits int, filepath string) error {
	//生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock := &pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   derStream,
	}

	err = ioutil.WriteFile(filepath+"/private.pem",pem.EncodeToMemory(priBlock),0644)
	if err != nil {
		return err
	}
	fmt.Println("==============私钥文件创建成功==============")
	//生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil
	}
	publicBlock := &pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   derPkix,
	}
	err = ioutil.WriteFile(filepath+"/public.pem",pem.EncodeToMemory(publicBlock),0644)
	if err != nil {
		return err
	}
	fmt.Println("==============公钥文件创建成功==============")
	return nil
}
