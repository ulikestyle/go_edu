package model

import (
	"bufio"
	"edu/container"
	"edu/util"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// 统一的模型接口
type Model interface {
	ToString() string // 格式化输出数据信息
}

var (
	path   string                 = "/home/wwwroot/go/devtools/src/learning/edu/data/" // 数据路径
	suffix string                 = ".sql"
	models map[string]Model // 记录标识 user =》 结构体
)

// "User" => User{} 不支持
func init() {

	container.CreateContainersFactory().Set("userModel", NewUser())
	UserDatas = make(map[string]Model, 0)

	rfdata("user", "username", UserDatas)
}

// 数据库文件 -> 通过配置设置
// name 数据库名称   user,admin
// pirmay 查询主键
// models 存放数据
func rfdata(name, primary string, datas map[string]Model) error {
	// 1. 读取数据库文件 -》读取那个文件？
	f, ferr := os.Open(path + name + suffix) // 根据路径读取文件信息
	if ferr != nil {
		fmt.Println("文件读取异常,", ferr)
		return errors.New("文件查询不到 error")
	}
	// 关闭文件的资源流
	defer f.Close()
	// 创建读取的文件的缓冲区
	buf := bufio.NewReader(f)
	// 2. 遍历数据  每一行的数据 字段根据 , 分割；数据通过 \n 分割
	titles := make([]string, 0)
	modelKey := ""
	for {
		row, rerr := buf.ReadBytes('\n') // 根据换行读取文件信息 , 返回的是byte[]
		if rerr != nil {
			if rerr == io.EOF { // 是否文件读取结束
				break
			}
			fmt.Println("抛出缓存读取文件异常", rerr)
		}
		// 去掉字符串，并分割数据
		fields := strings.Split(strings.TrimSuffix(string(row), "\r\n"), ",")
		//fmt.Println("field的值为：",fields)

		// 根据数据判断操作 ； 是记录字段还是设置数据
		if len(titles) == 0 {
			titles = fields // 存储字段
			continue
		} else {

			//反射得到类
			var instanceModel = container.CreateContainersFactory().Get(name+"Model").(Model)
			utilUserRef := util.UserReflectStruct{}
			utilUserRef.SetModelKey(instanceModel)

			for i, v := range titles {

				fieldName := v
				fieldValue := fields[i]
				if fieldName == primary {
					modelKey = fieldValue
				}

				//fmt.Printf("fieldValue %s, 值是 %T, 地址是： %T \n", fieldValue,fieldValue, fieldValue)

				fieldType := utilUserRef.GetTypeByField(fieldName)
				//fmt.Printf("fieldType %s  %T \n", fieldType, fieldType)
				if fieldType == "int" {
					age, _ := strconv.Atoi(fieldValue)
					utilUserRef.RunMethodByName(fieldName,[]reflect.Value{reflect.ValueOf(age)})
				}else{
					utilUserRef.RunMethodByName(fieldName,[]reflect.Value{reflect.ValueOf(fieldValue)})
				}
			}
			//fmt.Print(instanceModel)
			//fmt.Println()
			datas[modelKey] = instanceModel
		}
	}
	return nil
}

//获取数据
func RwData() {

}

func getFilePath(name string){
	//realPath := path + name + suffix
}

func GetCacheData() map[string]Model{
	return models
}

