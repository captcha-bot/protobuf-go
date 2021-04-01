// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: discord/model.proto

package model

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

type MessageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64  `protobuf:"fixed64,1,opt,name=id,proto3" json:"id,omitempty"`
	ChannelId uint64  `protobuf:"fixed64,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Content   string  `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Author    uint64  `protobuf:"fixed64,7,opt,name=author,proto3" json:"author,omitempty"`
	GuildId   *uint64 `protobuf:"fixed64,3,opt,name=guild_id,json=guildId,proto3,oneof" json:"guild_id,omitempty"`
	Member    *uint64 `protobuf:"fixed64,8,opt,name=member,proto3,oneof" json:"member,omitempty"`
}

func (x *MessageData) Reset() {
	*x = MessageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discord_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageData) ProtoMessage() {}

func (x *MessageData) ProtoReflect() protoreflect.Message {
	mi := &file_discord_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageData.ProtoReflect.Descriptor instead.
func (*MessageData) Descriptor() ([]byte, []int) {
	return file_discord_model_proto_rawDescGZIP(), []int{0}
}

func (x *MessageData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MessageData) GetChannelId() uint64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *MessageData) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MessageData) GetAuthor() uint64 {
	if x != nil {
		return x.Author
	}
	return 0
}

func (x *MessageData) GetGuildId() uint64 {
	if x != nil && x.GuildId != nil {
		return *x.GuildId
	}
	return 0
}

func (x *MessageData) GetMember() uint64 {
	if x != nil && x.Member != nil {
		return *x.Member
	}
	return 0
}

var File_discord_model_proto protoreflect.FileDescriptor

var file_discord_model_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xd7, 0x01, 0x0a,
	0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x42, 0x02, 0x30, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x21, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x06, 0x42, 0x02, 0x30, 0x01, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x06, 0x42, 0x02, 0x30,
	0x01, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x08, 0x67, 0x75, 0x69,
	0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x06, 0x42, 0x02, 0x30, 0x01, 0x48,
	0x00, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a,
	0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x06, 0x42, 0x02, 0x30,
	0x01, 0x48, 0x01, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2d, 0x62, 0x6f, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x72, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_discord_model_proto_rawDescOnce sync.Once
	file_discord_model_proto_rawDescData = file_discord_model_proto_rawDesc
)

func file_discord_model_proto_rawDescGZIP() []byte {
	file_discord_model_proto_rawDescOnce.Do(func() {
		file_discord_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_discord_model_proto_rawDescData)
	})
	return file_discord_model_proto_rawDescData
}

var file_discord_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_discord_model_proto_goTypes = []interface{}{
	(*MessageData)(nil), // 0: captcha.discord.model.MessageData
}
var file_discord_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_discord_model_proto_init() }
func file_discord_model_proto_init() {
	if File_discord_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_discord_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageData); i {
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
	file_discord_model_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_discord_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_discord_model_proto_goTypes,
		DependencyIndexes: file_discord_model_proto_depIdxs,
		MessageInfos:      file_discord_model_proto_msgTypes,
	}.Build()
	File_discord_model_proto = out.File
	file_discord_model_proto_rawDesc = nil
	file_discord_model_proto_goTypes = nil
	file_discord_model_proto_depIdxs = nil
}
