package controller

import (
	"edu/container"
	"edu/util"
	"fmt"
	"os"
)

type IndexController struct {

}

func init(){
	container.CreateContainersFactory().Set("Index", NewIndexController())
}

func NewIndexController() *IndexController{
	return &IndexController{}
}

func (index *IndexController)Index(){

	//for {

		fmt.Println("############开始########")

		isLogin := util.IsLogin()

		router := container.CreateContainersFactory().Get("dispatcher").(*BaseController)

		if isLogin == "" {

			fmt.Println("(0)	: 注册用户")
			fmt.Println("(1)	: 登入系统")
			fmt.Println("退出请输入 x或其他退出")

			// 用户输入的选择
			inputStr := util.GetInput()

			router.ControllerName = "User"
			if inputStr == "0" {
				router.ActionName = "Register"
			} else if inputStr == "1" {
				router.ActionName = "Login"
			} else {
				os.Exit(100)
			}

		}else{

			fmt.Println("(0)	: 展示用户信息")
			fmt.Println("退出请输入 x或其他退出")

			inputStr := util.GetInput()
			if inputStr == "0" {
				router := container.CreateContainersFactory().Get("dispatcher").(*BaseController)
				router.ControllerName = "User"
				router.ActionName = "Show"
			}else {
				os.Exit(100)
			}
		}
		fmt.Println("############结束########")

		fmt.Println("")

		fmt.Println("############开始########")
		controllerObj := container.CreateContainersFactory().Get(router.ControllerName)
		reflectObj := container.CreateContainersFactory().Get("reflect").(*util.UserReflectStruct)
		reflectObj.SetModelKey(controllerObj)
		reflectObj.RunMethodByAllName(router.ActionName, nil)
		fmt.Println("############结束########")

	//}

	//fmt.Println("Bye!!")
}