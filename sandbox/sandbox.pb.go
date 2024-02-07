// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: sandbox/sandbox.proto

package sandbox

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

type CreateSandboxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language       string `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Code           string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Stdin          string `protobuf:"bytes,3,opt,name=stdin,proto3" json:"stdin,omitempty"`
	Timeout        uint32 `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	DockerImage    string `protobuf:"bytes,5,opt,name=docker_image,json=dockerImage,proto3" json:"docker_image,omitempty"`
	CompileCommand string `protobuf:"bytes,6,opt,name=compile_command,json=compileCommand,proto3" json:"compile_command,omitempty"`
	CompileTimeout uint32 `protobuf:"varint,7,opt,name=compile_timeout,json=compileTimeout,proto3" json:"compile_timeout,omitempty"`
	MemoryLimit    uint32 `protobuf:"varint,8,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	CpuLimit       uint32 `protobuf:"varint,9,opt,name=cpu_limit,json=cpuLimit,proto3" json:"cpu_limit,omitempty"`
	RunCommand     string `protobuf:"bytes,10,opt,name=run_command,json=runCommand,proto3" json:"run_command,omitempty"`
	RunTimeout     uint32 `protobuf:"varint,11,opt,name=run_timeout,json=runTimeout,proto3" json:"run_timeout,omitempty"`
}

func (x *CreateSandboxRequest) Reset() {
	*x = CreateSandboxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sandbox_sandbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSandboxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSandboxRequest) ProtoMessage() {}

func (x *CreateSandboxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sandbox_sandbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSandboxRequest.ProtoReflect.Descriptor instead.
func (*CreateSandboxRequest) Descriptor() ([]byte, []int) {
	return file_sandbox_sandbox_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSandboxRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *CreateSandboxRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateSandboxRequest) GetStdin() string {
	if x != nil {
		return x.Stdin
	}
	return ""
}

func (x *CreateSandboxRequest) GetTimeout() uint32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *CreateSandboxRequest) GetDockerImage() string {
	if x != nil {
		return x.DockerImage
	}
	return ""
}

func (x *CreateSandboxRequest) GetCompileCommand() string {
	if x != nil {
		return x.CompileCommand
	}
	return ""
}

func (x *CreateSandboxRequest) GetCompileTimeout() uint32 {
	if x != nil {
		return x.CompileTimeout
	}
	return 0
}

func (x *CreateSandboxRequest) GetMemoryLimit() uint32 {
	if x != nil {
		return x.MemoryLimit
	}
	return 0
}

func (x *CreateSandboxRequest) GetCpuLimit() uint32 {
	if x != nil {
		return x.CpuLimit
	}
	return 0
}

func (x *CreateSandboxRequest) GetRunCommand() string {
	if x != nil {
		return x.RunCommand
	}
	return ""
}

func (x *CreateSandboxRequest) GetRunTimeout() uint32 {
	if x != nil {
		return x.RunTimeout
	}
	return 0
}

type CreateSandboxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateSandboxResponse) Reset() {
	*x = CreateSandboxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sandbox_sandbox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSandboxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSandboxResponse) ProtoMessage() {}

func (x *CreateSandboxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sandbox_sandbox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSandboxResponse.ProtoReflect.Descriptor instead.
func (*CreateSandboxResponse) Descriptor() ([]byte, []int) {
	return file_sandbox_sandbox_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSandboxResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateSandboxResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetSandboxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSandboxRequest) Reset() {
	*x = GetSandboxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sandbox_sandbox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSandboxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSandboxRequest) ProtoMessage() {}

func (x *GetSandboxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sandbox_sandbox_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSandboxRequest.ProtoReflect.Descriptor instead.
func (*GetSandboxRequest) Descriptor() ([]byte, []int) {
	return file_sandbox_sandbox_proto_rawDescGZIP(), []int{2}
}

func (x *GetSandboxRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSandboxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Language string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	Code     string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Stdin    string `protobuf:"bytes,4,opt,name=stdin,proto3" json:"stdin,omitempty"`
	Stdout   string `protobuf:"bytes,5,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr   string `protobuf:"bytes,6,opt,name=stderr,proto3" json:"stderr,omitempty"`
	ExitCode uint32 `protobuf:"varint,8,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	Error    string `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetSandboxResponse) Reset() {
	*x = GetSandboxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sandbox_sandbox_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSandboxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSandboxResponse) ProtoMessage() {}

func (x *GetSandboxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sandbox_sandbox_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSandboxResponse.ProtoReflect.Descriptor instead.
func (*GetSandboxResponse) Descriptor() ([]byte, []int) {
	return file_sandbox_sandbox_proto_rawDescGZIP(), []int{3}
}

func (x *GetSandboxResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetSandboxResponse) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *GetSandboxResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetSandboxResponse) GetStdin() string {
	if x != nil {
		return x.Stdin
	}
	return ""
}

func (x *GetSandboxResponse) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *GetSandboxResponse) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

func (x *GetSandboxResponse) GetExitCode() uint32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *GetSandboxResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_sandbox_sandbox_proto protoreflect.FileDescriptor

var file_sandbox_sandbox_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2f, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78,
	0x22, 0xed, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6e, 0x64, 0x62,
	0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x64,
	0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x70, 0x75, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x70, 0x75, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x75, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x75, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x75, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x22, 0x3d, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xcd, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6e, 0x64,
	0x62, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x64, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x64, 0x69,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64,
	0x65, 0x72, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x77, 0x75, 0x74, 0x74, 0x69, 0x6e, 0x61, 0x6e, 0x68, 0x69, 0x2f, 0x63, 0x6a,
	0x73, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sandbox_sandbox_proto_rawDescOnce sync.Once
	file_sandbox_sandbox_proto_rawDescData = file_sandbox_sandbox_proto_rawDesc
)

func file_sandbox_sandbox_proto_rawDescGZIP() []byte {
	file_sandbox_sandbox_proto_rawDescOnce.Do(func() {
		file_sandbox_sandbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_sandbox_sandbox_proto_rawDescData)
	})
	return file_sandbox_sandbox_proto_rawDescData
}

var file_sandbox_sandbox_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sandbox_sandbox_proto_goTypes = []interface{}{
	(*CreateSandboxRequest)(nil),  // 0: sandbox.CreateSandboxRequest
	(*CreateSandboxResponse)(nil), // 1: sandbox.CreateSandboxResponse
	(*GetSandboxRequest)(nil),     // 2: sandbox.GetSandboxRequest
	(*GetSandboxResponse)(nil),    // 3: sandbox.GetSandboxResponse
}
var file_sandbox_sandbox_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sandbox_sandbox_proto_init() }
func file_sandbox_sandbox_proto_init() {
	if File_sandbox_sandbox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sandbox_sandbox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSandboxRequest); i {
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
		file_sandbox_sandbox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSandboxResponse); i {
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
		file_sandbox_sandbox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSandboxRequest); i {
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
		file_sandbox_sandbox_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSandboxResponse); i {
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
			RawDescriptor: file_sandbox_sandbox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sandbox_sandbox_proto_goTypes,
		DependencyIndexes: file_sandbox_sandbox_proto_depIdxs,
		MessageInfos:      file_sandbox_sandbox_proto_msgTypes,
	}.Build()
	File_sandbox_sandbox_proto = out.File
	file_sandbox_sandbox_proto_rawDesc = nil
	file_sandbox_sandbox_proto_goTypes = nil
	file_sandbox_sandbox_proto_depIdxs = nil
}
