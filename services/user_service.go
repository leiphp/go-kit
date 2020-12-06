package services

import "gokit/models"

//存放xx模块的基础功能接口定义及实现
type UserInterfaceService interface {
	AddUser(query models.User) (interface{}, error)
	GetName(userId int) string
}

//初始化对象函数
func NewUserService() UserInterfaceService {
	return &userService{

	}
}

type userService struct {

}


type UserService struct {

}
func (this *UserService) GetName(userId int) string {
	if userId == 101 {
		return "leixiaotain"
	}
	return "guest"
}

//添加用户
func (this *userService) AddUser(query models.User) (interface{}, error){
	return "hello world", nil
}

func (this *userService) GetName(userId int) string {
	if userId == 101 {
		return "leixiaotain"
	}
	return "guest"
}