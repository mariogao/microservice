golang 微服务项目服务端代码

运行：go run main.go

情况说明 ：我在本地启动时，遇到了报错，然后把go.mod中的 google.golang.org/protobuf v1.25.0 改为 google.golang.org/protobuf v1.23.0 后项目正常启动
百度的结果是google.golang.org/grpc 版本的问题，有些版本是不支持 clientv3 的，所以改成v1.23.0。
如果1.25版本也能正常启动，请忽略此条信息。
客户端地址：https://github.com/mariogao/microclient.git
