// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: Qot_GetHistoryKL.proto

package qotgethistorykl

import (
	_ "github.com/Pai-Po/go-futu-api/pb/common"
	qotcommon "github.com/Pai-Po/go-futu-api/pb/qotcommon"
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

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RehabType        *int32              `protobuf:"varint,1,req,name=rehabType" json:"rehabType,omitempty"`               //Qot_Common.RehabType,复权类型
	KlType           *int32              `protobuf:"varint,2,req,name=klType" json:"klType,omitempty"`                     //Qot_Common.KLType,K线类型
	Security         *qotcommon.Security `protobuf:"bytes,3,req,name=security" json:"security,omitempty"`                  //股票市场以及股票代码
	BeginTime        *string             `protobuf:"bytes,4,req,name=beginTime" json:"beginTime,omitempty"`                //开始时间字符串
	EndTime          *string             `protobuf:"bytes,5,req,name=endTime" json:"endTime,omitempty"`                    //结束时间字符串
	MaxAckKLNum      *int32              `protobuf:"varint,6,opt,name=maxAckKLNum" json:"maxAckKLNum,omitempty"`           //最多返回多少根K线，如果未指定表示不限制
	NeedKLFieldsFlag *int64              `protobuf:"varint,7,opt,name=needKLFieldsFlag" json:"needKLFieldsFlag,omitempty"` //指定返回K线结构体特定某几项数据，KLFields枚举值或组合，如果未指定返回全部字段
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetHistoryKL_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetHistoryKL_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S.ProtoReflect.Descriptor instead.
func (*C2S) Descriptor() ([]byte, []int) {
	return file_Qot_GetHistoryKL_proto_rawDescGZIP(), []int{0}
}

func (x *C2S) GetRehabType() int32 {
	if x != nil && x.RehabType != nil {
		return *x.RehabType
	}
	return 0
}

func (x *C2S) GetKlType() int32 {
	if x != nil && x.KlType != nil {
		return *x.KlType
	}
	return 0
}

func (x *C2S) GetSecurity() *qotcommon.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *C2S) GetBeginTime() string {
	if x != nil && x.BeginTime != nil {
		return *x.BeginTime
	}
	return ""
}

func (x *C2S) GetEndTime() string {
	if x != nil && x.EndTime != nil {
		return *x.EndTime
	}
	return ""
}

func (x *C2S) GetMaxAckKLNum() int32 {
	if x != nil && x.MaxAckKLNum != nil {
		return *x.MaxAckKLNum
	}
	return 0
}

func (x *C2S) GetNeedKLFieldsFlag() int64 {
	if x != nil && x.NeedKLFieldsFlag != nil {
		return *x.NeedKLFieldsFlag
	}
	return 0
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Security        *qotcommon.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	KlList          []*qotcommon.KLine  `protobuf:"bytes,2,rep,name=klList" json:"klList,omitempty"`                     //K线数据
	NextKLTime      *string             `protobuf:"bytes,3,opt,name=nextKLTime" json:"nextKLTime,omitempty"`             //如请求不指定maxAckKLNum值，则不会返回该字段，该字段表示超过指定限制的下一K线时间字符串
	NextKLTimestamp *float64            `protobuf:"fixed64,4,opt,name=nextKLTimestamp" json:"nextKLTimestamp,omitempty"` //时间戳，如请求不指定maxAckKLNum值，则不会返回该字段，该字段表示超过指定限制的下一K线时间戳
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetHistoryKL_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetHistoryKL_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C.ProtoReflect.Descriptor instead.
func (*S2C) Descriptor() ([]byte, []int) {
	return file_Qot_GetHistoryKL_proto_rawDescGZIP(), []int{1}
}

func (x *S2C) GetSecurity() *qotcommon.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *S2C) GetKlList() []*qotcommon.KLine {
	if x != nil {
		return x.KlList
	}
	return nil
}

func (x *S2C) GetNextKLTime() string {
	if x != nil && x.NextKLTime != nil {
		return *x.NextKLTime
	}
	return ""
}

