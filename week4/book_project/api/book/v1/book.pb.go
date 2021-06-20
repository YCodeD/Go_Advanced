// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: book/v1/book.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetBookInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBookInfoRequest) Reset() {
	*x = GetBookInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_v1_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookInfoRequest) ProtoMessage() {}

func (x *GetBookInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_v1_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookInfoRequest.ProtoReflect.Descriptor instead.
func (*GetBookInfoRequest) Descriptor() ([]byte, []int) {
	return file_book_v1_book_proto_rawDescGZIP(), []int{0}
}

type GetBookInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookTitle  string `protobuf:"bytes,1,opt,name=bookTitle,proto3" json:"bookTitle,omitempty"`
	BookAuthor string `protobuf:"bytes,2,opt,name=bookAuthor,proto3" json:"bookAuthor,omitempty"`
}

func (x *GetBookInfoResponse) Reset() {
	*x = GetBookInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_v1_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookInfoResponse) ProtoMessage() {}

func (x *GetBookInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_v1_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookInfoResponse.ProtoReflect.Descriptor instead.
func (*GetBookInfoResponse) Descriptor() ([]byte, []int) {
	return file_book_v1_book_proto_rawDescGZIP(), []int{1}
}

func (x *GetBookInfoResponse) GetBookTitle() string {
	if x != nil {
		return x.BookTitle
	}
	return ""
}

func (x *GetBookInfoResponse) GetBookAuthor() string {
	if x != nil {
		return x.BookAuthor
	}
	return ""
}

var File_book_v1_book_proto protoreflect.FileDescriptor

var file_book_v1_book_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x22, 0x14, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x53, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f,
	0x6f, 0x6b, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x6f, 0x6f, 0x6b, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6f,
	0x6f, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x32, 0x58, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_book_v1_book_proto_rawDescOnce sync.Once
	file_book_v1_book_proto_rawDescData = file_book_v1_book_proto_rawDesc
)

func file_book_v1_book_proto_rawDescGZIP() []byte {
	file_book_v1_book_proto_rawDescOnce.Do(func() {
		file_book_v1_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_book_v1_book_proto_rawDescData)
	})
	return file_book_v1_book_proto_rawDescData
}

var file_book_v1_book_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_book_v1_book_proto_goTypes = []interface{}{
	(*GetBookInfoRequest)(nil),  // 0: book.v1.GetBookInfoRequest
	(*GetBookInfoResponse)(nil), // 1: book.v1.GetBookInfoResponse
}
var file_book_v1_book_proto_depIdxs = []int32{
	0, // 0: book.v1.BookServer.GetBookInfo:input_type -> book.v1.GetBookInfoRequest
	1, // 1: book.v1.BookServer.GetBookInfo:output_type -> book.v1.GetBookInfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_book_v1_book_proto_init() }
func file_book_v1_book_proto_init() {
	if File_book_v1_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_book_v1_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookInfoRequest); i {
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
		file_book_v1_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookInfoResponse); i {
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
			RawDescriptor: file_book_v1_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_book_v1_book_proto_goTypes,
		DependencyIndexes: file_book_v1_book_proto_depIdxs,
		MessageInfos:      file_book_v1_book_proto_msgTypes,
	}.Build()
	File_book_v1_book_proto = out.File
	file_book_v1_book_proto_rawDesc = nil
	file_book_v1_book_proto_goTypes = nil
	file_book_v1_book_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BookServerClient is the client API for BookServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookServerClient interface {
	GetBookInfo(ctx context.Context, in *GetBookInfoRequest, opts ...grpc.CallOption) (*GetBookInfoResponse, error)
}

type bookServerClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServerClient(cc grpc.ClientConnInterface) BookServerClient {
	return &bookServerClient{cc}
}

func (c *bookServerClient) GetBookInfo(ctx context.Context, in *GetBookInfoRequest, opts ...grpc.CallOption) (*GetBookInfoResponse, error) {
	out := new(GetBookInfoResponse)
	err := c.cc.Invoke(ctx, "/book.v1.BookServer/GetBookInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServerServer is the server API for BookServer service.
type BookServerServer interface {
	GetBookInfo(context.Context, *GetBookInfoRequest) (*GetBookInfoResponse, error)
}

// UnimplementedBookServerServer can be embedded to have forward compatible implementations.
type UnimplementedBookServerServer struct {
}

func (*UnimplementedBookServerServer) GetBookInfo(context.Context, *GetBookInfoRequest) (*GetBookInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookInfo not implemented")
}

func RegisterBookServerServer(s *grpc.Server, srv BookServerServer) {
	s.RegisterService(&_BookServer_serviceDesc, srv)
}

func _BookServer_GetBookInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServerServer).GetBookInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.v1.BookServer/GetBookInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServerServer).GetBookInfo(ctx, req.(*GetBookInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "book.v1.BookServer",
	HandlerType: (*BookServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBookInfo",
			Handler:    _BookServer_GetBookInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book/v1/book.proto",
}
