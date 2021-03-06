// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/net.proto

package core // import "github.com/duckchat/duckchat-gateway/proto/core"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TransportDataHeaderKey int32

const (
	TransportDataHeaderKey_HeaderInvalid   TransportDataHeaderKey = 0
	TransportDataHeaderKey_HeaderErrorCode TransportDataHeaderKey = 1
	TransportDataHeaderKey_HeaderErrorInfo TransportDataHeaderKey = 2
	TransportDataHeaderKey_HeaderSessionid TransportDataHeaderKey = 3
	// the value must be set for websocket and zaly request.
	//
	// (scheme://)?host:(port)?(/path)?
	//
	// http://127.0.0.1
	// http://127.0.0.1:8080
	// http://127.0.0.1/child-dir/
	// zaly://127.0.0.1
	// zaly://127.0.0.1:3681
	// ws://127.0.0.1:8888/child-dir/
	TransportDataHeaderKey_HeaderHostUrl TransportDataHeaderKey = 4
	// zalyUrl
	TransportDataHeaderKey_HeaderReferer TransportDataHeaderKey = 5
	// the value is a string.
	// the jsRequest in web must set the value with the browser user agent.
	TransportDataHeaderKey_HeaderUserAgent TransportDataHeaderKey = 6
	// whether the response can be cached by the client.
	TransportDataHeaderKey_HeaderAllowCache TransportDataHeaderKey = 7
	// (string)UserClientLangType please when set its value.
	TransportDataHeaderKey_HeaderUserClientLang TransportDataHeaderKey = 8
	// Reference: (message Version).json
	TransportDataHeaderKey_HeaderApplicationVersion TransportDataHeaderKey = 10
)

var TransportDataHeaderKey_name = map[int32]string{
	0:  "HeaderInvalid",
	1:  "HeaderErrorCode",
	2:  "HeaderErrorInfo",
	3:  "HeaderSessionid",
	4:  "HeaderHostUrl",
	5:  "HeaderReferer",
	6:  "HeaderUserAgent",
	7:  "HeaderAllowCache",
	8:  "HeaderUserClientLang",
	10: "HeaderApplicationVersion",
}
var TransportDataHeaderKey_value = map[string]int32{
	"HeaderInvalid":            0,
	"HeaderErrorCode":          1,
	"HeaderErrorInfo":          2,
	"HeaderSessionid":          3,
	"HeaderHostUrl":            4,
	"HeaderReferer":            5,
	"HeaderUserAgent":          6,
	"HeaderAllowCache":         7,
	"HeaderUserClientLang":     8,
	"HeaderApplicationVersion": 10,
}

func (x TransportDataHeaderKey) String() string {
	return proto.EnumName(TransportDataHeaderKey_name, int32(x))
}
func (TransportDataHeaderKey) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_net_3bb38ce5f35e60f5, []int{0}
}

type UserClientLangType int32

const (
	UserClientLangType_UserClientLangEN UserClientLangType = 0
	UserClientLangType_UserClientLangZH UserClientLangType = 1
)

var UserClientLangType_name = map[int32]string{
	0: "UserClientLangEN",
	1: "UserClientLangZH",
}
var UserClientLangType_value = map[string]int32{
	"UserClientLangEN": 0,
	"UserClientLangZH": 1,
}

func (x UserClientLangType) String() string {
	return proto.EnumName(UserClientLangType_name, int32(x))
}
func (UserClientLangType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_net_3bb38ce5f35e60f5, []int{1}
}

type GatewayType int32

const (
	GatewayType_GatewayInvalid   GatewayType = 0
	GatewayType_GatewayWebsocket GatewayType = 1
	GatewayType_GatewayZaly      GatewayType = 2
)

var GatewayType_name = map[int32]string{
	0: "GatewayInvalid",
	1: "GatewayWebsocket",
	2: "GatewayZaly",
}
var GatewayType_value = map[string]int32{
	"GatewayInvalid":   0,
	"GatewayWebsocket": 1,
	"GatewayZaly":      2,
}

func (x GatewayType) String() string {
	return proto.EnumName(GatewayType_name, int32(x))
}
func (GatewayType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_net_3bb38ce5f35e60f5, []int{2}
}

type UserClientType int32

const (
	UserClientType_UserClientMobileInvalid UserClientType = 0
	UserClientType_UserClientMobileApp     UserClientType = 1
	UserClientType_UserClientWeb           UserClientType = 2
)

