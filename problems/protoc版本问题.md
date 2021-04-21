# 生成的客户端和服务端接口代码在哪里？

使用protoc生成 xx.pb.go 后，想找到server的接口以及用grpc声明的函数，然而没有找到。在 [gRPC官方文档](http://doc.oschina.net/grpc?t=60133) 中寻找问题。发现一个问题
- 官方文档中生成`pb.go`的指令中的 `--go_out=plugins=grpc`，实际使用提示不支持。是官方文档没有更新吗?
- 文档给了 一个[教程仓库](https://github.com/grpc/grpc-go/tree/master/examples/route_guide)。 在 `issue` 中果然有人提了一样的问题。答案是：
  ```
  You most likely have the wrong version of protoc-gen-go (yours is from google.golang.org/protobuf, but the expected one is from github.com/golang/protobuf).
  ```

关键字 `protoc-gen-go` 意思是版本错误。项目开始之前，确实使用了`go get` 下载了这个插件。应该是没有将`GOPATH/bin` 添加到环境变量的原因。

环境问题解决后，输入` protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb` 发现 现在的gRPC版本生成的go代码将客户端，服务端接口等代码都放在`xxx_grpc.pb.go`文件。
# 问题总结

当时使用`--go_out=plugins=grpc` 报错不支持plugins，没有对其深究。
