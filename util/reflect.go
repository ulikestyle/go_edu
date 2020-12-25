package util

import (
	"edu/container"
	"reflect"
	"strings"
)


type reflectInterface interface {
	GetReflectType(i interface{}) reflect.Type
	GetReflectValue(i interface{}) reflect.Value
}

func init() {
	container.CreateContainersFactory().Set("reflect", NewReflect())
}

func NewReflect() *UserReflectStruct{
	return &UserReflectStruct{}
}

type UserReflectStruct struct {
	pk string
	model interface{}
	fieldStructArr []FieldStruct
}

// 解析user体后，出来的结构化数据
type FieldStruct struct {
	Sname string // 字段名称
	Stype string //字段类型
	Stag map[string]string // 字段tag
}

// 获取反射的动态类型
func (urs *UserReflectStruct) GetReflectType() reflect.Type{
	i := urs.model
	reType := reflect.TypeOf(i)

	if reType.Kind().String() != "struct" {

		reValue := reflect.ValueOf(i)
		//fmt.Println("当前传值不是struct, 采用reValue.Elem().Type()进行转化")
		reType = reValue.Elem().Type()

		//fmt.Println("当前传值不是struct, 采用reValue.Elem().Type()进行转化")
		//reType = reflect.Indirect(reValue)
	}
	return reType
}

// 获取反射的值
func (urs *UserReflectStruct) GetReflectValue() reflect.Value{

	i := urs.model
	return reflect.ValueOf(i)
}

func (urs *UserReflectStruct) GetModelFields() []FieldStruct {

	reType := urs.GetReflectType()

	fieldStructList := make([]FieldStruct,reType.NumField())
	for i :=0;i < reType.NumField(); i++ {

		reTypeField := reType.Field(i)

		tagMap := map[string]string{}

		tagKeyReg := reTypeField.Tag.Get("Register")
		tagKeyLogin := reTypeField.Tag.Get("Login")

		tagMap["Register"] = tagKeyReg
		tagMap["Login"] = tagKeyLogin

		fieldStructList[i] = FieldStruct{Sname:reTypeField.Name, Stype:reTypeField.Type.Name(), Stag:tagMap}
	}
	return fieldStructList
}

// 获取属性值的，类型
func (urs *UserReflectStruct) GetTypeByField(fieldStr string) string {

	reType := urs.GetReflectType()
	structField,_ := reType.FieldByName(fieldStr)
	return structField.Type.Name()
}

func (urs *UserReflectStruct) GetMethodFunc(name string) reflect.Value {
	funcName := "Set" + strings.Title(name)
	return urs.GetReflectValue().MethodByName(funcName)
}

/**
rmm := reflectValue.MethodByName("Set" + strings.Title(fieldName))
				rmm.Call([]reflect.Value{reflect.ValueOf(fieldValue)})
 */
func (urs *UserReflectStruct) RunMethodByName(name string, params []reflect.Value) []reflect.Value {
	rval := urs.GetMethodFunc(name)
	//fmt.Printf("name %s:   %T \n", name, rval)
	return rval.Call(params)
}

func (urs *UserReflectStruct) RunMethodByAllName(name string, params []reflect.Value) []reflect.Value {
	rval := urs.GetReflectValue().MethodByName(name)
	//fmt.Printf("name %s:   %T \n", name, rval)
	return rval.Call(params)
}

func (urs *UserReflectStruct) __Set(name string, value reflect.Value) string {
	return "2"
}

func (urs *UserReflectStruct) GetPk() string {
	return urs.pk
}

func (urs *UserReflectStruct) SetPk(pk string) {
	urs.pk = pk
}

func (urs *UserReflectStruct) GetModelKey() interface{} {
	return urs.model
}

func (urs *UserReflectStruct) SetModelKey(model interface{} ) {
	urs.model  = model
}
