// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Error struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetLogRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Start         string                 `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End           string                 `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLogRequest) Reset() {
	*x = GetLogRequest{}
	mi := &file_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogRequest) ProtoMessage() {}

func (x *GetLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogRequest.ProtoReflect.Descriptor instead.
func (*GetLogRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetLogRequest) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *GetLogRequest) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

type GetLogResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLogResponse) Reset() {
	*x = GetLogResponse{}
	mi := &file_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogResponse) ProtoMessage() {}

func (x *GetLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogResponse.ProtoReflect.Descriptor instead.
func (*GetLogResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetLogResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type DiscoveryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Port          int32                  `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DiscoveryResponse) Reset() {
	*x = DiscoveryResponse{}
	mi := &file_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscoveryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryResponse) ProtoMessage() {}

func (x *DiscoveryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveryResponse.ProtoReflect.Descriptor instead.
func (*DiscoveryResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *DiscoveryResponse) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type SetCmdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Error         *Error                 `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetCmdResponse) Reset() {
	*x = SetCmdResponse{}
	mi := &file_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetCmdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCmdResponse) ProtoMessage() {}

func (x *SetCmdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCmdResponse.ProtoReflect.Descriptor instead.
func (*SetCmdResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *SetCmdResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type SetCmdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         []byte                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Ttl           int32                  `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetCmdRequest) Reset() {
	*x = SetCmdRequest{}
	mi := &file_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetCmdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCmdRequest) ProtoMessage() {}

func (x *SetCmdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCmdRequest.ProtoReflect.Descriptor instead.
func (*SetCmdRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *SetCmdRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetCmdRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *SetCmdRequest) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type GetCmdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCmdRequest) Reset() {
	*x = GetCmdRequest{}
	mi := &file_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCmdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCmdRequest) ProtoMessage() {}

func (x *GetCmdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCmdRequest.ProtoReflect.Descriptor instead.
func (*GetCmdRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetCmdRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetCmdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         []byte                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Error         *Error                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCmdResponse) Reset() {
	*x = GetCmdResponse{}
	mi := &file_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCmdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCmdResponse) ProtoMessage() {}

func (x *GetCmdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCmdResponse.ProtoReflect.Descriptor instead.
func (*GetCmdResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetCmdResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *GetCmdResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type DeleteCmdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCmdRequest) Reset() {
	*x = DeleteCmdRequest{}
	mi := &file_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCmdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCmdRequest) ProtoMessage() {}

func (x *DeleteCmdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCmdRequest.ProtoReflect.Descriptor instead.
func (*DeleteCmdRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCmdRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteCmdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Error         *Error                 `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCmdResponse) Reset() {
	*x = DeleteCmdResponse{}
	mi := &file_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCmdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCmdResponse) ProtoMessage() {}

func (x *DeleteCmdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCmdResponse.ProtoReflect.Descriptor instead.
func (*DeleteCmdResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteCmdResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

const file_service_proto_rawDesc = "" +
	"\n" +
	"\rservice.proto\x12\x02pb\x1a\x1bgoogle/protobuf/empty.proto\"!\n" +
	"\x05Error\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"7\n" +
	"\rGetLogRequest\x12\x14\n" +
	"\x05start\x18\x01 \x01(\tR\x05start\x12\x10\n" +
	"\x03end\x18\x02 \x01(\tR\x03end\"$\n" +
	"\x0eGetLogResponse\x12\x12\n" +
	"\x04data\x18\x01 \x01(\fR\x04data\"'\n" +
	"\x11DiscoveryResponse\x12\x12\n" +
	"\x04port\x18\x01 \x01(\x05R\x04port\"1\n" +
	"\x0eSetCmdResponse\x12\x1f\n" +
	"\x05error\x18\x01 \x01(\v2\t.pb.ErrorR\x05error\"I\n" +
	"\rSetCmdRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\fR\x05value\x12\x10\n" +
	"\x03ttl\x18\x03 \x01(\x05R\x03ttl\"!\n" +
	"\rGetCmdRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\"G\n" +
	"\x0eGetCmdResponse\x12\x14\n" +
	"\x05value\x18\x01 \x01(\fR\x05value\x12\x1f\n" +
	"\x05error\x18\x02 \x01(\v2\t.pb.ErrorR\x05error\"$\n" +
	"\x10DeleteCmdRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\"4\n" +
	"\x11DeleteCmdResponse\x12\x1f\n" +
	"\x05error\x18\x01 \x01(\v2\t.pb.ErrorR\x05error2\x99\x01\n" +
	"\aCommand\x12,\n" +
	"\x03Set\x12\x11.pb.SetCmdRequest\x1a\x12.pb.SetCmdResponse\x12,\n" +
	"\x03Get\x12\x11.pb.GetCmdRequest\x1a\x12.pb.GetCmdResponse\x122\n" +
	"\x03Del\x12\x14.pb.DeleteCmdRequest\x1a\x15.pb.DeleteCmdResponse2G\n" +
	"\tDiscovery\x12:\n" +
	"\tGetLeader\x12\x16.google.protobuf.Empty\x1a\x15.pb.DiscoveryResponse27\n" +
	"\x04Node\x12/\n" +
	"\x06GetLog\x12\x11.pb.GetLogRequest\x1a\x12.pb.GetLogResponseB\x05Z\x03/pbb\x06proto3"

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData []byte
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_service_proto_rawDesc), len(file_service_proto_rawDesc)))
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_service_proto_goTypes = []any{
	(*Error)(nil),             // 0: pb.Error
	(*GetLogRequest)(nil),     // 1: pb.GetLogRequest
	(*GetLogResponse)(nil),    // 2: pb.GetLogResponse
	(*DiscoveryResponse)(nil), // 3: pb.DiscoveryResponse
	(*SetCmdResponse)(nil),    // 4: pb.SetCmdResponse
	(*SetCmdRequest)(nil),     // 5: pb.SetCmdRequest
	(*GetCmdRequest)(nil),     // 6: pb.GetCmdRequest
	(*GetCmdResponse)(nil),    // 7: pb.GetCmdResponse
	(*DeleteCmdRequest)(nil),  // 8: pb.DeleteCmdRequest
	(*DeleteCmdResponse)(nil), // 9: pb.DeleteCmdResponse
	(*emptypb.Empty)(nil),     // 10: google.protobuf.Empty
}
var file_service_proto_depIdxs = []int32{
	0,  // 0: pb.SetCmdResponse.error:type_name -> pb.Error
	0,  // 1: pb.GetCmdResponse.error:type_name -> pb.Error
	0,  // 2: pb.DeleteCmdResponse.error:type_name -> pb.Error
	5,  // 3: pb.Command.Set:input_type -> pb.SetCmdRequest
	6,  // 4: pb.Command.Get:input_type -> pb.GetCmdRequest
	8,  // 5: pb.Command.Del:input_type -> pb.DeleteCmdRequest
	10, // 6: pb.Discovery.GetLeader:input_type -> google.protobuf.Empty
	1,  // 7: pb.Node.GetLog:input_type -> pb.GetLogRequest
	4,  // 8: pb.Command.Set:output_type -> pb.SetCmdResponse
	7,  // 9: pb.Command.Get:output_type -> pb.GetCmdResponse
	9,  // 10: pb.Command.Del:output_type -> pb.DeleteCmdResponse
	3,  // 11: pb.Discovery.GetLeader:output_type -> pb.DiscoveryResponse
	2,  // 12: pb.Node.GetLog:output_type -> pb.GetLogResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_service_proto_rawDesc), len(file_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
