package services

import (
	"context"
	"encoding/json"
	"errors"
	"gokit/utils"
	"io/ioutil"
	"net/http"
	"github.com/tidwall/gjson"
)

func DecodeAccessRequest(c context.Context, r *http.Request) (interface{}, error){
	body,_ := ioutil.ReadAll(r.Body)
	result := gjson.Parse(string(body)) //第三方库解析json
	if result.IsObject() {
		username := result.Get("username")
		userpass := result.Get("userpass")
		return AccessRequest{Username:username.String(),Userpass:userpass.String(),Method:r.Method},nil
	}
	return nil, errors.New("参数错误")
}

func EncodeAccessRequest(c context.Context, w http.ResponseWriter, response interface{}) error{
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func MyErrorEncoder2(_ context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	w.Header().Set("Content-Type", contentType)
	if myerr, ok := err.(*utils.MyError) ;ok {
		w.WriteHeader(myerr.Code)
		w.Write(body)
	}else {
		w.WriteHeader(500)
		w.Write(body)
	}
}