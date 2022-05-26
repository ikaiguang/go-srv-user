# User

用户服务

参考文档：

- [kratos-github](https://github.com/go-kratos/kratos)
- [kratos-docs](https://go-kratos.dev/docs/)

## 运行

在当前目录先运行

```shell
# 运行程序
# make run
go run ./cmd/main/... -conf=./configs

# 运行测试
# make ping
curl http://127.0.0.1:8081/api/v1/ping/hello
curl http://127.0.0.1:8081/api/v1/ping/error
```

## 执行生成脚本 与 编译proto

```shell

# 执行生成脚本 与 编译proto
# make proto_user
# make proto_xxx
# kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user/v1/enums/user.enum.v1.proto
go run ./cmd/proto/... -path=./api/user
go run ./cmd/proto/... -path=./api/xxx
    
```

## 创建数据库

```shell
CREATE DATABASE srv_user DEFAULT CHARSET utf8mb4;
```
