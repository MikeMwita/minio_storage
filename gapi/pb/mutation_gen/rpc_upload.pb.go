// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpc_upload.proto

package pb

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

type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileContent *FileContent `protobuf:"bytes,1,opt,name=file_content,json=fileContent,proto3" json:"file_content,omitempty"`
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_upload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_upload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_rpc_upload_proto_rawDescGZIP(), []int{0}
}

func (x *UploadFileRequest) GetFileContent() *FileContent {
	if x != nil {
		return x.FileContent
	}
	return nil
}

type UploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success   bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Etag      string `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
	VersionId string `protobuf:"bytes,4,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
}

func (x *UploadFileResponse) Reset() {
	*x = UploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_upload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResponse) ProtoMessage() {}

func (x *UploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_upload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResponse.ProtoReflect.Descriptor instead.
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return file_rpc_upload_proto_rawDescGZIP(), []int{1}
}

func (x *UploadFileResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UploadFileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UploadFileResponse) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *UploadFileResponse) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

var File_rpc_upload_proto protoreflect.FileDescriptor

var file_rpc_upload_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x47, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x0b,
	0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x7b, 0x0a, 0x12, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x69, 0x6c, 0x74, 0x72, 0x6f, 0x6e, 0x69, 0x63,
	0x2f, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x2f, 0x67, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_upload_proto_rawDescOnce sync.Once
	file_rpc_upload_proto_rawDescData = file_rpc_upload_proto_rawDesc
)

func file_rpc_upload_proto_rawDescGZIP() []byte {
	file_rpc_upload_proto_rawDescOnce.Do(func() {
		file_rpc_upload_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_upload_proto_rawDescData)
	})
	return file_rpc_upload_proto_rawDescData
}

var file_rpc_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_upload_proto_goTypes = []interface{}{
	(*UploadFileRequest)(nil),  // 0: pb.UploadFileRequest
	(*UploadFileResponse)(nil), // 1: pb.UploadFileResponse
	(*FileContent)(nil),        // 2: pb.FileContent
}
var file_rpc_upload_proto_depIdxs = []int32{
	2, // 0: pb.UploadFileRequest.file_content:type_name -> pb.FileContent
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_upload_proto_init() }
func file_rpc_upload_proto_init() {
	if File_rpc_upload_proto != nil {
		return
	}
	file_file_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_upload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
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
		file_rpc_upload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileResponse); i {
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
			RawDescriptor: file_rpc_upload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_upload_proto_goTypes,
		DependencyIndexes: file_rpc_upload_proto_depIdxs,
		MessageInfos:      file_rpc_upload_proto_msgTypes,
	}.Build()
	File_rpc_upload_proto = out.File
	file_rpc_upload_proto_rawDesc = nil
	file_rpc_upload_proto_goTypes = nil
	file_rpc_upload_proto_depIdxs = nil
}