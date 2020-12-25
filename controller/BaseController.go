package controller

import (
	"edu/container"
	"edu/util"
	"fmt"
)

type BaseInterface interface {
	Run()
}

type BaseController struct {

	ControllerName string
	ActionName string
}

func init() {
	container.CreateContainersFactory().Set("dispatcher", NewBaseController())
}

func NewBaseController() *BaseController{
	return &BaseController{}
}

func (router*BaseController)Run(){

	fmt.Println(router.ControllerName, router.ActionName)
	controllerObj := container.CreateContainersFactory().Get(router.ControllerName)
	reflectObj := container.CreateContainersFactory().Get("reflect").(*util.UserReflectStruct)
	reflectObj.SetModelKey(controllerObj)
	reflectObj.RunMethodByAllName(router.ActionName, nil)
}

