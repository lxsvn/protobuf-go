// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/test_import_a1m1.proto

package imports

import (
	test_a_1 "github.com/golang/protobuf/v2/cmd/protoc-gen-go/testdata/imports/test_a_1"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoiface "github.com/golang/protobuf/v2/runtime/protoiface"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

type A1M1 struct {
	F                    *test_a_1.M1 `protobuf:"bytes,1,opt,name=f,proto3" json:"f,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (x *A1M1) Reset() {
	*x = A1M1{}
}

func (x *A1M1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*A1M1) ProtoMessage() {}

func (x *A1M1) ProtoReflect() protoreflect.Message {
	return file_imports_test_import_a1m1_proto_msgTypes[0].MessageOf(x)
}

func (m *A1M1) XXX_Methods() *protoiface.Methods {
	return file_imports_test_import_a1m1_proto_msgTypes[0].Methods()
}

// Deprecated: Use A1M1.ProtoReflect.Type instead.
func (*A1M1) Descriptor() ([]byte, []int) {
	return file_imports_test_import_a1m1_proto_rawDescGZIP(), []int{0}
}

func (x *A1M1) GetF() *test_a_1.M1 {
	if x != nil {
		return x.F
	}
	return nil
}

var File_imports_test_import_a1m1_proto protoreflect.FileDescriptor

var file_imports_test_import_a1m1_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x31, 0x6d, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x5f, 0x31, 0x2f, 0x6d, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x20, 0x0a, 0x04, 0x41, 0x31, 0x4d, 0x31, 0x12, 0x18, 0x0a, 0x01, 0x66, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x2e, 0x4d, 0x31,
	0x52, 0x01, 0x66, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_imports_test_import_a1m1_proto_rawDescOnce sync.Once
	file_imports_test_import_a1m1_proto_rawDescData = file_imports_test_import_a1m1_proto_rawDesc
)

func file_imports_test_import_a1m1_proto_rawDescGZIP() []byte {
	file_imports_test_import_a1m1_proto_rawDescOnce.Do(func() {
		file_imports_test_import_a1m1_proto_rawDescData = protoimpl.X.CompressGZIP(file_imports_test_import_a1m1_proto_rawDescData)
	})
	return file_imports_test_import_a1m1_proto_rawDescData
}

var file_imports_test_import_a1m1_proto_msgTypes = make([]protoimpl.MessageType, 1)
var file_imports_test_import_a1m1_proto_goTypes = []interface{}{
	(*A1M1)(nil),        // 0: test.A1M1
	(*test_a_1.M1)(nil), // 1: test.a.M1
}
var file_imports_test_import_a1m1_proto_depIdxs = []int32{
	1, // test.A1M1.f:type_name -> test.a.M1
}

func init() { file_imports_test_import_a1m1_proto_init() }
func file_imports_test_import_a1m1_proto_init() {
	if File_imports_test_import_a1m1_proto != nil {
		return
	}
	File_imports_test_import_a1m1_proto = protoimpl.FileBuilder{
		RawDescriptor:      file_imports_test_import_a1m1_proto_rawDesc,
		GoTypes:            file_imports_test_import_a1m1_proto_goTypes,
		DependencyIndexes:  file_imports_test_import_a1m1_proto_depIdxs,
		MessageOutputTypes: file_imports_test_import_a1m1_proto_msgTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	file_imports_test_import_a1m1_proto_rawDesc = nil
	file_imports_test_import_a1m1_proto_goTypes = nil
	file_imports_test_import_a1m1_proto_depIdxs = nil
}
