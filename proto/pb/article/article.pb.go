// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.11.4
// source: article.proto

package article

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/mwitkow/go-proto-validators"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type HealthzResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReturnMessage string `protobuf:"bytes,1,opt,name=return_message,json=returnMessage,proto3" json:"return_message,omitempty"`
}

func (x *HealthzResponse) Reset() {
	*x = HealthzResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthzResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthzResponse) ProtoMessage() {}

func (x *HealthzResponse) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthzResponse.ProtoReflect.Descriptor instead.
func (*HealthzResponse) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{0}
}

func (x *HealthzResponse) GetReturnMessage() string {
	if x != nil {
		return x.ReturnMessage
	}
	return ""
}

type ArticleListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemCode string `protobuf:"bytes,1,opt,name=system_code,json=systemCode,proto3" json:"system_code,omitempty"`
}

func (x *ArticleListRequest) Reset() {
	*x = ArticleListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleListRequest) ProtoMessage() {}

func (x *ArticleListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleListRequest.ProtoReflect.Descriptor instead.
func (*ArticleListRequest) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{1}
}

func (x *ArticleListRequest) GetSystemCode() string {
	if x != nil {
		return x.SystemCode
	}
	return ""
}

type ArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemCode string `protobuf:"bytes,1,opt,name=system_code,json=systemCode,proto3" json:"system_code,omitempty"`
	ArticleId  string `protobuf:"bytes,2,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
}

func (x *ArticleRequest) Reset() {
	*x = ArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleRequest) ProtoMessage() {}

func (x *ArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleRequest.ProtoReflect.Descriptor instead.
func (*ArticleRequest) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{2}
}

func (x *ArticleRequest) GetSystemCode() string {
	if x != nil {
		return x.SystemCode
	}
	return ""
}

func (x *ArticleRequest) GetArticleId() string {
	if x != nil {
		return x.ArticleId
	}
	return ""
}

type ArticleListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articles []*ArticleResponse `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
}

func (x *ArticleListResponse) Reset() {
	*x = ArticleListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleListResponse) ProtoMessage() {}

func (x *ArticleListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleListResponse.ProtoReflect.Descriptor instead.
func (*ArticleListResponse) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{3}
}

func (x *ArticleListResponse) GetArticles() []*ArticleResponse {
	if x != nil {
		return x.Articles
	}
	return nil
}

type ArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *ArticleResponse) Reset() {
	*x = ArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleResponse) ProtoMessage() {}

func (x *ArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleResponse.ProtoReflect.Descriptor instead.
func (*ArticleResponse) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{4}
}

func (x *ArticleResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ArticleResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_article_proto protoreflect.FileDescriptor

var file_article_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a,
	0x0f, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x47, 0x0a, 0x12, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a,
	0x0b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x10, 0xe2, 0xdf, 0x1f, 0x0c, 0x0a, 0x0a, 0x5e, 0x28, 0x5b, 0x41, 0x42, 0x5a,
	0x5d, 0x29, 0x3f, 0x24, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x50, 0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x49, 0x64, 0x22, 0x4b, 0x0a, 0x13, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x22,
	0x3f, 0x0a, 0x0f, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x32, 0xda, 0x01, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x07,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x18, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x7a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a,
	0x09, 0x2e, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_article_proto_rawDescOnce sync.Once
	file_article_proto_rawDescData = file_article_proto_rawDesc
)

func file_article_proto_rawDescGZIP() []byte {
	file_article_proto_rawDescOnce.Do(func() {
		file_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_article_proto_rawDescData)
	})
	return file_article_proto_rawDescData
}

