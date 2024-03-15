// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: proto/eqstruct/world.proto

package eqstruct

import (
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

type EnterWorld struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *EnterWorld) Reset() {
	*x = EnterWorld{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eqstruct_world_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterWorld) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterWorld) ProtoMessage() {}

func (x *EnterWorld) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eqstruct_world_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterWorld.ProtoReflect.Descriptor instead.
func (*EnterWorld) Descriptor() ([]byte, []int) {
	return file_proto_eqstruct_world_proto_rawDescGZIP(), []int{0}
}

func (x *EnterWorld) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ServerMOTD struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ServerMOTD) Reset() {
	*x = ServerMOTD{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eqstruct_world_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerMOTD) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerMOTD) ProtoMessage() {}

func (x *ServerMOTD) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eqstruct_world_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerMOTD.ProtoReflect.Descriptor instead.
func (*ServerMOTD) Descriptor() ([]byte, []int) {
	return file_proto_eqstruct_world_proto_rawDescGZIP(), []int{1}
}

func (x *ServerMOTD) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServerMOTD) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_eqstruct_world_proto protoreflect.FileDescriptor

var file_proto_eqstruct_world_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x71, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x71,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x20, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4d, 0x4f, 0x54, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x6d, 0x6f, 0x72, 0x65, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x7a,
	0x2f, 0x67, 0x68, 0x6f, 0x65, 0x71, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x71, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_eqstruct_world_proto_rawDescOnce sync.Once
	file_proto_eqstruct_world_proto_rawDescData = file_proto_eqstruct_world_proto_rawDesc
)

func file_proto_eqstruct_world_proto_rawDescGZIP() []byte {
	file_proto_eqstruct_world_proto_rawDescOnce.Do(func() {
		file_proto_eqstruct_world_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_eqstruct_world_proto_rawDescData)
	})
	return file_proto_eqstruct_world_proto_rawDescData
}

var file_proto_eqstruct_world_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_eqstruct_world_proto_goTypes = []interface{}{
	(*EnterWorld)(nil), // 0: eqstruct.EnterWorld
	(*ServerMOTD)(nil), // 1: eqstruct.ServerMOTD
}
var file_proto_eqstruct_world_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_eqstruct_world_proto_init() }
func file_proto_eqstruct_world_proto_init() {
	if File_proto_eqstruct_world_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_eqstruct_world_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterWorld); i {
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
		file_proto_eqstruct_world_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerMOTD); i {
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
			RawDescriptor: file_proto_eqstruct_world_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_eqstruct_world_proto_goTypes,
		DependencyIndexes: file_proto_eqstruct_world_proto_depIdxs,
		MessageInfos:      file_proto_eqstruct_world_proto_msgTypes,
	}.Build()
	File_proto_eqstruct_world_proto = out.File
	file_proto_eqstruct_world_proto_rawDesc = nil
	file_proto_eqstruct_world_proto_goTypes = nil
	file_proto_eqstruct_world_proto_depIdxs = nil
}
