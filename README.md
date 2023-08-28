# jax
该项目为新webserver，取名来自LOL的武器大师：贾克斯(Jax)，希望新webserver作为上层服务，可以像武器大师任何武器都能精通一样，可以流畅的接入任何服务，并能够向前端高效、稳定的输出。


[更多功能介绍及使用手册参考wiki(链接待完善)>>](http://127.0.0.1)


## 快速开始
### go version
 GO >= 1.13

### 服务构建和启动

* 直接启动
```
go run main.go
```


## 框架规范

* [Go 语言规范]
* [接口规范]
* [错误处理规范]
* [配置规范]
* [日志规范]
* [CICD规范]
* [包管理规范]
* [运行规范]
* [目录规范] 
强烈建议按照以下目录规范来规范你的项目：

 ```
|-demo-project
    |-api 访问第三方请求目录
    |-conf 配置文件目录
        |-app 用户自定义配置，跟随代码发布的业务配置
        |-mount 用来放置环境相关的配置，可通过配置中心发布的配置
    |-controllers 控制器目录 业务处理方法统一入口
        |-http http控制器目录
        |-command 任务控制器入口，包括cycle任务、crontab任务、一次性任务
        |-mq 消息队列回调入口
        |-rpcx rpcx服务控制器入口
    |-data 数据层。当项目比较复杂时，可以增加data层用于组装数据，包括不限于数据库查询到的数据、api调用后查询到的数据
    |-defines 一些常量声明，类型定义，初始化
    |-dto 数据传输对象定义目录，各层之间数据传输的结构体在此目录定义
    |-helpers 公共类目录，可以用来初始化一些全局变量，http和rpcx客户端
    |-message 返回给调用方的错误码，或多国语言信息放在此目录
    |-middleware 业务中间件
        |-http gin的中间件放在此目录
        |-rpcx rpcx的plugin放在此目录
    |-models 数据模型访问目录。数据库，redis相关调用封装。
    |-router 路由目录，一般对应controllers目录结构
        |-http http路由
        |-command 人物类路由
        |-mq 消息队列路由
        |-rpcx rpcx路由
    |-service 业务逻辑聚合目录。主要强调业务逻辑，能够看出一个功能的核心处理流程。
    |-sql mysql建库建表相关语句
    |-utils  一些工具类
    |-go.mod go module使用，记录项目的依赖
    |-go.sum go mod tidy 后生成，记录依赖的详细依赖
    |-main.go 程序执行入口
  ```
