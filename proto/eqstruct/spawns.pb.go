// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: proto/eqstruct/spawns.proto

package eqstruct

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

type Spawn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpawnId      uint32  `protobuf:"varint,1,opt,name=spawn_id,json=spawnId,proto3" json:"spawn_id,omitempty"`
	Acceleration int32   `protobuf:"varint,2,opt,name=acceleration,proto3" json:"acceleration,omitempty"`
	Heading      uint32  `protobuf:"varint,3,opt,name=heading,proto3" json:"heading,omitempty"`
	HeadingDelta int32   `protobuf:"varint,4,opt,name=heading_delta,json=headingDelta,proto3" json:"heading_delta,omitempty"`
	Y            int32   `protobuf:"varint,5,opt,name=y,proto3" json:"y,omitempty"`
	X            int32   `protobuf:"varint,6,opt,name=x,proto3" json:"x,omitempty"`
	Z            int32   `protobuf:"varint,7,opt,name=z,proto3" json:"z,omitempty"`
	YDelta       uint32  `protobuf:"varint,8,opt,name=y_delta,json=yDelta,proto3" json:"y_delta,omitempty"`
	XDelta       uint32  `protobuf:"varint,9,opt,name=x_delta,json=xDelta,proto3" json:"x_delta,omitempty"`
	ZDelta       uint32  `protobuf:"varint,10,opt,name=z_delta,json=zDelta,proto3" json:"z_delta,omitempty"`
	PetOwner     uint32  `protobuf:"varint,11,opt,name=pet_owner,json=petOwner,proto3" json:"pet_owner,omitempty"`
	Animation    uint32  `protobuf:"varint,12,opt,name=animation,proto3" json:"animation,omitempty"`
	HairColor    uint32  `protobuf:"varint,13,opt,name=hair_color,json=hairColor,proto3" json:"hair_color,omitempty"`
	BeardColor   uint32  `protobuf:"varint,14,opt,name=beard_color,json=beardColor,proto3" json:"beard_color,omitempty"`
	EyeColor_1   uint32  `protobuf:"varint,15,opt,name=eye_color_1,json=eyeColor1,proto3" json:"eye_color_1,omitempty"`
	EyeColor_2   uint32  `protobuf:"varint,16,opt,name=eye_color_2,json=eyeColor2,proto3" json:"eye_color_2,omitempty"`
	HairStyle    uint32  `protobuf:"varint,17,opt,name=hair_style,json=hairStyle,proto3" json:"hair_style,omitempty"`
	Beard        uint32  `protobuf:"varint,18,opt,name=beard,proto3" json:"beard,omitempty"`
	Title        uint32  `protobuf:"varint,19,opt,name=title,proto3" json:"title,omitempty"`
	Size         float32 `protobuf:"fixed32,20,opt,name=size,proto3" json:"size,omitempty"`
	WalkSpeed    float32 `protobuf:"fixed32,21,opt,name=walk_speed,json=walkSpeed,proto3" json:"walk_speed,omitempty"`
	RunSpeed     float32 `protobuf:"fixed32,22,opt,name=run_speed,json=runSpeed,proto3" json:"run_speed,omitempty"`
	BodyType     int32   `protobuf:"varint,23,opt,name=body_type,json=bodyType,proto3" json:"body_type,omitempty"`
	HpCurr       int32   `protobuf:"varint,24,opt,name=hp_curr,json=hpCurr,proto3" json:"hp_curr,omitempty"`
	GuildId      int32   `protobuf:"varint,25,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	Race         int32   `protobuf:"varint,26,opt,name=race,proto3" json:"race,omitempty"`
	Npc          uint32  `protobuf:"varint,27,opt,name=npc,proto3" json:"npc,omitempty"`
	ClassB       uint32  `protobuf:"varint,28,opt,name=class_b,json=classB,proto3" json:"class_b,omitempty"`
	Gender       uint32  `protobuf:"varint,29,opt,name=gender,proto3" json:"gender,omitempty"`
	Level        uint32  `protobuf:"varint,30,opt,name=level,proto3" json:"level,omitempty"`
	Invis        uint32  `protobuf:"varint,31,opt,name=invis,proto3" json:"invis,omitempty"`
	Sneaking     uint32  `protobuf:"varint,32,opt,name=sneaking,proto3" json:"sneaking,omitempty"`
	Pvp          uint32  `protobuf:"varint,33,opt,name=pvp,proto3" json:"pvp,omitempty"`
	AnimType     uint32  `protobuf:"varint,34,opt,name=anim_type,json=animType,proto3" json:"anim_type,omitempty"`
	Light        uint32  `protobuf:"varint,35,opt,name=light,proto3" json:"light,omitempty"`
	Anon         int32   `protobuf:"varint,36,opt,name=anon,proto3" json:"anon,omitempty"`
	Afk          int32   `protobuf:"varint,37,opt,name=afk,proto3" json:"afk,omitempty"`
	SummonedPc   int32   `protobuf:"varint,38,opt,name=summoned_pc,json=summonedPc,proto3" json:"summoned_pc,omitempty"`
	Ld           int32   `protobuf:"varint,39,opt,name=ld,proto3" json:"ld,omitempty"`
	Gm           int32   `protobuf:"varint,40,opt,name=gm,proto3" json:"gm,omitempty"`
	FlyMode      int32   `protobuf:"varint,41,opt,name=fly_mode,json=flyMode,proto3" json:"fly_mode,omitempty"`
	BodyTexture  int32   `protobuf:"varint,42,opt,name=body_texture,json=bodyTexture,proto3" json:"body_texture,omitempty"`
	Helm         int32   `protobuf:"varint,43,opt,name=helm,proto3" json:"helm,omitempty"`
	Face         uint32  `protobuf:"varint,44,opt,name=face,proto3" json:"face,omitempty"`
	GuildRank    int32   `protobuf:"varint,45,opt,name=guild_rank,json=guildRank,proto3" json:"guild_rank,omitempty"`
	Diety        int32   `protobuf:"varint,46,opt,name=diety,proto3" json:"diety,omitempty"`
	TempPet      int32   `protobuf:"varint,47,opt,name=temp_pet,json=tempPet,proto3" json:"temp_pet,omitempty"`
	Name         string  `protobuf:"bytes,48,opt,name=name,proto3" json:"name,omitempty"`
	Surname      string  `protobuf:"bytes,49,opt,name=surname,proto3" json:"surname,omitempty"`
}

func (x *Spawn) Reset() {
	*x = Spawn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eqstruct_spawns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spawn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spawn) ProtoMessage() {}

func (x *Spawn) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eqstruct_spawns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spawn.ProtoReflect.Descriptor instead.
func (*Spawn) Descriptor() ([]byte, []int) {
	return file_proto_eqstruct_spawns_proto_rawDescGZIP(), []int{0}
}

func (x *Spawn) GetSpawnId() uint32 {
	if x != nil {
		return x.SpawnId
	}
	return 0
}

func (x *Spawn) GetAcceleration() int32 {
	if x != nil {
		return x.Acceleration
	}
	return 0
}

func (x *Spawn) GetHeading() uint32 {
	if x != nil {
		return x.Heading
	}
	return 0
}

func (x *Spawn) GetHeadingDelta() int32 {
	if x != nil {
		return x.HeadingDelta
	}
	return 0
}

func (x *Spawn) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Spawn) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Spawn) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *Spawn) GetYDelta() uint32 {
	if x != nil {
		return x.YDelta
	}
	return 0
}

func (x *Spawn) GetXDelta() uint32 {
	if x != nil {
		return x.XDelta
	}
	return 0
}

func (x *Spawn) GetZDelta() uint32 {
	if x != nil {
		return x.ZDelta
	}
	return 0
}

func (x *Spawn) GetPetOwner() uint32 {
	if x != nil {
		return x.PetOwner
	}
	return 0
}

func (x *Spawn) GetAnimation() uint32 {
	if x != nil {
		return x.Animation
	}
	return 0
}

func (x *Spawn) GetHairColor() uint32 {
	if x != nil {
		return x.HairColor
	}
	return 0
}

func (x *Spawn) GetBeardColor() uint32 {
	if x != nil {
		return x.BeardColor
	}
	return 0
}

func (x *Spawn) GetEyeColor_1() uint32 {
	if x != nil {
		return x.EyeColor_1
	}
	return 0
}

func (x *Spawn) GetEyeColor_2() uint32 {
	if x != nil {
		return x.EyeColor_2
	}
	return 0
}

func (x *Spawn) GetHairStyle() uint32 {
	if x != nil {
		return x.HairStyle
	}
	return 0
}

func (x *Spawn) GetBeard() uint32 {
	if x != nil {
		return x.Beard
	}
	return 0
}

func (x *Spawn) GetTitle() uint32 {
	if x != nil {
		return x.Title
	}
	return 0
}

func (x *Spawn) GetSize() float32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Spawn) GetWalkSpeed() float32 {
	if x != nil {
		return x.WalkSpeed
	}
	return 0
}

func (x *Spawn) GetRunSpeed() float32 {
	if x != nil {
		return x.RunSpeed
	}
	return 0
}

func (x *Spawn) GetBodyType() int32 {
	if x != nil {
		return x.BodyType
	}
	return 0
}

func (x *Spawn) GetHpCurr() int32 {
	if x != nil {
		return x.HpCurr
	}
	return 0
}

func (x *Spawn) GetGuildId() int32 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

func (x *Spawn) GetRace() int32 {
	if x != nil {
		return x.Race
	}
	return 0
}

func (x *Spawn) GetNpc() uint32 {
	if x != nil {
		return x.Npc
	}
	return 0
}

func (x *Spawn) GetClassB() uint32 {
	if x != nil {
		return x.ClassB
	}
	return 0
}

func (x *Spawn) GetGender() uint32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *Spawn) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Spawn) GetInvis() uint32 {
	if x != nil {
		return x.Invis
	}
	return 0
}

func (x *Spawn) GetSneaking() uint32 {
	if x != nil {
		return x.Sneaking
	}
	return 0
}

func (x *Spawn) GetPvp() uint32 {
	if x != nil {
		return x.Pvp
	}
	return 0
}

func (x *Spawn) GetAnimType() uint32 {
	if x != nil {
		return x.AnimType
	}
	return 0
}

func (x *Spawn) GetLight() uint32 {
	if x != nil {
		return x.Light
	}
	return 0
}

func (x *Spawn) GetAnon() int32 {
	if x != nil {
		return x.Anon
	}
	return 0
}

func (x *Spawn) GetAfk() int32 {
	if x != nil {
		return x.Afk
	}
	return 0
}

func (x *Spawn) GetSummonedPc() int32 {
	if x != nil {
		return x.SummonedPc
	}
	return 0
}

func (x *Spawn) GetLd() int32 {
	if x != nil {
		return x.Ld
	}
	return 0
}

func (x *Spawn) GetGm() int32 {
	if x != nil {
		return x.Gm
	}
	return 0
}

func (x *Spawn) GetFlyMode() int32 {
	if x != nil {
		return x.FlyMode
	}
	return 0
}

func (x *Spawn) GetBodyTexture() int32 {
	if x != nil {
		return x.BodyTexture
	}
	return 0
}

func (x *Spawn) GetHelm() int32 {
	if x != nil {
		return x.Helm
	}
	return 0
}

func (x *Spawn) GetFace() uint32 {
	if x != nil {
		return x.Face
	}
	return 0
}

func (x *Spawn) GetGuildRank() int32 {
	if x != nil {
		return x.GuildRank
	}
	return 0
}

func (x *Spawn) GetDiety() int32 {
	if x != nil {
		return x.Diety
	}
	return 0
}

func (x *Spawn) GetTempPet() int32 {
	if x != nil {
		return x.TempPet
	}
	return 0
}

func (x *Spawn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Spawn) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

type Spawns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spawns []*Spawn `protobuf:"bytes,1,rep,name=spawns,proto3" json:"spawns,omitempty"`
}

func (x *Spawns) Reset() {
	*x = Spawns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eqstruct_spawns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spawns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spawns) ProtoMessage() {}

func (x *Spawns) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eqstruct_spawns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spawns.ProtoReflect.Descriptor instead.
func (*Spawns) Descriptor() ([]byte, []int) {
	return file_proto_eqstruct_spawns_proto_rawDescGZIP(), []int{1}
}

func (x *Spawns) GetSpawns() []*Spawn {
	if x != nil {
		return x.Spawns
	}
	return nil
}

type SpawnPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpawnId      uint32 `protobuf:"varint,1,opt,name=spawn_id,json=spawnId,proto3" json:"spawn_id,omitempty"`
	AnimType     int32  `protobuf:"varint,2,opt,name=anim_type,json=animType,proto3" json:"anim_type,omitempty"`
	Heading      uint32 `protobuf:"varint,3,opt,name=heading,proto3" json:"heading,omitempty"`
	HeadingDelta int32  `protobuf:"varint,4,opt,name=heading_delta,json=headingDelta,proto3" json:"heading_delta,omitempty"`
	X            int32  `protobuf:"varint,5,opt,name=x,proto3" json:"x,omitempty"`
	Y            int32  `protobuf:"varint,6,opt,name=y,proto3" json:"y,omitempty"`
	Z            int32  `protobuf:"varint,7,opt,name=z,proto3" json:"z,omitempty"`
	Delta        uint32 `protobuf:"varint,8,opt,name=delta,proto3" json:"delta,omitempty"`
}

func (x *SpawnPosition) Reset() {
	*x = SpawnPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eqstruct_spawns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpawnPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpawnPosition) ProtoMessage() {}

func (x *SpawnPosition) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eqstruct_spawns_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpawnPosition.ProtoReflect.Descriptor instead.
func (*SpawnPosition) Descriptor() ([]byte, []int) {
	return file_proto_eqstruct_spawns_proto_rawDescGZIP(), []int{2}
}

func (x *SpawnPosition) GetSpawnId() uint32 {
	if x != nil {
		return x.SpawnId
	}
	return 0
}

func (x *SpawnPosition) GetAnimType() int32 {
	if x != nil {
		return x.AnimType
	}
	return 0
}

func (x *SpawnPosition) GetHeading() uint32 {
	if x != nil {
		return x.Heading
	}
	return 0
}

func (x *SpawnPosition) GetHeadingDelta() int32 {
	if x != nil {
		return x.HeadingDelta
	}
	return 0
}

func (x *SpawnPosition) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *SpawnPosition) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *SpawnPosition) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *SpawnPosition) GetDelta() uint32 {
	if x != nil {
		return x.Delta
	}
	return 0
}

type SpawnPositions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Positions []*SpawnPosition `protobuf:"bytes,1,rep,name=positions,proto3" json:"positions,omitempty"`
}

func (x *SpawnPositions) Reset() {
	*x = SpawnPositions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eqstruct_spawns_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpawnPositions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpawnPositions) ProtoMessage() {}

func (x *SpawnPositions) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eqstruct_spawns_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpawnPositions.ProtoReflect.Descriptor instead.
func (*SpawnPositions) Descriptor() ([]byte, []int) {
	return file_proto_eqstruct_spawns_proto_rawDescGZIP(), []int{3}
}

func (x *SpawnPositions) GetPositions() []*SpawnPosition {
	if x != nil {
		return x.Positions
	}
	return nil
}

type SpawnAppearance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpawnId   uint32 `protobuf:"varint,1,opt,name=spawn_id,json=spawnId,proto3" json:"spawn_id,omitempty"`
	Type      uint32 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Parameter uint32 `protobuf:"varint,3,opt,name=parameter,proto3" json:"parameter,omitempty"`
}

func (x *SpawnAppearance) Reset() {
	*x = SpawnAppearance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eqstruct_spawns_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpawnAppearance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpawnAppearance) ProtoMessage() {}

func (x *SpawnAppearance) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eqstruct_spawns_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpawnAppearance.ProtoReflect.Descriptor instead.
func (*SpawnAppearance) Descriptor() ([]byte, []int) {
	return file_proto_eqstruct_spawns_proto_rawDescGZIP(), []int{4}
}

func (x *SpawnAppearance) GetSpawnId() uint32 {
	if x != nil {
		return x.SpawnId
	}
	return 0
}

func (x *SpawnAppearance) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *SpawnAppearance) GetParameter() uint32 {
	if x != nil {
		return x.Parameter
	}
	return 0
}

var File_proto_eqstruct_spawns_proto protoreflect.FileDescriptor

var file_proto_eqstruct_spawns_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x71, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2f, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65,
	0x71, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0xd0, 0x09, 0x0a, 0x05, 0x53, 0x70, 0x61, 0x77,
	0x6e, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x65,
	0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12,
	0x0c, 0x0a, 0x01, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x7a,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x7a, 0x12, 0x17, 0x0a, 0x07, 0x79, 0x5f, 0x64,
	0x65, 0x6c, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x79, 0x44, 0x65, 0x6c,
	0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x78, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x78, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x7a,
	0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x7a, 0x44,
	0x65, 0x6c, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x74, 0x5f, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x68, 0x61, 0x69, 0x72, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x68, 0x61, 0x69, 0x72, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x62, 0x65, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x62, 0x65, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x1e, 0x0a, 0x0b, 0x65, 0x79, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x5f, 0x31, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x79, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x31, 0x12,
	0x1e, 0x0a, 0x0b, 0x65, 0x79, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x5f, 0x32, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x79, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x32, 0x12,
	0x1d, 0x0a, 0x0a, 0x68, 0x61, 0x69, 0x72, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x68, 0x61, 0x69, 0x72, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x65, 0x61, 0x72, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x62,
	0x65, 0x61, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x77, 0x61, 0x6c, 0x6b, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x77, 0x61, 0x6c, 0x6b, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x75, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x72, 0x75, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f,
	0x64, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62,
	0x6f, 0x64, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x70, 0x5f, 0x63, 0x75,
	0x72, 0x72, 0x18, 0x18, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x70, 0x43, 0x75, 0x72, 0x72,
	0x12, 0x19, 0x0a, 0x08, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x61, 0x63, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x61, 0x63, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x70, 0x63, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6e, 0x70,
	0x63, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x62, 0x18, 0x1c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x76, 0x69,
	0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x76, 0x69, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x6e, 0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x73, 0x6e, 0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x76,
	0x70, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x76, 0x70, 0x12, 0x1b, 0x0a, 0x09,
	0x61, 0x6e, 0x69, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x22, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x61, 0x6e, 0x69, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x23, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x6e, 0x6f, 0x6e, 0x18, 0x24, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x61,
	0x6e, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x66, 0x6b, 0x18, 0x25, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x61, 0x66, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65,
	0x64, 0x5f, 0x70, 0x63, 0x18, 0x26, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x75, 0x6d, 0x6d,
	0x6f, 0x6e, 0x65, 0x64, 0x50, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x6c, 0x64, 0x18, 0x27, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x6c, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x67, 0x6d, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x67, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x6c, 0x79, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x29, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x6c, 0x79, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x62, 0x6f, 0x64, 0x79, 0x54, 0x65, 0x78,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x6c, 0x6d, 0x18, 0x2b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x68, 0x65, 0x6c, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x61, 0x63, 0x65,
	0x18, 0x2c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x69, 0x65, 0x74, 0x79, 0x18, 0x2e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x69, 0x65, 0x74,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x70, 0x65, 0x74, 0x18, 0x2f, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x65, 0x6d, 0x70, 0x50, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x30, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x31, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x06, 0x53, 0x70,
	0x61, 0x77, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x71, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x53, 0x70, 0x61, 0x77, 0x6e, 0x52, 0x06, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x73, 0x22, 0xc6, 0x01,
	0x0a, 0x0d, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6e,
	0x69, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61,
	0x6e, 0x69, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x6c,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x7a,
	0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x22, 0x47, 0x0a, 0x0e, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x71,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x5e, 0x0a, 0x0f, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x42,
	0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f,
	0x6d, 0x6f, 0x72, 0x65, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x7a, 0x2f, 0x67, 0x68, 0x6f, 0x65,
	0x71, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x71, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_eqstruct_spawns_proto_rawDescOnce sync.Once
	file_proto_eqstruct_spawns_proto_rawDescData = file_proto_eqstruct_spawns_proto_rawDesc
)

func file_proto_eqstruct_spawns_proto_rawDescGZIP() []byte {
	file_proto_eqstruct_spawns_proto_rawDescOnce.Do(func() {
		file_proto_eqstruct_spawns_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_eqstruct_spawns_proto_rawDescData)
	})
	return file_proto_eqstruct_spawns_proto_rawDescData
}

var file_proto_eqstruct_spawns_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_eqstruct_spawns_proto_goTypes = []interface{}{
	(*Spawn)(nil),           // 0: eqstruct.Spawn
	(*Spawns)(nil),          // 1: eqstruct.Spawns
	(*SpawnPosition)(nil),   // 2: eqstruct.SpawnPosition
	(*SpawnPositions)(nil),  // 3: eqstruct.SpawnPositions
	(*SpawnAppearance)(nil), // 4: eqstruct.SpawnAppearance
}
var file_proto_eqstruct_spawns_proto_depIdxs = []int32{
	0, // 0: eqstruct.Spawns.spawns:type_name -> eqstruct.Spawn
	2, // 1: eqstruct.SpawnPositions.positions:type_name -> eqstruct.SpawnPosition
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_eqstruct_spawns_proto_init() }
func file_proto_eqstruct_spawns_proto_init() {
	if File_proto_eqstruct_spawns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_eqstruct_spawns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spawn); i {
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
		file_proto_eqstruct_spawns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spawns); i {
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
		file_proto_eqstruct_spawns_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpawnPosition); i {
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
		file_proto_eqstruct_spawns_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpawnPositions); i {
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
		file_proto_eqstruct_spawns_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpawnAppearance); i {
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
			RawDescriptor: file_proto_eqstruct_spawns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_eqstruct_spawns_proto_goTypes,
		DependencyIndexes: file_proto_eqstruct_spawns_proto_depIdxs,
		MessageInfos:      file_proto_eqstruct_spawns_proto_msgTypes,
	}.Build()
	File_proto_eqstruct_spawns_proto = out.File
	file_proto_eqstruct_spawns_proto_rawDesc = nil
	file_proto_eqstruct_spawns_proto_goTypes = nil
	file_proto_eqstruct_spawns_proto_depIdxs = nil
}