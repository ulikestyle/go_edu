package controller

import (
	"edu/container"
	"fmt"
)

type AutoController struct {

}

func init(){
	container.CreateContainersFactory().Set("Auto", NewAutoController())
}

func NewAutoController() *AutoController{
	return &AutoController{}
}

func (auto *AutoController)Run(){

	router := container.CreateContainersFactory().Get("dispatcher").(*BaseController)
	router.ControllerName = "Index"
	router.ActionName = "Index"

	for {
		fmt.Printf("当前位置 %s - %s \n", router.ControllerName, router.ActionName)
		router.Run()
	}
}