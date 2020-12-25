package main

import (
	"edu/container"
	"edu/controller"
	"fmt"
)

func init() {
	container.CreateContainersFactory().Set("isLogin", "")
}

/**
 思路：

避免重复实例化和NewUser()操作，从其他项目借鉴了 container思路，保存在容器中;
- userModel ,对于*User
- isLogin，存放的是用户登录标识（已登录，存放的是用户的唯一标识，在这里使用username）

view层实现思路，已实现
1. 利用标识符(controllerName, actionName),
- controllerName 为controller的struct结构体名称
- actionName 为所需执行的方法
2. 反射拿到value， call调用方法

 */
func main() {

	fmt.Println("欢迎来到学员管理系统 >>>")
	fmt.Println("你的执行操作:")

	AutoController := container.CreateContainersFactory().Get("Auto").(*controller.AutoController)
	AutoController.Run()

}
