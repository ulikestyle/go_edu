# go_edu

#### 目标
- 通过edu系统，对数组，切片，map操作，文件读写，反射调用进行练习操作，加深印象

#### 功能需求
- 数据操作，这里未使用数据库，直接使用文件存储数据
    - 文件读取
    - 由于文件不具备搜索功能，项目初始化时，需要读出文件数据，并以用户名为键保存在map中
- 视图层
    - 用户登录
        - 登录完成后，可以选择展示用户信息
    - 用户注册
 
#### 实现思路
 
```
避免重复实例化和NewUser()操作，从其他项目借鉴了 container思路，保存在容器中;
- userModel ,对于*User
- isLogin，存放的是用户登录标识（已登录，存放的是用户的唯一标识，在这里使用username）

view层实现思路，已实现
1. 利用标识符(controllerName, actionName),
- controllerName 为controller的struct结构体名称
- actionName 为所需执行的方法
2. 反射拿到value， call调用方法
```