var UserClientType_name = map[int32]string{
	0: "UserClientMobileInvalid",
	1: "UserClientMobileApp",
	2: "UserClientWeb",
}
var UserClientType_value = map[string]int32{
	"UserClientMobileInvalid": 0,
	"UserClientMobileApp":     1,
	"UserClientWeb":           2,
}

func (x UserClientType) String() string {
	return proto.EnumName(UserClientType_name, int32(x))
}
func (UserClientType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_net_3bb38ce5f35e60f5, []int{3}
}

type HttpQueryType int32

const (
	HttpQueryType_HttpQueryInvalid HttpQueryType = 0
	HttpQueryType_HttpQueryGet     HttpQueryType = 1
	HttpQueryType_HttpQueryPost    HttpQueryType = 2
)

var HttpQueryType_name = map[int32]string{
	0: "HttpQueryInvalid",
	1: "HttpQueryGet",
	2: "HttpQueryPost",
}
var HttpQueryType_value = map[string]int32{
	"HttpQueryInvalid": 0,
	"HttpQueryGet":     1,
	"HttpQueryPost":    2,
}

func (x HttpQueryType) String() string {
	return proto.EnumName(HttpQueryType_name, int32(x))
}
func (HttpQueryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_net_3bb38ce5f35e60f5, []int{4}
}

// there are lots of mime.types here:
// http://svn.apache.org/repos/asf/httpd/httpd/trunk/docs/conf/mime.types
type FileType int32

const (
	FileType_FileInvalid FileType = 0
	FileType_FileImage   FileType = 1
	FileType_FileAudio   FileType = 2
)

var FileType_name = map[int32]string{
	0: "FileInvalid",
	1: "FileImage",
	2: "FileAudio",
}
var FileType_value = map[string]int32{
	"FileInvalid": 0,
	"FileImage":   1,
	"FileAudio":   2,
}

func (x FileType) String() string {
	return proto.EnumName(FileType_name, int32(x))
}
func (FileType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_net_3bb38ce5f35e60f5, []int{5}
}

type DataWriteType int32

const (
	// db = requestValue
	DataWriteType_WriteUpdate DataWriteType = 0
	// db = db + requestValue
	DataWriteType_WriteAdd DataWriteType = 1
	// db = db - requestValue
	DataWriteType_WriteDel DataWriteType = 2
)

var DataWriteType_name = map[int32]string{
	0: "WriteUpdate",
	1: "WriteAdd",
	2: "WriteDel",
}
var DataWriteType_value = map[string]int32{
	"WriteUpdate": 0,
	"WriteAdd":    1,
	"WriteDel":    2,
}

func (x DataWriteType) String() string {
	return proto.EnumName(DataWriteType_name, int32(x))
}
func (DataWriteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_net_3bb38ce5f35e60f5, []int{6}
}

