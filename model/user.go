package model

import (
	"strconv"
)

type User struct {
	username string	`Register:"请输入注册用户名" Login:"请输入用户名"`
	password string	`Register:"请输入注册密码" Login:"请输入注册密码"`
	age      int	`Register:"请输入年龄"`
	sex      string	`Register:"请输入性别"`
}

var UserDatas map[string]Model

func NewUser() *User {
	return &User{}
}

func (u *User) SetUsername(username string) {
	u.username = username
}
func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) SetAge(age int) {
	u.age = age
}

func (u *User) SetSex(sex string) {
	u.sex = sex
}

func (u *User) GetUsername() string {
	return u.username
}
func (u *User) GetPassword() string {
	return u.password
}
func (u *User) GetAge() int {
	return u.age
}
func (u *User) GetSex() string {
	return u.sex
}

func (u *User) RegisterToDo(username,password string,age int, sex string) string {
	return u.username + "," + u.password + "," + strconv.Itoa(u.age) + "," + u.sex
}

//func (u *User) LoginToDo(username string,password string) bool {
//
//	userModel := u.GetUserModelByName(username)
//
//	if userModel == nil {
//		return false
//	}
//
//	if userModel.GetPassword() != password {
//		return false
//	}
//
//	return true
//}

// 格式化输出数据信息
func (u *User) ToString() string {
	return u.username + "," + u.password + "," + strconv.Itoa(u.age) + "," + u.sex
}

func (u *User) GetDatas() map[string]Model{
	return UserDatas
}

func (u *User) GetUserModelByName(username string) (*User, bool){

	if data, ok := UserDatas[username]; ok {
		return data.(*User),true
	} else {
		return &User{},false
	}
}
