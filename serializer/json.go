package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// ProtobufToJSON converts protocol buffer message into JSON string
func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false, // 枚举类型是否使用整型
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true, // json的字段名是否使用原名（否则使用驼峰法）
	}
	return marshaler.MarshalToString(message)
}
