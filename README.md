# 环境配置

在启动微服务之前需要安装etcd。
注意：这里使用的是v2版本的go-micro，其与v1不兼容，在使用时需要导入并更新以下包，否则会报错。
```
"github.com/micro/go-micro/v2"
"github.com/micro/go-micro/v2/registry"
"github.com/micro/go-plugins/registry/etcdv3/v2"
"github.com/micro/protoc-gen-micro/v2"

```

# 启动程序

## 启动服务端

```
go run hello/main.go
```

## 启动客户端

```
go run hello/client/client.go
```