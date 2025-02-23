// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.0
// source: api/product/v1/service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

// 创建商品的请求消息
type CreateGoodsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`               // 商品名称
	Tags          string                 `protobuf:"bytes,2,opt,name=tags,proto3" json:"tags,omitempty"`               // 商品标签
	Type          string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`               // 商品类型
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"` // 商品描述
	Price         float64                `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`           // 商品价格
	Quantity      int32                  `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`      // 商品数量
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateGoodsRequest) Reset() {
	*x = CreateGoodsRequest{}
	mi := &file_api_product_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoodsRequest) ProtoMessage() {}

func (x *CreateGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoodsRequest.ProtoReflect.Descriptor instead.
func (*CreateGoodsRequest) Descriptor() ([]byte, []int) {
	return file_api_product_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGoodsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGoodsRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *CreateGoodsRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateGoodsRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateGoodsRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateGoodsRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

// 创建商品的响应消息
type CreateGoodsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 商品ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateGoodsResponse) Reset() {
	*x = CreateGoodsResponse{}
	mi := &file_api_product_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGoodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoodsResponse) ProtoMessage() {}

func (x *CreateGoodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoodsResponse.ProtoReflect.Descriptor instead.
func (*CreateGoodsResponse) Descriptor() ([]byte, []int) {
	return file_api_product_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGoodsResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 更新商品的请求消息
type UpdateGoodsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                  // 商品ID
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`               // 商品名称
	Tags          string                 `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`               // 商品标签
	Type          string                 `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`               // 商品类型
	Description   string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"` // 商品描述
	Price         float64                `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`           // 商品价格
	Quantity      int32                  `protobuf:"varint,7,opt,name=quantity,proto3" json:"quantity,omitempty"`      // 商品数量
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateGoodsRequest) Reset() {
	*x = UpdateGoodsRequest{}
	mi := &file_api_product_v1_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoodsRequest) ProtoMessage() {}

func (x *UpdateGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoodsRequest.ProtoReflect.Descriptor instead.
func (*UpdateGoodsRequest) Descriptor() ([]byte, []int) {
	return file_api_product_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateGoodsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateGoodsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateGoodsRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *UpdateGoodsRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateGoodsRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateGoodsRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateGoodsRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

// 更新商品的响应消息
type UpdateGoodsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 更新是否成功
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateGoodsResponse) Reset() {
	*x = UpdateGoodsResponse{}
	mi := &file_api_product_v1_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateGoodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoodsResponse) ProtoMessage() {}

func (x *UpdateGoodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoodsResponse.ProtoReflect.Descriptor instead.
func (*UpdateGoodsResponse) Descriptor() ([]byte, []int) {
	return file_api_product_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateGoodsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 删除商品的请求消息
type DeleteGoodsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 商品ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteGoodsRequest) Reset() {
	*x = DeleteGoodsRequest{}
	mi := &file_api_product_v1_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGoodsRequest) ProtoMessage() {}

func (x *DeleteGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGoodsRequest.ProtoReflect.Descriptor instead.
func (*DeleteGoodsRequest) Descriptor() ([]byte, []int) {
	return file_api_product_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteGoodsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 删除商品的响应消息
type DeleteGoodsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 删除是否成功
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteGoodsResponse) Reset() {
	*x = DeleteGoodsResponse{}
	mi := &file_api_product_v1_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteGoodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGoodsResponse) ProtoMessage() {}

func (x *DeleteGoodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGoodsResponse.ProtoReflect.Descriptor instead.
func (*DeleteGoodsResponse) Descriptor() ([]byte, []int) {
	return file_api_product_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteGoodsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 过滤商品的请求消息
type GoodsFilterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Keywords      string                 `protobuf:"bytes,1,opt,name=keywords,proto3" json:"keywords,omitempty"`   // 关键词
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`           // 商品类型
	MinPrice      float64                `protobuf:"fixed64,3,opt,name=minPrice,proto3" json:"minPrice,omitempty"` // 最低价格
	MaxPrice      float64                `protobuf:"fixed64,4,opt,name=maxPrice,proto3" json:"maxPrice,omitempty"` // 最高价格
	Page          int32                  `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`          // 页码
	PageSize      int32                  `protobuf:"varint,6,opt,name=pageSize,proto3" json:"pageSize,omitempty"`  // 每页数量
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GoodsFilterRequest) Reset() {
	*x = GoodsFilterRequest{}
	mi := &file_api_product_v1_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GoodsFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsFilterRequest) ProtoMessage() {}

func (x *GoodsFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsFilterRequest.ProtoReflect.Descriptor instead.
func (*GoodsFilterRequest) Descriptor() ([]byte, []int) {
	return file_api_product_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *GoodsFilterRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *GoodsFilterRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GoodsFilterRequest) GetMinPrice() float64 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *GoodsFilterRequest) GetMaxPrice() float64 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

func (x *GoodsFilterRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GoodsFilterRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// 列出商品的响应消息
type GoodsListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int64                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"` // 商品总数
	List          []*GoodsInfoResponse   `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`    // 商品列表
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GoodsListResponse) Reset() {
	*x = GoodsListResponse{}
	mi := &file_api_product_v1_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GoodsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsListResponse) ProtoMessage() {}

func (x *GoodsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsListResponse.ProtoReflect.Descriptor instead.
func (*GoodsListResponse) Descriptor() ([]byte, []int) {
	return file_api_product_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *GoodsListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GoodsListResponse) GetList() []*GoodsInfoResponse {
	if x != nil {
		return x.List
	}
	return nil
}

// 商品信息响应消息
type GoodsInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                  // 商品ID
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`               // 商品名称
	Tags          string                 `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`               // 商品标签
	Type          string                 `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`               // 商品类型
	Description   string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"` // 商品描述
	Price         float64                `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`           // 商品价格
	Quantity      int32                  `protobuf:"varint,7,opt,name=quantity,proto3" json:"quantity,omitempty"`      // 商品数量
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GoodsInfoResponse) Reset() {
	*x = GoodsInfoResponse{}
	mi := &file_api_product_v1_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GoodsInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsInfoResponse) ProtoMessage() {}

func (x *GoodsInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsInfoResponse.ProtoReflect.Descriptor instead.
func (*GoodsInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_product_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *GoodsInfoResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GoodsInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoodsInfoResponse) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *GoodsInfoResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GoodsInfoResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GoodsInfoResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GoodsInfoResponse) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

// 搜索商品的请求消息
type SearchGoodsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *structpb.Struct       `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`        // 搜索查询条件
	Page          int32                  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`         // 页码
	PageSize      int32                  `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"` // 每页数量
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchGoodsRequest) Reset() {
	*x = SearchGoodsRequest{}
	mi := &file_api_product_v1_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchGoodsRequest) ProtoMessage() {}

func (x *SearchGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchGoodsRequest.ProtoReflect.Descriptor instead.
func (*SearchGoodsRequest) Descriptor() ([]byte, []int) {
	return file_api_product_v1_service_proto_rawDescGZIP(), []int{9}
}

func (x *SearchGoodsRequest) GetQuery() *structpb.Struct {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *SearchGoodsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchGoodsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// 搜索商品的响应消息
type SearchGoodsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*GoodsInfoResponse   `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"` // 搜索结果
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchGoodsResponse) Reset() {
	*x = SearchGoodsResponse{}
	mi := &file_api_product_v1_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchGoodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchGoodsResponse) ProtoMessage() {}

func (x *SearchGoodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchGoodsResponse.ProtoReflect.Descriptor instead.
func (*SearchGoodsResponse) Descriptor() ([]byte, []int) {
	return file_api_product_v1_service_proto_rawDescGZIP(), []int{10}
}

func (x *SearchGoodsResponse) GetResults() []*GoodsInfoResponse {
	if x != nil {
		return x.Results
	}
	return nil
}

// 自动补全请求消息
type AutoCompleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Prefix        string                 `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"` // 前缀
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AutoCompleteRequest) Reset() {
	*x = AutoCompleteRequest{}
	mi := &file_api_product_v1_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AutoCompleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoCompleteRequest) ProtoMessage() {}

func (x *AutoCompleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoCompleteRequest.ProtoReflect.Descriptor instead.
func (*AutoCompleteRequest) Descriptor() ([]byte, []int) {
	return file_api_product_v1_service_proto_rawDescGZIP(), []int{11}
}

func (x *AutoCompleteRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

// 自动补全响应消息
type AutoCompleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Suggestions   []string               `protobuf:"bytes,1,rep,name=suggestions,proto3" json:"suggestions,omitempty"` // 建议列表
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AutoCompleteResponse) Reset() {
	*x = AutoCompleteResponse{}
	mi := &file_api_product_v1_service_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AutoCompleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoCompleteResponse) ProtoMessage() {}

func (x *AutoCompleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_service_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoCompleteResponse.ProtoReflect.Descriptor instead.
func (*AutoCompleteResponse) Descriptor() ([]byte, []int) {
	return file_api_product_v1_service_proto_rawDescGZIP(), []int{12}
}

func (x *AutoCompleteResponse) GetSuggestions() []string {
	if x != nil {
		return x.Suggestions
	}
	return nil
}

var File_api_product_v1_service_proto protoreflect.FileDescriptor

var file_api_product_v1_service_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x25, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xb4, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x2f, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x24, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x2f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0xac, 0x01, 0x0a, 0x12, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69,
	0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x69,
	0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x22, 0x5c, 0x0a, 0x11, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x31, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x22, 0xb3, 0x01, 0x0a, 0x11, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x73, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x4e, 0x0a, 0x13, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x2d, 0x0a, 0x13, 0x41,
	0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x38, 0x0a, 0x14, 0x41, 0x75,
	0x74, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x32, 0x8e, 0x05, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a,
	0x01, 0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x5d, 0x0a,
	0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x69, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x1a, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x66, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a,
	0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x6b, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x77, 0x0a, 0x12,
	0x41, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x50, 0x0a, 0x19, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x42, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x56, 0x31, 0x50, 0x01, 0x5a, 0x21, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_product_v1_service_proto_rawDescOnce sync.Once
	file_api_product_v1_service_proto_rawDescData []byte
)

func file_api_product_v1_service_proto_rawDescGZIP() []byte {
	file_api_product_v1_service_proto_rawDescOnce.Do(func() {
		file_api_product_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_product_v1_service_proto_rawDesc), len(file_api_product_v1_service_proto_rawDesc)))
	})
	return file_api_product_v1_service_proto_rawDescData
}

var file_api_product_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_product_v1_service_proto_goTypes = []any{
	(*CreateGoodsRequest)(nil),   // 0: product.v1.CreateGoodsRequest
	(*CreateGoodsResponse)(nil),  // 1: product.v1.CreateGoodsResponse
	(*UpdateGoodsRequest)(nil),   // 2: product.v1.UpdateGoodsRequest
	(*UpdateGoodsResponse)(nil),  // 3: product.v1.UpdateGoodsResponse
	(*DeleteGoodsRequest)(nil),   // 4: product.v1.DeleteGoodsRequest
	(*DeleteGoodsResponse)(nil),  // 5: product.v1.DeleteGoodsResponse
	(*GoodsFilterRequest)(nil),   // 6: product.v1.GoodsFilterRequest
	(*GoodsListResponse)(nil),    // 7: product.v1.GoodsListResponse
	(*GoodsInfoResponse)(nil),    // 8: product.v1.GoodsInfoResponse
	(*SearchGoodsRequest)(nil),   // 9: product.v1.SearchGoodsRequest
	(*SearchGoodsResponse)(nil),  // 10: product.v1.SearchGoodsResponse
	(*AutoCompleteRequest)(nil),  // 11: product.v1.AutoCompleteRequest
	(*AutoCompleteResponse)(nil), // 12: product.v1.AutoCompleteResponse
	(*structpb.Struct)(nil),      // 13: google.protobuf.Struct
}
var file_api_product_v1_service_proto_depIdxs = []int32{
	8,  // 0: product.v1.GoodsListResponse.list:type_name -> product.v1.GoodsInfoResponse
	13, // 1: product.v1.SearchGoodsRequest.query:type_name -> google.protobuf.Struct
	8,  // 2: product.v1.SearchGoodsResponse.results:type_name -> product.v1.GoodsInfoResponse
	0,  // 3: product.v1.ProductService.CreateGoods:input_type -> product.v1.CreateGoodsRequest
	6,  // 4: product.v1.ProductService.GoodsList:input_type -> product.v1.GoodsFilterRequest
	2,  // 5: product.v1.ProductService.UpdateGoods:input_type -> product.v1.UpdateGoodsRequest
	4,  // 6: product.v1.ProductService.DeleteGoods:input_type -> product.v1.DeleteGoodsRequest
	9,  // 7: product.v1.ProductService.SearchGoods:input_type -> product.v1.SearchGoodsRequest
	11, // 8: product.v1.ProductService.AutocompleteSearch:input_type -> product.v1.AutoCompleteRequest
	1,  // 9: product.v1.ProductService.CreateGoods:output_type -> product.v1.CreateGoodsResponse
	7,  // 10: product.v1.ProductService.GoodsList:output_type -> product.v1.GoodsListResponse
	3,  // 11: product.v1.ProductService.UpdateGoods:output_type -> product.v1.UpdateGoodsResponse
	5,  // 12: product.v1.ProductService.DeleteGoods:output_type -> product.v1.DeleteGoodsResponse
	10, // 13: product.v1.ProductService.SearchGoods:output_type -> product.v1.SearchGoodsResponse
	12, // 14: product.v1.ProductService.AutocompleteSearch:output_type -> product.v1.AutoCompleteResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_product_v1_service_proto_init() }
func file_api_product_v1_service_proto_init() {
	if File_api_product_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_product_v1_service_proto_rawDesc), len(file_api_product_v1_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_product_v1_service_proto_goTypes,
		DependencyIndexes: file_api_product_v1_service_proto_depIdxs,
		MessageInfos:      file_api_product_v1_service_proto_msgTypes,
	}.Build()
	File_api_product_v1_service_proto = out.File
	file_api_product_v1_service_proto_goTypes = nil
	file_api_product_v1_service_proto_depIdxs = nil
}