var file_article_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_article_proto_goTypes = []interface{}{
	(*HealthzResponse)(nil),     // 0: article.HealthzResponse
	(*ArticleListRequest)(nil),  // 1: article.ArticleListRequest
	(*ArticleRequest)(nil),      // 2: article.ArticleRequest
	(*ArticleListResponse)(nil), // 3: article.ArticleListResponse
	(*ArticleResponse)(nil),     // 4: article.ArticleResponse
	(*empty.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_article_proto_depIdxs = []int32{
	4, // 0: article.ArticleListResponse.articles:type_name -> article.ArticleResponse
	5, // 1: article.Article.Healthz:input_type -> google.protobuf.Empty
	1, // 2: article.Article.GetArticleList:input_type -> article.ArticleListRequest
	2, // 3: article.Article.GetArticle:input_type -> article.ArticleRequest
	0, // 4: article.Article.Healthz:output_type -> article.HealthzResponse
	3, // 5: article.Article.GetArticleList:output_type -> article.ArticleListResponse
	4, // 6: article.Article.GetArticle:output_type -> article.ArticleResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_article_proto_init() }
func file_article_proto_init() {
	if File_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthzResponse); i {
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
		file_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleListRequest); i {
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
		file_article_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleRequest); i {
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
		file_article_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleListResponse); i {
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
		file_article_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleResponse); i {
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
			RawDescriptor: file_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_article_proto_goTypes,
		DependencyIndexes: file_article_proto_depIdxs,
		MessageInfos:      file_article_proto_msgTypes,
	}.Build()
	File_article_proto = out.File
	file_article_proto_rawDesc = nil
	file_article_proto_goTypes = nil
	file_article_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ArticleClient is the client API for Article service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArticleClient interface {
	Healthz(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthzResponse, error)
	GetArticleList(ctx context.Context, in *ArticleListRequest, opts ...grpc.CallOption) (*ArticleListResponse, error)
	GetArticle(ctx context.Context, in *ArticleRequest, opts ...grpc.CallOption) (*ArticleResponse, error)
}

type articleClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleClient(cc grpc.ClientConnInterface) ArticleClient {
	return &articleClient{cc}
}

func (c *articleClient) Healthz(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthzResponse, error) {
	out := new(HealthzResponse)
	err := c.cc.Invoke(ctx, "/article.Article/Healthz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) GetArticleList(ctx context.Context, in *ArticleListRequest, opts ...grpc.CallOption) (*ArticleListResponse, error) {
	out := new(ArticleListResponse)
	err := c.cc.Invoke(ctx, "/article.Article/GetArticleList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) GetArticle(ctx context.Context, in *ArticleRequest, opts ...grpc.CallOption) (*ArticleResponse, error) {
	out := new(ArticleResponse)
	err := c.cc.Invoke(ctx, "/article.Article/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServer is the server API for Article service.
type ArticleServer interface {
	Healthz(context.Context, *empty.Empty) (*HealthzResponse, error)
	GetArticleList(context.Context, *ArticleListRequest) (*ArticleListResponse, error)
	GetArticle(context.Context, *ArticleRequest) (*ArticleResponse, error)
}

// UnimplementedArticleServer can be embedded to have forward compatible implementations.
type UnimplementedArticleServer struct {
}

func (*UnimplementedArticleServer) Healthz(context.Context, *empty.Empty) (*HealthzResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthz not implemented")
}
func (*UnimplementedArticleServer) GetArticleList(context.Context, *ArticleListRequest) (*ArticleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleList not implemented")
}
func (*UnimplementedArticleServer) GetArticle(context.Context, *ArticleRequest) (*ArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}

func RegisterArticleServer(s *grpc.Server, srv ArticleServer) {
	s.RegisterService(&_Article_serviceDesc, srv)
}

func _Article_Healthz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).Healthz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.Article/Healthz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).Healthz(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_GetArticleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).GetArticleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.Article/GetArticleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).GetArticleList(ctx, req.(*ArticleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.Article/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).GetArticle(ctx, req.(*ArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Article_serviceDesc = grpc.ServiceDesc{
	ServiceName: "article.Article",
	HandlerType: (*ArticleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthz",
			Handler:    _Article_Healthz_Handler,
		},
		{
			MethodName: "GetArticleList",
			Handler:    _Article_GetArticleList_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _Article_GetArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}
