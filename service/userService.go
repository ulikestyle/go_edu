package service

import (
	"edu/model"
	"fmt"
)

type serviceInterface interface {
	//查询,是否存在

	//获取信息

	//根据username查询用户信息
	FindOne(username string) model.User

	//保存用户信息
	SaveInfo(user model.User) bool
}

type UserService struct {
	userModel *model.User
}

func init()  {
	//UserService{userModel: model.NewUser()}

	fmt.Print("UserService init ......")
}

func (uService *UserService) FindOne(username string) model.Model {

	allData := uService.GetUserDatasFromCache()
	if data, ok := allData[username]; ok {
		return data
	}else{
		return nil
	}
}

func (uService *UserService) GetUserDatasFromCache() map[string]model.Model {
	return uService.userModel.GetDatas()
}