package controller

import (
	"edu/container"
	"edu/model"
	"edu/util"
	"fmt"
	"strconv"
)

type UserController struct {

}

func init(){
	container.CreateContainersFactory().Set("User", NewUserController())
}

func NewUserController() *UserController{
	return &UserController{}
}

func (User *UserController)Login() {

	username := util.GetInput("请输入用户名")
	password := util.GetInput("请输入密码")

	fmt.Println("用户输入的是 ", username, "长度: ", len(username), password)

	userModel := container.CreateContainersFactory().Get("userModel").(*model.User)
	userData,exist := userModel.GetUserModelByName(username)

	router := container.CreateContainersFactory().Get("dispatcher").(*BaseController)
	router.ControllerName = "Index"
	router.ActionName = "Index"

	if !exist {
		fmt.Println("登录失败！ 用户名或密码错误")
	}else {
		if userData.GetPassword() == password {

			fmt.Println("登录成功")
			container.CreateContainersFactory().Delete("isLogin")
			container.CreateContainersFactory().Set("isLogin", username)
		}
	}
}

func (User *UserController)Register() {
	username := util.GetInput("请输入用户名")
	password := util.GetInput("请输入密码")
	age := util.GetInput("请输入年龄")
	age2,_:= strconv.Atoi(age)
	sex := util.GetInput("请输入性别")

	regUserModel := model.NewUser()

	regUserModel.SetUsername(username)
	regUserModel.SetPassword(password)
	regUserModel.SetAge(age2)
	regUserModel.SetSex(sex)

	regUserModel.ToString()


}

func (User *UserController)Show() {

	username := container.CreateContainersFactory().Get("isLogin").(string)

	userModel := container.CreateContainersFactory().Get("userModel").(*model.User)
	userData,_ := userModel.GetUserModelByName(username)

	fmt.Println("当前用户信息是： ", userData.ToString())
}