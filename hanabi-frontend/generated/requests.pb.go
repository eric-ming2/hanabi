// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: requests.proto

package generated

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

type RequestType int32

const (
	RequestType_START_GAME   RequestType = 0
	RequestType_DISCARD_CARD RequestType = 1
	RequestType_PLAY_CARD    RequestType = 2
	RequestType_GIVE_HINT    RequestType = 3
)

// Enum value maps for RequestType.
var (
	RequestType_name = map[int32]string{
		0: "START_GAME",
		1: "DISCARD_CARD",
		2: "PLAY_CARD",
		3: "GIVE_HINT",
	}
	RequestType_value = map[string]int32{
		"START_GAME":   0,
		"DISCARD_CARD": 1,
		"PLAY_CARD":    2,
		"GIVE_HINT":    3,
	}
)

func (x RequestType) Enum() *RequestType {
	p := new(RequestType)
	*p = x
	return p
}

func (x RequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_requests_proto_enumTypes[0].Descriptor()
}

func (RequestType) Type() protoreflect.EnumType {
	return &file_requests_proto_enumTypes[0]
}

func (x RequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestType.Descriptor instead.
func (RequestType) EnumDescriptor() ([]byte, []int) {
	return file_requests_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestType RequestType `protobuf:"varint,1,opt,name=request_type,json=requestType,proto3,enum=requests.RequestType" json:"request_type,omitempty"`
	// Types that are assignable to Request:
	//
	//	*Request_StartGame
	//	*Request_DiscardCard
	//	*Request_PlayCard
	//	*Request_GiveHint
	Request isRequest_Request `protobuf_oneof:"request"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_requests_proto_msgTypes[0]
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
	return file_requests_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetRequestType() RequestType {
	if x != nil {
		return x.RequestType
	}
	return RequestType_START_GAME
}

func (m *Request) GetRequest() isRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *Request) GetStartGame() *StartGameRequest {
	if x, ok := x.GetRequest().(*Request_StartGame); ok {
		return x.StartGame
	}
	return nil
}

func (x *Request) GetDiscardCard() *DiscardCardRequest {
	if x, ok := x.GetRequest().(*Request_DiscardCard); ok {
		return x.DiscardCard
	}
	return nil
}

func (x *Request) GetPlayCard() *PlayCardRequest {
	if x, ok := x.GetRequest().(*Request_PlayCard); ok {
		return x.PlayCard
	}
	return nil
}

func (x *Request) GetGiveHint() *GiveHintRequest {
	if x, ok := x.GetRequest().(*Request_GiveHint); ok {
		return x.GiveHint
	}
	return nil
}

type isRequest_Request interface {
	isRequest_Request()
}

type Request_StartGame struct {
	StartGame *StartGameRequest `protobuf:"bytes,2,opt,name=start_game,json=startGame,proto3,oneof"`
}

type Request_DiscardCard struct {
	DiscardCard *DiscardCardRequest `protobuf:"bytes,3,opt,name=discard_card,json=discardCard,proto3,oneof"`
}

type Request_PlayCard struct {
	PlayCard *PlayCardRequest `protobuf:"bytes,4,opt,name=play_card,json=playCard,proto3,oneof"`
}

type Request_GiveHint struct {
	GiveHint *GiveHintRequest `protobuf:"bytes,5,opt,name=give_hint,json=giveHint,proto3,oneof"`
}

func (*Request_StartGame) isRequest_Request() {}

func (*Request_DiscardCard) isRequest_Request() {}

func (*Request_PlayCard) isRequest_Request() {}

func (*Request_GiveHint) isRequest_Request() {}

type StartGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartGameRequest) Reset() {
	*x = StartGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartGameRequest) ProtoMessage() {}

func (x *StartGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requests_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartGameRequest.ProtoReflect.Descriptor instead.
func (*StartGameRequest) Descriptor() ([]byte, []int) {
	return file_requests_proto_rawDescGZIP(), []int{1}
}

type DiscardCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardIndex int32 `protobuf:"varint,1,opt,name=card_index,json=cardIndex,proto3" json:"card_index,omitempty"`
}

func (x *DiscardCardRequest) Reset() {
	*x = DiscardCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscardCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscardCardRequest) ProtoMessage() {}

func (x *DiscardCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requests_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscardCardRequest.ProtoReflect.Descriptor instead.
func (*DiscardCardRequest) Descriptor() ([]byte, []int) {
	return file_requests_proto_rawDescGZIP(), []int{2}
}

func (x *DiscardCardRequest) GetCardIndex() int32 {
	if x != nil {
		return x.CardIndex
	}
	return 0
}

type PlayCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardIndex int32 `protobuf:"varint,1,opt,name=card_index,json=cardIndex,proto3" json:"card_index,omitempty"`
}

func (x *PlayCardRequest) Reset() {
	*x = PlayCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayCardRequest) ProtoMessage() {}

func (x *PlayCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requests_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayCardRequest.ProtoReflect.Descriptor instead.
func (*PlayCardRequest) Descriptor() ([]byte, []int) {
	return file_requests_proto_rawDescGZIP(), []int{3}
}

func (x *PlayCardRequest) GetCardIndex() int32 {
	if x != nil {
		return x.CardIndex
	}
	return 0
}

type GiveHintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardIndex   int32 `protobuf:"varint,1,opt,name=card_index,json=cardIndex,proto3" json:"card_index,omitempty"`
	PlayerIndex int32 `protobuf:"varint,2,opt,name=player_index,json=playerIndex,proto3" json:"player_index,omitempty"`
}

func (x *GiveHintRequest) Reset() {
	*x = GiveHintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiveHintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiveHintRequest) ProtoMessage() {}

func (x *GiveHintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requests_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiveHintRequest.ProtoReflect.Descriptor instead.
func (*GiveHintRequest) Descriptor() ([]byte, []int) {
	return file_requests_proto_rawDescGZIP(), []int{4}
}

func (x *GiveHintRequest) GetCardIndex() int32 {
	if x != nil {
		return x.CardIndex
	}
	return 0
}

func (x *GiveHintRequest) GetPlayerIndex() int32 {
	if x != nil {
		return x.PlayerIndex
	}
	return 0
}

var File_requests_proto protoreflect.FileDescriptor

var file_requests_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0xc2, 0x02, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x3b, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x43, 0x61, 0x72, 0x64,
	0x12, 0x38, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x61, 0x72, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x67, 0x69,
	0x76, 0x65, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x69, 0x76, 0x65, 0x48, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x67, 0x69, 0x76, 0x65,
	0x48, 0x69, 0x6e, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x12, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x72,
	0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63,
	0x61, 0x72, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x30, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x61, 0x72, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x63, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x53, 0x0a, 0x0f, 0x47, 0x69,
	0x76, 0x65, 0x48, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x63, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2a,
	0x4d, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e,
	0x0a, 0x0a, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x44, 0x49, 0x53, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x02, 0x12,
	0x0d, 0x0a, 0x09, 0x47, 0x49, 0x56, 0x45, 0x5f, 0x48, 0x49, 0x4e, 0x54, 0x10, 0x03, 0x42, 0x2e,
	0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x69,
	0x63, 0x2d, 0x6d, 0x69, 0x6e, 0x67, 0x32, 0x2f, 0x68, 0x61, 0x6e, 0x61, 0x62, 0x69, 0x2f, 0x68,
	0x61, 0x6e, 0x61, 0x62, 0x69, 0x2d, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_requests_proto_rawDescOnce sync.Once
	file_requests_proto_rawDescData = file_requests_proto_rawDesc
)

func file_requests_proto_rawDescGZIP() []byte {
	file_requests_proto_rawDescOnce.Do(func() {
		file_requests_proto_rawDescData = protoimpl.X.CompressGZIP(file_requests_proto_rawDescData)
	})
	return file_requests_proto_rawDescData
}

var file_requests_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_requests_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_requests_proto_goTypes = []any{
	(RequestType)(0),           // 0: requests.RequestType
	(*Request)(nil),            // 1: requests.Request
	(*StartGameRequest)(nil),   // 2: requests.StartGameRequest
	(*DiscardCardRequest)(nil), // 3: requests.DiscardCardRequest
	(*PlayCardRequest)(nil),    // 4: requests.PlayCardRequest
	(*GiveHintRequest)(nil),    // 5: requests.GiveHintRequest
}
var file_requests_proto_depIdxs = []int32{
	0, // 0: requests.Request.request_type:type_name -> requests.RequestType
	2, // 1: requests.Request.start_game:type_name -> requests.StartGameRequest
	3, // 2: requests.Request.discard_card:type_name -> requests.DiscardCardRequest
	4, // 3: requests.Request.play_card:type_name -> requests.PlayCardRequest
	5, // 4: requests.Request.give_hint:type_name -> requests.GiveHintRequest
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_requests_proto_init() }
func file_requests_proto_init() {
	if File_requests_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_requests_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_requests_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StartGameRequest); i {
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
		file_requests_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DiscardCardRequest); i {
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
		file_requests_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PlayCardRequest); i {
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
		file_requests_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GiveHintRequest); i {
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
	file_requests_proto_msgTypes[0].OneofWrappers = []any{
		(*Request_StartGame)(nil),
		(*Request_DiscardCard)(nil),
		(*Request_PlayCard)(nil),
		(*Request_GiveHint)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_requests_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_requests_proto_goTypes,
		DependencyIndexes: file_requests_proto_depIdxs,
		EnumInfos:         file_requests_proto_enumTypes,
		MessageInfos:      file_requests_proto_msgTypes,
	}.Build()
	File_requests_proto = out.File
	file_requests_proto_rawDesc = nil
	file_requests_proto_goTypes = nil
	file_requests_proto_depIdxs = nil
}
