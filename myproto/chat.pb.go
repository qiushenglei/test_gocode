// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: chat.proto

// package  proto文件的包名(与go无关)  生成的语言类放在什么包里面  这个package必须紧跟着syntax="proto3"

package myproto

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeqId     int64  `protobuf:"varint,1,opt,name=seqId,proto3" json:"seqId,omitempty"`
	Sender    int64  `protobuf:"varint,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver  int64  `protobuf:"varint,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Forwarder int64  `protobuf:"varint,4,opt,name=forwarder,proto3" json:"forwarder,omitempty"`
	GroupId   int64  `protobuf:"varint,5,opt,name=groupId,proto3" json:"groupId,omitempty"`
	MsgType   int64  `protobuf:"varint,6,opt,name=msgType,proto3" json:"msgType,omitempty"`
	MsgFlag   int64  `protobuf:"varint,7,opt,name=msgFlag,proto3" json:"msgFlag,omitempty"`
	Timestamp int64  `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Mid       string `protobuf:"bytes,9,opt,name=mid,proto3" json:"mid,omitempty"`
	ReplyId   string `protobuf:"bytes,10,opt,name=replyId,proto3" json:"replyId,omitempty"`
	Data      string `protobuf:"bytes,11,opt,name=data,proto3" json:"data,omitempty"`
	Version   string `protobuf:"bytes,12,opt,name=version,proto3" json:"version,omitempty"`
	Uuid      string `protobuf:"bytes,13,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Nickname  string `protobuf:"bytes,14,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar    string `protobuf:"bytes,15,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Qid       string `protobuf:"bytes,16,opt,name=qid,proto3" json:"qid,omitempty"`
	GlobalId  int64  `protobuf:"varint,17,opt,name=globalId,proto3" json:"globalId,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetSeqId() int64 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

func (x *Message) GetSender() int64 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *Message) GetReceiver() int64 {
	if x != nil {
		return x.Receiver
	}
	return 0
}

func (x *Message) GetForwarder() int64 {
	if x != nil {
		return x.Forwarder
	}
	return 0
}

func (x *Message) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *Message) GetMsgType() int64 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *Message) GetMsgFlag() int64 {
	if x != nil {
		return x.MsgFlag
	}
	return 0
}

func (x *Message) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Message) GetMid() string {
	if x != nil {
		return x.Mid
	}
	return ""
}

func (x *Message) GetReplyId() string {
	if x != nil {
		return x.ReplyId
	}
	return ""
}

func (x *Message) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Message) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Message) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Message) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Message) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Message) GetQid() string {
	if x != nil {
		return x.Qid
	}
	return ""
}

func (x *Message) GetGlobalId() int64 {
	if x != nil {
		return x.GlobalId
	}
	return 0
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x22, 0xad, 0x03, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73,
	0x65, 0x71, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73,
	0x67, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x73, 0x67,
	0x46, 0x6c, 0x61, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x71, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x49, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chat_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: chat.Message
}
var file_chat_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
