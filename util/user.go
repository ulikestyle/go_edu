package util

import "edu/container"

func IsLogin() interface{} {

	return container.CreateContainersFactory().Get("isLogin")
}
