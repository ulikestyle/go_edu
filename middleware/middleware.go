package middleware

type eduMiddleware interface {
	next(controllerName string, actionName string)
}

// 标识和下一步
func next() bool{

}