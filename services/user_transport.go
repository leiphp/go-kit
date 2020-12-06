package services

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error){
	//http://127.0.0.1:8080/uid=101
	if r.URL.Query().Get("uid") != "" {
		uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
		return UserRequest{Uid:uid}, nil
	}
	return nil, errors.New("参数错误")
}

func EncodeUserRequest(c context.Context, w http.ResponseWriter, response interface{}) error{
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}