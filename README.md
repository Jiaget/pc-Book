# protocol message
```
syntax = "proto3";

message <MessageName> {
    <data-type> fiedl_1 = tag_ 1;
    ...
}
```
- message 首字母大写
- 字段名小写
- 数据类型
  - string, bool, bytes
  - float,double
  - int32,int64,uint32....
- 数据类型还可以是自定义的枚举型或者其他message类型
- **tags比字段名还要重要**
  - tag 可以是任意的整数
    - tag 可以是 1 到 2**(29-1)之间的数
    - 除了19000 到 19999（因为这段数字作为保留数字供网络通信使用的）
  - 1到15占1个字节
  - 16到2047占两个字节
  - 上面两条意味着，频繁使用的字段应该标记1到15，减少开支
  - tag可以不用按顺序排列，但是同一级的字段的tag必须是独一无二的
# 生成代码
写好proto文件后，下载`protoc`与相关包
```
https://github.com/protocolbuffers/protobuf/releases
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get google.golang.org/protobuf/cmd/protoc-gen-go

```
写下面的命令生成go代码

`protoc -I=proto proto/*.proto --go_out=pb` 

- 其中，`-I` 代表porto文件的相对路径，`-go_out`代表输出的go代码路径， `proto/*.proto` 代表proto目录下所有的proto文件。
- 注意，当前gRPC版本需要在文件中添加两条指令
```
package proto; // 声明包名，避免冲突
option go_package ="./"; // 声明go代码输出路径，该路径已写在生成指令里，但这里又必须写，随便写一个当前路径。
```