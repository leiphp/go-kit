package models

type User struct {
	ID           int     `gorm:"primary_key" json:"id"` //id
	Name   		 string  `json:"name"`           		//名称
	Age			 int64   `json:"age"`        			//年龄
	UserId       int64   `json:"user_id"`               //用户id
	CreateTime   int64   `json:"create_time"`           //时间
	Remark       string  `json:"remark"`                //备注
	Status       int     `json:"status"`                //状态
}
