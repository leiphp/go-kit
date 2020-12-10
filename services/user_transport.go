package services

import (
	"context"
	"encoding/json"
	"errors"
	mymux "github.com/gorilla/mux"
	"gokit/utils"
	"net/http"
	"strconv"
)

func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error){
	vars := mymux.Vars(r)
	if uid,ok := vars["uid"]; ok {
		uid, _ := strconv.Atoi(uid)
		return UserRequest{
			Uid:uid,
			Method:r.Method,
		}, nil
	}
	return nil, errors.New("参数错误")
}

func EncodeUserRequest(c context.Context, w http.ResponseWriter, response interface{}) error{
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func MyErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
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