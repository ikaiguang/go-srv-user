# User

用户服务

参考文档：

- [kratos-github](https://github.com/go-kratos/kratos)
- [kratos-docs](https://go-kratos.dev/docs/)

## 运行

在当前目录先运行

```shell
go run ./cmd/main/... -conf=./configs

curl http://127.0.0.1:8081/api/v1/ping/hello
curl http://127.0.0.1:8081/api/v1/ping/error
```

## 执行生成脚本 与 编译proto

```shell

# 执行生成脚本 与 编译proto
go run ./cmd/proto/... -path=./api/user
go run ./cmd/proto/... -path=./api/xxx
    
```