func (x *S2C) GetNextKLTimestamp() float64 {
	if x != nil && x.NextKLTimestamp != nil {
		return *x.NextKLTimestamp
	}
	return 0
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetHistoryKL_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetHistoryKL_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_Qot_GetHistoryKL_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetC2S() *C2S {
	if x != nil {
		return x.C2S
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"` //RetType,返回结果
	RetMsg  *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C     *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

// Default values for Response fields.
const (
	Default_Response_RetType = int32(-400)
)

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetHistoryKL_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetHistoryKL_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_Qot_GetHistoryKL_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetRetType() int32 {
	if x != nil && x.RetType != nil {
		return *x.RetType
	}
	return Default_Response_RetType
}

func (x *Response) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *Response) GetErrCode() int32 {
	if x != nil && x.ErrCode != nil {
		return *x.ErrCode
	}
	return 0
}

func (x *Response) GetS2C() *S2C {
	if x != nil {
		return x.S2C
	}
	return nil
}

var File_Qot_GetHistoryKL_proto protoreflect.FileDescriptor

var file_Qot_GetHistoryKL_proto_rawDesc = []byte{
	0x0a, 0x16, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x4b, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x4c, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x03, 0x43,
	0x32, 0x53, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x68, 0x61, 0x62, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x68, 0x61, 0x62, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6b, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05,
	0x52, 0x06, 0x6b, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f, 0x74,
	0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x41, 0x63, 0x6b, 0x4b, 0x4c, 0x4e, 0x75,
	0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x41, 0x63, 0x6b, 0x4b,
	0x4c, 0x4e, 0x75, 0x6d, 0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x65, 0x65, 0x64, 0x4b, 0x4c, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10,
	0x6e, 0x65, 0x65, 0x64, 0x4b, 0x4c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x46, 0x6c, 0x61, 0x67,
	0x22, 0xac, 0x01, 0x0a, 0x03, 0x53, 0x32, 0x43, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f, 0x74,
	0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x06, 0x6b, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x51, 0x6f, 0x74,
	0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4b, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x06, 0x6b,
	0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x4b, 0x4c, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x4b,
	0x4c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x4b, 0x4c, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x4b, 0x4c, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x32, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x03, 0x63, 0x32,
	0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x4c, 0x2e, 0x43, 0x32, 0x53, 0x52, 0x03,
	0x63, 0x32, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x05, 0x3a, 0x04, 0x2d, 0x34, 0x30, 0x30, 0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x73, 0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x4b, 0x4c, 0x2e, 0x53, 0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x42, 0x4a, 0x0a, 0x13, 0x63,
	0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x74, 0x75, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x62, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x75, 0x72, 0x69, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x66, 0x75, 0x74, 0x75,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x71, 0x6f, 0x74, 0x67, 0x65, 0x74, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x6b, 0x6c,
}

var (
	file_Qot_GetHistoryKL_proto_rawDescOnce sync.Once
	file_Qot_GetHistoryKL_proto_rawDescData = file_Qot_GetHistoryKL_proto_rawDesc
)

func file_Qot_GetHistoryKL_proto_rawDescGZIP() []byte {
	file_Qot_GetHistoryKL_proto_rawDescOnce.Do(func() {
		file_Qot_GetHistoryKL_proto_rawDescData = protoimpl.X.CompressGZIP(file_Qot_GetHistoryKL_proto_rawDescData)
	})
	return file_Qot_GetHistoryKL_proto_rawDescData
}

var file_Qot_GetHistoryKL_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Qot_GetHistoryKL_proto_goTypes = []interface{}{
	(*C2S)(nil),                // 0: Qot_GetHistoryKL.C2S
	(*S2C)(nil),                // 1: Qot_GetHistoryKL.S2C
	(*Request)(nil),            // 2: Qot_GetHistoryKL.Request
	(*Response)(nil),           // 3: Qot_GetHistoryKL.Response
	(*qotcommon.Security)(nil), // 4: Qot_Common.Security
	(*qotcommon.KLine)(nil),    // 5: Qot_Common.KLine
}
var file_Qot_GetHistoryKL_proto_depIdxs = []int32{
	4, // 0: Qot_GetHistoryKL.C2S.security:type_name -> Qot_Common.Security
	4, // 1: Qot_GetHistoryKL.S2C.security:type_name -> Qot_Common.Security
	5, // 2: Qot_GetHistoryKL.S2C.klList:type_name -> Qot_Common.KLine
	0, // 3: Qot_GetHistoryKL.Request.c2s:type_name -> Qot_GetHistoryKL.C2S
	1, // 4: Qot_GetHistoryKL.Response.s2c:type_name -> Qot_GetHistoryKL.S2C
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_Qot_GetHistoryKL_proto_init() }
func file_Qot_GetHistoryKL_proto_init() {
	if File_Qot_GetHistoryKL_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Qot_GetHistoryKL_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S); i {
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
		file_Qot_GetHistoryKL_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C); i {
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
		file_Qot_GetHistoryKL_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_Qot_GetHistoryKL_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_Qot_GetHistoryKL_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Qot_GetHistoryKL_proto_goTypes,
		DependencyIndexes: file_Qot_GetHistoryKL_proto_depIdxs,
		MessageInfos:      file_Qot_GetHistoryKL_proto_msgTypes,
	}.Build()
	File_Qot_GetHistoryKL_proto = out.File
	file_Qot_GetHistoryKL_proto_rawDesc = nil
	file_Qot_GetHistoryKL_proto_goTypes = nil
	file_Qot_GetHistoryKL_proto_depIdxs = nil
}
