## gin框架练习
### 项目用意
使用wed开发常用的golang-gin框架进行练习demo。

### 项目依赖
#### 1. mysql
**mysql** 需要使用者自行使用docker启动mysql容器，并且自行创建**库**。

目前已实现自定义配置功能，**src/db/config.ini** 路径下可配置mysql依赖。


目前已经实现migrate功能，可以自动建立表，不需手动建立。


目前已新增主从mysql，可以藉由配置不同db与在go中引入不同db实例，实现读写分离。
```bigquery
# 配置文件
[mysql]
Db = mysql
DbHost = 127.0.0.1
DbPort = 3306
DbUser = root
DbPassWord = root
DbName = test
```
```bigquery
# 预计支持从数据库 读数据
[mysql_slave]
Db = mysql
DbHost = 127.0.0.1
DbPort = 3307
DbUser = root
DbPassWord = 123456
DbName = test
```

#### 2. redis

目前预计支持引入redis组件，对某些热点数据进行缓存，且预防缓存穿透。

### 项目目录
```bigquery
.
├── README.md
├── go.mod
├── go.sum
├── main.go // server run 和 路由逻辑 ，预计会进行拆分，把路由和server代码拆开
└── src // gin 使用的所有代码
    ├── common  
    │   └── middlewares.go  // 中间件 目前只有提供recover功能
    ├── data    // 类似dao层
    │   └── Getter
    │       └── UserGetter.go
    ├── db  // db配置
    │   ├── config.go
    │   ├── config.ini
    │   ├── db_test.go
    │   └── mysqlinit.go
    ├── handlers    // 路由
    │   ├── CommonHandler.go
    │   └── UserHandler.go
    ├── models  // 数据模型
    │   └── UserModel
    │       ├── attrs.go // 链式调用
    │       └── model.go  
    ├── result  // 统一output处理
    │   └── ErrorResult.go
    ├── test    // 测试小demo
    │   └── test.go
    └── validators  // 校验请求字段功能
        ├── Common.go
        └── UserName.go

```

### 项目启动
默认开启localhost:8080端口

在项目根目录执行 go run main.go
```
go run main.go
```
