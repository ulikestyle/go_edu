package util

import "fmt"

func Debug(objName string, obj interface{}){

	fmt.Println("########################", objName)

	fmt.Printf("%s 值是： %v \n", objName, obj)
	fmt.Printf("%s 类型是： %T \n", objName, obj)
	fmt.Printf("%s 指针是： %p \n", objName, obj)

	fmt.Println("########################", objName)
	fmt.Println()
}
