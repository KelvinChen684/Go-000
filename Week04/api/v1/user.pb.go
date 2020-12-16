// 定义项目 API 的 proto 文件 可以同时描述 gRPC 和 HTTP API
// protobuf 文件参考:
//  - https://developers.google.com/protocol-buffers/

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: api/v1/user.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Pass    string `protobuf:"bytes,2,opt,name=Pass,proto3" json:"Pass,omitempty"`
	NewPass string `protobuf:"bytes,3,opt,name=NewPass,proto3" json:"NewPass,omitempty"`
	Gender  string `protobuf:"bytes,4,opt,name=Gender,proto3" json:"Gender,omitempty"`
}

func (x *UserReq) Reset() {
	*x = UserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReq) ProtoMessage() {}

func (x *UserReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReq.ProtoReflect.Descriptor instead.
func (*UserReq) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserReq) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

func (x *UserReq) GetNewPass() string {
	if x != nil {
		return x.NewPass
	}
	return ""
}

func (x *UserReq) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Gender string `protobuf:"bytes,2,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Age    int32  `protobuf:"varint,3,opt,name=Age,proto3" json:"Age,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserInfo) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type UserStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UserStatus) Reset() {
	*x = UserStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStatus) ProtoMessage() {}

func (x *UserStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStatus.ProtoReflect.Descriptor instead.
func (*UserStatus) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserStatus) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_v1_user_proto protoreflect.FileDescriptor

var file_api_v1_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x50, 0x61, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x41,
	0x67, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x98, 0x01, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x08, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x0b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0c,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x12, 0x08,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x65, 0x6c, 0x76, 0x69, 0x6e, 0x43, 0x68, 0x65, 0x6e, 0x36, 0x38,
	0x34, 0x2f, 0x47, 0x6f, 0x2d, 0x30, 0x30, 0x30, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61,
	0x69, 0x6e, 0x2f, 0x57, 0x65, 0x65, 0x6b, 0x30, 0x34, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_user_proto_rawDescOnce sync.Once
	file_api_v1_user_proto_rawDescData = file_api_v1_user_proto_rawDesc
)

func file_api_v1_user_proto_rawDescGZIP() []byte {
	file_api_v1_user_proto_rawDescOnce.Do(func() {
		file_api_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_user_proto_rawDescData)
	})
	return file_api_v1_user_proto_rawDescData
}

var file_api_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1_user_proto_goTypes = []interface{}{
	(*UserReq)(nil),    // 0: UserReq
	(*UserInfo)(nil),   // 1: UserInfo
	(*UserStatus)(nil), // 2: UserStatus
}
var file_api_v1_user_proto_depIdxs = []int32{
	0, // 0: User.GetUserInfo:input_type -> UserReq
	0, // 1: User.CreateUser:input_type -> UserReq
	0, // 2: User.Authenticate:input_type -> UserReq
	0, // 3: User.SetPass:input_type -> UserReq
	1, // 4: User.GetUserInfo:output_type -> UserInfo
	2, // 5: User.CreateUser:output_type -> UserStatus
	2, // 6: User.Authenticate:output_type -> UserStatus
	2, // 7: User.SetPass:output_type -> UserStatus
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_user_proto_init() }
func file_api_v1_user_proto_init() {
	if File_api_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_user_proto_goTypes,
		DependencyIndexes: file_api_v1_user_proto_depIdxs,
		MessageInfos:      file_api_v1_user_proto_msgTypes,
	}.Build()
	File_api_v1_user_proto = out.File
	file_api_v1_user_proto_rawDesc = nil
	file_api_v1_user_proto_goTypes = nil
	file_api_v1_user_proto_depIdxs = nil
}
