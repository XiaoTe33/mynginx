# mynginx

## 启动方式

`./etc/config.yaml`

配置好配置文件

`go run ./cmd/main.go`

## example

`go run ./cmd/main.go`

`go run ./test/routers/main.go`

多次访问`localhost:80/a/mynginx`(负载均衡)

访问`localhost:80/fs/main.go`(文件服务)