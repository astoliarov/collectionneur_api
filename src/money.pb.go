// Code generated by protoc-gen-go. DO NOT EDIT.
// source: money.proto

package src

/*
May cause problems
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type OrderField int32

const (
	OrderField_id   OrderField = 0
	OrderField_date OrderField = 1
)

var OrderField_name = map[int32]string{
	0: "id",
	1: "date",
}
var OrderField_value = map[string]int32{
	"id":   0,
	"date": 1,
}

func (x OrderField) String() string {
	return proto.EnumName(OrderField_name, int32(x))
}
func (OrderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_money_7753967e7af42a70, []int{0}
}

type OrderDirection int32

const (
	OrderDirection_asc  OrderDirection = 0
	OrderDirection_desc OrderDirection = 1
)

var OrderDirection_name = map[int32]string{
	0: "asc",
	1: "desc",
}
var OrderDirection_value = map[string]int32{
	"asc":  0,
	"desc": 1,
}

func (x OrderDirection) String() string {
	return proto.EnumName(OrderDirection_name, int32(x))
}
func (OrderDirection) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_money_7753967e7af42a70, []int{1}
}

type SpendInfo struct {
	Currency             string   `protobuf:"bytes,1,opt,name=currency" json:"currency,omitempty"`
	Sum                  float32  `protobuf:"fixed32,2,opt,name=sum" json:"sum,omitempty"`
	Card                 string   `protobuf:"bytes,3,opt,name=card" json:"card,omitempty"`
	Date                 uint64   `protobuf:"varint,4,opt,name=date" json:"date,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Category             string   `protobuf:"bytes,6,opt,name=category" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpendInfo) Reset()         { *m = SpendInfo{} }
func (m *SpendInfo) String() string { return proto.CompactTextString(m) }
func (*SpendInfo) ProtoMessage()    {}
func (*SpendInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_money_7753967e7af42a70, []int{0}
}
func (m *SpendInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpendInfo.Unmarshal(m, b)
}
func (m *SpendInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpendInfo.Marshal(b, m, deterministic)
}
func (dst *SpendInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpendInfo.Merge(dst, src)
}
func (m *SpendInfo) XXX_Size() int {
	return xxx_messageInfo_SpendInfo.Size(m)
}
func (m *SpendInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SpendInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SpendInfo proto.InternalMessageInfo

func (m *SpendInfo) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *SpendInfo) GetSum() float32 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *SpendInfo) GetCard() string {
	if m != nil {
		return m.Card
	}
	return ""
}

func (m *SpendInfo) GetDate() uint64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *SpendInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SpendInfo) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

type CategoryResolver struct {
	Description          string   `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	Category             string   `protobuf:"bytes,2,opt,name=category" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryResolver) Reset()         { *m = CategoryResolver{} }
func (m *CategoryResolver) String() string { return proto.CompactTextString(m) }
func (*CategoryResolver) ProtoMessage()    {}
func (*CategoryResolver) Descriptor() ([]byte, []int) {
	return fileDescriptor_money_7753967e7af42a70, []int{1}
}
func (m *CategoryResolver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryResolver.Unmarshal(m, b)
}
func (m *CategoryResolver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryResolver.Marshal(b, m, deterministic)
}
func (dst *CategoryResolver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryResolver.Merge(dst, src)
}
func (m *CategoryResolver) XXX_Size() int {
	return xxx_messageInfo_CategoryResolver.Size(m)
}
func (m *CategoryResolver) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryResolver.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryResolver proto.InternalMessageInfo

func (m *CategoryResolver) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CategoryResolver) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

type AddInfosResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddInfosResponse) Reset()         { *m = AddInfosResponse{} }
func (m *AddInfosResponse) String() string { return proto.CompactTextString(m) }
func (*AddInfosResponse) ProtoMessage()    {}
func (*AddInfosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_money_7753967e7af42a70, []int{2}
}
func (m *AddInfosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddInfosResponse.Unmarshal(m, b)
}
func (m *AddInfosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddInfosResponse.Marshal(b, m, deterministic)
}
func (dst *AddInfosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddInfosResponse.Merge(dst, src)
}
func (m *AddInfosResponse) XXX_Size() int {
	return xxx_messageInfo_AddInfosResponse.Size(m)
}
func (m *AddInfosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddInfosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddInfosResponse proto.InternalMessageInfo

type AddInfosRequest struct {
	Infos                []*SpendInfo `protobuf:"bytes,1,rep,name=infos" json:"infos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AddInfosRequest) Reset()         { *m = AddInfosRequest{} }
func (m *AddInfosRequest) String() string { return proto.CompactTextString(m) }
func (*AddInfosRequest) ProtoMessage()    {}
func (*AddInfosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_money_7753967e7af42a70, []int{3}
}
func (m *AddInfosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddInfosRequest.Unmarshal(m, b)
}
func (m *AddInfosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddInfosRequest.Marshal(b, m, deterministic)
}
func (dst *AddInfosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddInfosRequest.Merge(dst, src)
}
func (m *AddInfosRequest) XXX_Size() int {
	return xxx_messageInfo_AddInfosRequest.Size(m)
}
func (m *AddInfosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddInfosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddInfosRequest proto.InternalMessageInfo

func (m *AddInfosRequest) GetInfos() []*SpendInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

type OrderBy struct {
	Field                OrderField     `protobuf:"varint,1,opt,name=field,enum=src.OrderField" json:"field,omitempty"`
	Direction            OrderDirection `protobuf:"varint,2,opt,name=direction,enum=src.OrderDirection" json:"direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *OrderBy) Reset()         { *m = OrderBy{} }
func (m *OrderBy) String() string { return proto.CompactTextString(m) }
func (*OrderBy) ProtoMessage()    {}
func (*OrderBy) Descriptor() ([]byte, []int) {
	return fileDescriptor_money_7753967e7af42a70, []int{4}
}
func (m *OrderBy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderBy.Unmarshal(m, b)
}
func (m *OrderBy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderBy.Marshal(b, m, deterministic)
}
func (dst *OrderBy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderBy.Merge(dst, src)
}
func (m *OrderBy) XXX_Size() int {
	return xxx_messageInfo_OrderBy.Size(m)
}
func (m *OrderBy) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderBy.DiscardUnknown(m)
}

var xxx_messageInfo_OrderBy proto.InternalMessageInfo

func (m *OrderBy) GetField() OrderField {
	if m != nil {
		return m.Field
	}
	return OrderField_id
}

func (m *OrderBy) GetDirection() OrderDirection {
	if m != nil {
		return m.Direction
	}
	return OrderDirection_asc
}

type RequestMeta struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestMeta) Reset()         { *m = RequestMeta{} }
func (m *RequestMeta) String() string { return proto.CompactTextString(m) }
func (*RequestMeta) ProtoMessage()    {}
func (*RequestMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_money_7753967e7af42a70, []int{5}
}
func (m *RequestMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMeta.Unmarshal(m, b)
}
func (m *RequestMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMeta.Marshal(b, m, deterministic)
}
func (dst *RequestMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMeta.Merge(dst, src)
}
func (m *RequestMeta) XXX_Size() int {
	return xxx_messageInfo_RequestMeta.Size(m)
}
func (m *RequestMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMeta.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMeta proto.InternalMessageInfo

func (m *RequestMeta) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *RequestMeta) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ResponseMeta struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	Count                int32    `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseMeta) Reset()         { *m = ResponseMeta{} }
func (m *ResponseMeta) String() string { return proto.CompactTextString(m) }
func (*ResponseMeta) ProtoMessage()    {}
func (*ResponseMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_money_7753967e7af42a70, []int{6}
}
func (m *ResponseMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMeta.Unmarshal(m, b)
}
func (m *ResponseMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMeta.Marshal(b, m, deterministic)
}
func (dst *ResponseMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMeta.Merge(dst, src)
}
func (m *ResponseMeta) XXX_Size() int {
	return xxx_messageInfo_ResponseMeta.Size(m)
}
func (m *ResponseMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMeta proto.InternalMessageInfo

func (m *ResponseMeta) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ResponseMeta) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ResponseMeta) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetInfosRequest struct {
	Meta                 *RequestMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	OrderBy              *OrderBy     `protobuf:"bytes,2,opt,name=orderBy" json:"orderBy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetInfosRequest) Reset()         { *m = GetInfosRequest{} }
func (m *GetInfosRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfosRequest) ProtoMessage()    {}
func (*GetInfosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_money_7753967e7af42a70, []int{7}
}
func (m *GetInfosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfosRequest.Unmarshal(m, b)
}
func (m *GetInfosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfosRequest.Marshal(b, m, deterministic)
}
func (dst *GetInfosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfosRequest.Merge(dst, src)
}
func (m *GetInfosRequest) XXX_Size() int {
	return xxx_messageInfo_GetInfosRequest.Size(m)
}
func (m *GetInfosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfosRequest proto.InternalMessageInfo

func (m *GetInfosRequest) GetMeta() *RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetInfosRequest) GetOrderBy() *OrderBy {
	if m != nil {
		return m.OrderBy
	}
	return nil
}

type GetInfosResponse struct {
	Meta                 *ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Infos                []*SpendInfo  `protobuf:"bytes,2,rep,name=infos" json:"infos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetInfosResponse) Reset()         { *m = GetInfosResponse{} }
func (m *GetInfosResponse) String() string { return proto.CompactTextString(m) }
func (*GetInfosResponse) ProtoMessage()    {}
func (*GetInfosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_money_7753967e7af42a70, []int{8}
}
func (m *GetInfosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfosResponse.Unmarshal(m, b)
}
func (m *GetInfosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfosResponse.Marshal(b, m, deterministic)
}
func (dst *GetInfosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfosResponse.Merge(dst, src)
}
func (m *GetInfosResponse) XXX_Size() int {
	return xxx_messageInfo_GetInfosResponse.Size(m)
}
func (m *GetInfosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfosResponse proto.InternalMessageInfo

func (m *GetInfosResponse) GetMeta() *ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetInfosResponse) GetInfos() []*SpendInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

func init() {
	proto.RegisterType((*SpendInfo)(nil), "src.SpendInfo")
	proto.RegisterType((*CategoryResolver)(nil), "src.CategoryResolver")
	proto.RegisterType((*AddInfosResponse)(nil), "src.AddInfosResponse")
	proto.RegisterType((*AddInfosRequest)(nil), "src.AddInfosRequest")
	proto.RegisterType((*OrderBy)(nil), "src.OrderBy")
	proto.RegisterType((*RequestMeta)(nil), "src.RequestMeta")
	proto.RegisterType((*ResponseMeta)(nil), "src.ResponseMeta")
	proto.RegisterType((*GetInfosRequest)(nil), "src.GetInfosRequest")
	proto.RegisterType((*GetInfosResponse)(nil), "src.GetInfosResponse")
	proto.RegisterEnum("src.OrderField", OrderField_name, OrderField_value)
	proto.RegisterEnum("src.OrderDirection", OrderDirection_name, OrderDirection_value)
}

func init() { proto.RegisterFile("money.proto", fileDescriptor_money_7753967e7af42a70) }

var fileDescriptor_money_7753967e7af42a70 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xce, 0xda, 0x71, 0x7e, 0xc6, 0x55, 0x62, 0x96, 0x82, 0xac, 0x1e, 0x90, 0x65, 0x28, 0x8a,
	0x7a, 0x88, 0x84, 0x39, 0x20, 0xc4, 0x89, 0x16, 0x81, 0x38, 0x20, 0xd0, 0xf6, 0x01, 0x2a, 0xb3,
	0x9e, 0x20, 0x4b, 0x89, 0x37, 0xec, 0x6e, 0x2a, 0xf9, 0xcc, 0x83, 0xf0, 0xaa, 0xc8, 0xe3, 0x75,
	0xdc, 0xfa, 0x10, 0x89, 0xdb, 0xce, 0xcc, 0x37, 0xf3, 0xcd, 0x37, 0x33, 0x0b, 0xe1, 0x4e, 0x55,
	0x58, 0xaf, 0xf7, 0x5a, 0x59, 0xc5, 0x7d, 0xa3, 0x65, 0xfa, 0x97, 0xc1, 0xfc, 0x76, 0x8f, 0x55,
	0xf1, 0xb5, 0xda, 0x28, 0x7e, 0x01, 0x33, 0x79, 0xd0, 0x1a, 0x2b, 0x59, 0xc7, 0x2c, 0x61, 0xab,
	0xb9, 0x38, 0xda, 0x3c, 0x02, 0xdf, 0x1c, 0x76, 0xb1, 0x97, 0xb0, 0x95, 0x27, 0x9a, 0x27, 0xe7,
	0x30, 0x96, 0xb9, 0x2e, 0x62, 0x9f, 0x90, 0xf4, 0x6e, 0x7c, 0x45, 0x6e, 0x31, 0x1e, 0x27, 0x6c,
	0x35, 0x16, 0xf4, 0xe6, 0x09, 0x84, 0x05, 0x1a, 0xa9, 0xcb, 0xbd, 0x2d, 0x55, 0x15, 0x07, 0x04,
	0x7f, 0xe8, 0x22, 0xde, 0xdc, 0xe2, 0x2f, 0xa5, 0xeb, 0x78, 0xe2, 0x78, 0x9d, 0x9d, 0xfe, 0x80,
	0xe8, 0xc6, 0xbd, 0x05, 0x1a, 0xb5, 0xbd, 0x47, 0x3d, 0xac, 0xc8, 0x4e, 0x57, 0xf4, 0x06, 0x15,
	0x39, 0x44, 0x1f, 0x0b, 0x12, 0x6c, 0x04, 0x9a, 0xbd, 0xaa, 0x0c, 0xa6, 0xef, 0x60, 0xd9, 0xfb,
	0x7e, 0x1f, 0xd0, 0x58, 0xfe, 0x0a, 0x82, 0xb2, 0xb1, 0x63, 0x96, 0xf8, 0xab, 0x30, 0x5b, 0xac,
	0x8d, 0x96, 0xeb, 0xe3, 0xac, 0x44, 0x1b, 0x4c, 0x25, 0x4c, 0xbf, 0xeb, 0x02, 0xf5, 0x75, 0xcd,
	0x2f, 0x21, 0xd8, 0x94, 0xb8, 0x2d, 0xa8, 0x9f, 0x45, 0xb6, 0xa4, 0x04, 0x0a, 0x7e, 0x6e, 0xdc,
	0xa2, 0x8d, 0xf2, 0x37, 0x30, 0x2f, 0x4a, 0x8d, 0x92, 0x5a, 0xf7, 0x08, 0xfa, 0xb4, 0x87, 0x7e,
	0xea, 0x42, 0xa2, 0x47, 0xa5, 0x1f, 0x20, 0x74, 0x5d, 0x7d, 0x43, 0x9b, 0xf3, 0x73, 0x08, 0xb6,
	0xe5, 0xae, 0xb4, 0x44, 0x14, 0x88, 0xd6, 0xe0, 0xcf, 0x61, 0xa2, 0x36, 0x1b, 0x83, 0x96, 0x8a,
	0x06, 0xc2, 0x59, 0xa9, 0x80, 0xb3, 0x4e, 0xe6, 0xff, 0x67, 0x37, 0x68, 0xa9, 0x0e, 0x95, 0xa5,
	0x2d, 0x07, 0xa2, 0x35, 0xd2, 0x3b, 0x58, 0x7e, 0x41, 0x3b, 0x18, 0xd7, 0x78, 0x87, 0x36, 0xa7,
	0xaa, 0x61, 0x16, 0x91, 0xa2, 0x07, 0x4d, 0x0b, 0x8a, 0xf2, 0xd7, 0x30, 0x55, 0xed, 0xb8, 0x88,
	0x27, 0xcc, 0xce, 0x7a, 0xe9, 0xd7, 0xb5, 0xe8, 0x82, 0xe9, 0x1d, 0x44, 0x3d, 0x41, 0xdb, 0x3c,
	0xbf, 0x7c, 0xc4, 0xf0, 0xc4, 0x31, 0xf4, 0xca, 0x1c, 0xc5, 0x71, 0x6f, 0xde, 0x89, 0xbd, 0x5d,
	0xbd, 0x00, 0xe8, 0x57, 0xc3, 0x27, 0xe0, 0x95, 0x45, 0x34, 0xe2, 0xb3, 0xf6, 0x7c, 0x23, 0x76,
	0xf5, 0x12, 0x16, 0x8f, 0xf7, 0xc1, 0xa7, 0xe0, 0xe7, 0x46, 0x3a, 0x10, 0x1a, 0x19, 0xb1, 0xec,
	0x0f, 0x83, 0xf0, 0x46, 0x19, 0x7b, 0x8b, 0xfa, 0xbe, 0x94, 0xc8, 0xdf, 0xc3, 0xac, 0xbb, 0x22,
	0x7e, 0x4e, 0xbc, 0x83, 0xa3, 0xba, 0x78, 0x36, 0xf0, 0xba, 0xf3, 0x1b, 0x35, 0xa9, 0x9d, 0x60,
	0x97, 0x3a, 0x18, 0xb0, 0x4b, 0x1d, 0x4e, 0x25, 0x1d, 0xfd, 0x9c, 0xd0, 0x7f, 0x7e, 0xfb, 0x2f,
	0x00, 0x00, 0xff, 0xff, 0x48, 0x68, 0xc5, 0x9d, 0xde, 0x03, 0x00, 0x00,
}