type TransportData struct {
	Action string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Body   *any.Any `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	// transfrom TransportDataHeaderKey to string.
	//
	// key = "_" + TransportDataHeaderKey
	// ex: "_1", "_9"
	//
	// Q: why donot use TransportDataHeaderKey directly?
	// A: javascript(json) doesnot support int key in object.
	//
	// Q: why donnot use string(TransportDataHeaderKey)?
	// A: php treat a string with a int pattern as an int, so the probuf cannot mergeFromJsonString.
	//
	Header map[string]string `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// uniqid for blockRequest
	// the default php config doesnot support bcmath required by int64
	PackageId            int64    `protobuf:"varint,4,opt,name=packageId,proto3" json:"packageId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransportData) Reset()         { *m = TransportData{} }
func (m *TransportData) String() string { return proto.CompactTextString(m) }
func (*TransportData) ProtoMessage()    {}
func (*TransportData) Descriptor() ([]byte, []int) {
	return fileDescriptor_net_3bb38ce5f35e60f5, []int{0}
}
func (m *TransportData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransportData.Unmarshal(m, b)
}
func (m *TransportData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransportData.Marshal(b, m, deterministic)
}
func (dst *TransportData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransportData.Merge(dst, src)
}
func (m *TransportData) XXX_Size() int {
	return xxx_messageInfo_TransportData.Size(m)
}
func (m *TransportData) XXX_DiscardUnknown() {
	xxx_messageInfo_TransportData.DiscardUnknown(m)
}

var xxx_messageInfo_TransportData proto.InternalMessageInfo

func (m *TransportData) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *TransportData) GetBody() *any.Any {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *TransportData) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *TransportData) GetPackageId() int64 {
	if m != nil {
		return m.PackageId
	}
	return 0
}

func init() {
	proto.RegisterType((*TransportData)(nil), "core.TransportData")
	proto.RegisterMapType((map[string]string)(nil), "core.TransportData.HeaderEntry")
	proto.RegisterEnum("core.TransportDataHeaderKey", TransportDataHeaderKey_name, TransportDataHeaderKey_value)
	proto.RegisterEnum("core.UserClientLangType", UserClientLangType_name, UserClientLangType_value)
	proto.RegisterEnum("core.GatewayType", GatewayType_name, GatewayType_value)
	proto.RegisterEnum("core.UserClientType", UserClientType_name, UserClientType_value)
	proto.RegisterEnum("core.HttpQueryType", HttpQueryType_name, HttpQueryType_value)
	proto.RegisterEnum("core.FileType", FileType_name, FileType_value)
	proto.RegisterEnum("core.DataWriteType", DataWriteType_name, DataWriteType_value)
}

func init() { proto.RegisterFile("core/net.proto", fileDescriptor_net_3bb38ce5f35e60f5) }

var fileDescriptor_net_3bb38ce5f35e60f5 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0xdd, 0x4e, 0xdb, 0x4e,
	0x10, 0xc5, 0xb1, 0x13, 0xf2, 0x27, 0x13, 0x42, 0x86, 0x25, 0x82, 0xfc, 0x29, 0x52, 0xa3, 0x5e,
	0x45, 0x96, 0xea, 0x48, 0xf4, 0xa2, 0xa5, 0xea, 0x45, 0xd3, 0xf0, 0x11, 0xfa, 0x25, 0x9a, 0x42,
	0x69, 0x11, 0x37, 0x6b, 0x7b, 0x70, 0xac, 0x38, 0x5e, 0x6b, 0xbd, 0x01, 0xb9, 0x8f, 0xd8, 0xd7,
	0xe8, 0x23, 0xf4, 0x05, 0x2a, 0xaf, 0x4d, 0x4c, 0xe8, 0x9d, 0xcf, 0x6f, 0xe7, 0x9c, 0x99, 0x4c,
	0x76, 0x61, 0xc3, 0x15, 0x92, 0xfa, 0x11, 0x29, 0x3b, 0x96, 0x42, 0x09, 0x56, 0xcd, 0xf4, 0xee,
	0xff, 0xbe, 0x10, 0x7e, 0x48, 0x7d, 0xcd, 0x9c, 0xf9, 0x4d, 0x9f, 0x47, 0x69, 0x5e, 0xf0, 0xec,
	0xb7, 0x01, 0xcd, 0x73, 0xc9, 0xa3, 0x24, 0x16, 0x52, 0x1d, 0x72, 0xc5, 0xd9, 0x36, 0xd4, 0xb8,
	0xab, 0x02, 0x11, 0x75, 0x8c, 0xae, 0xd1, 0xab, 0x8f, 0x0b, 0xc5, 0x7a, 0x50, 0x75, 0x84, 0x97,
	0x76, 0xcc, 0xae, 0xd1, 0x6b, 0xec, 0xb7, 0xed, 0x3c, 0xd3, 0xbe, 0xcf, 0xb4, 0x07, 0x51, 0x3a,
	0xd6, 0x15, 0xec, 0x25, 0xd4, 0x26, 0xc4, 0x3d, 0x92, 0x9d, 0x4a, 0xb7, 0xd2, 0x6b, 0xec, 0x3f,
	0xb5, 0xb3, 0x29, 0xec, 0xa5, 0x36, 0xf6, 0x48, 0x57, 0x1c, 0x45, 0x4a, 0xa6, 0xe3, 0xa2, 0x9c,
	0xed, 0x41, 0x3d, 0xe6, 0xee, 0x94, 0xfb, 0x74, 0xea, 0x75, 0xaa, 0x5d, 0xa3, 0x57, 0x19, 0x97,
	0x60, 0xf7, 0x00, 0x1a, 0x0f, 0x4c, 0x0c, 0xa1, 0x32, 0xa5, 0xb4, 0x18, 0x32, 0xfb, 0x64, 0x6d,
	0x58, 0xbd, 0xe5, 0xe1, 0x9c, 0xf4, 0x88, 0xf5, 0x71, 0x2e, 0x5e, 0x9b, 0xaf, 0x0c, 0xeb, 0x8f,
	0x01, 0xdb, 0x4b, 0xed, 0xf3, 0xa0, 0x0f, 0x94, 0xb2, 0x4d, 0x68, 0xe6, 0xe2, 0x34, 0xba, 0xe5,
	0x61, 0xe0, 0xe1, 0x0a, 0xdb, 0x82, 0x56, 0xd1, 0x48, 0x4a, 0x21, 0x87, 0xc2, 0x23, 0x34, 0x1e,
	0xc1, 0xd3, 0xe8, 0x46, 0xa0, 0x59, 0xc2, 0xaf, 0x94, 0x24, 0x81, 0x88, 0x02, 0x0f, 0x2b, 0x65,
	0xe2, 0x48, 0x24, 0xea, 0x42, 0x86, 0x58, 0x2d, 0xd1, 0x98, 0x6e, 0x48, 0x92, 0xc4, 0xd5, 0xd2,
	0x7a, 0x91, 0x90, 0x1c, 0xf8, 0x14, 0x29, 0xac, 0xb1, 0x36, 0x60, 0x0e, 0x07, 0x61, 0x28, 0xee,
	0x86, 0xdc, 0x9d, 0x10, 0xfe, 0xc7, 0x3a, 0xd0, 0x2e, 0x4b, 0x87, 0x61, 0x40, 0x91, 0xfa, 0xc8,
	0x23, 0x1f, 0xd7, 0xd8, 0x1e, 0x74, 0x8a, 0xfa, 0x38, 0x0e, 0x03, 0x97, 0x67, 0x7f, 0xd4, 0x37,
	0x92, 0xd9, 0x28, 0x08, 0xd6, 0x5b, 0x60, 0xcb, 0x8e, 0xf3, 0x34, 0xa6, 0xac, 0xc7, 0x32, 0x3d,
	0xfa, 0x8c, 0x2b, 0xff, 0xd2, 0xab, 0x11, 0x1a, 0xd6, 0x08, 0x1a, 0x27, 0x5c, 0xd1, 0x1d, 0x4f,
	0xb5, 0x95, 0xc1, 0x46, 0x21, 0xcb, 0x65, 0xb5, 0x01, 0x0b, 0x76, 0x49, 0x4e, 0x22, 0xdc, 0x29,
	0x29, 0x34, 0x58, 0x6b, 0x61, 0xbc, 0xe2, 0x61, 0x8a, 0xa6, 0xf5, 0x03, 0x36, 0xca, 0x7c, 0x1d,
	0xf6, 0x04, 0x76, 0x4a, 0xf2, 0x49, 0x38, 0x41, 0x48, 0x65, 0xea, 0x0e, 0x6c, 0x3d, 0x3e, 0x1c,
	0xc4, 0x31, 0x1a, 0xd9, 0x26, 0xcb, 0x83, 0x4b, 0x72, 0xd0, 0xb4, 0xde, 0x43, 0x73, 0xa4, 0x54,
	0xfc, 0x65, 0x4e, 0x32, 0xbd, 0xff, 0x85, 0x0b, 0x50, 0x46, 0x22, 0xac, 0x2f, 0xe8, 0x89, 0x1e,
	0x72, 0xf3, 0x81, 0xf1, 0x4c, 0x24, 0x0a, 0x4d, 0xeb, 0x00, 0xd6, 0x8e, 0x83, 0x90, 0x74, 0x4c,
	0x0b, 0x1a, 0xc7, 0x4b, 0x43, 0x35, 0xa1, 0xae, 0xc1, 0x8c, 0xfb, 0xd9, 0x8d, 0x28, 0xe4, 0x60,
	0xee, 0x05, 0x02, 0x4d, 0xeb, 0x0d, 0x34, 0xb3, 0x9b, 0x75, 0x29, 0x03, 0xb5, 0xf0, 0x6b, 0x71,
	0x11, 0x7b, 0x5c, 0x11, 0xae, 0xb0, 0x75, 0x58, 0xd3, 0x60, 0xe0, 0x79, 0x68, 0x2c, 0xd4, 0x21,
	0x85, 0x68, 0xbe, 0xfb, 0x0e, 0x5b, 0xae, 0x98, 0xd9, 0x3f, 0x79, 0x58, 0xbc, 0x4c, 0xfd, 0x66,
	0xae, 0xfa, 0x7e, 0xa0, 0x26, 0x73, 0xc7, 0x76, 0xc5, 0xac, 0xef, 0xcd, 0xdd, 0xa9, 0x3b, 0xe1,
	0x6a, 0xf1, 0xf1, 0xdc, 0xcf, 0x57, 0x9c, 0xbf, 0xeb, 0x7e, 0x66, 0xf8, 0x65, 0xb6, 0xb2, 0x85,
	0x5f, 0x9f, 0x65, 0xe4, 0x7a, 0x28, 0x24, 0x39, 0x35, 0x7d, 0xfa, 0xe2, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x75, 0x72, 0x01, 0xa4, 0x1b, 0x04, 0x00, 0x00,
}
