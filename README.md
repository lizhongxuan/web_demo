# web_demo

## 文件说明
> conf 配置文件

> handler api接口逻辑
>> home.name 简单的接口

> pkg 工具包
>> errno 错误返回封装,根据公司要求封装,例子没使用

> router 路由

> vendor 第三方依赖包 (以下命令已经执行,无须再执行)
>> 命令 go mod init example.com/m/v2 初始化go.mod
>> 命令 go mod vendor 拉取依赖包

> main.go 程序入口

## 启动
> go build -o web_demo main.go

## 运行
> ./web_demo

## 请求
>http://127.0.0.1:6666/home/name?_w_fname=av
>返回 av hellow