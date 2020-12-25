package model

import (
	"bufio"
	"edu/controller"
	"fmt"
	"os"
	"testing"
)

func TestRfdata(t *testing.T) {

	//userDatas = make(map[string]Model, 0)
	//rfdata("user", "username", userDatas)
	//fmt.Println(userDatas)
}

func TestNewUser(t *testing.T) {

	// 从标准输入流中接收输入数据
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {

		inputStr := input.Text()

		fmt.Printf("输入的字符串是%s, %T \n",inputStr, inputStr)

		// 输入bye时 结束
		if inputStr == "x" {
			fmt.Printf(">>>>>>>>>>>>> 已退出 bye！")
			break
		}

		if inputStr == "0" {
			handler := controller.LoginController{}
			handler.ViewLayer("user")
		}

		fmt.Println("##########End#########")
	}
}
