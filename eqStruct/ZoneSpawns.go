package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type ZoneSpawn struct {
	Unused01     uint32 // 000
	Accel        int8   // 004
	Heading      uint8  // 005
	HeadingDelta int8   // 006
	Y            int16
	X            int16
	Z            int16
	YDelta       uint32 // velocities
	ZDelta       uint32
	XDelta       uint32
	Unused02     int8
	PetOwnerID   uint16
	Animation    uint8
	Haircolor    uint8
	Beardcolor   uint8
	Eyecolor1    uint8
	Eyecolor2    uint8
	Hairstyle    uint8
	Beard        uint8
	Title        uint8
	Size         float32
	Walkspeed    float32
	Runspeed     float32
	EquipColors  []byte // skip for now
	SpawnID      uint16
	BodyType     int16
	HPCur        int16
	GuildID      int16
	Race         uint16
	NPC          uint8
	ClassB       uint8
	Gender       uint8
	Level        uint8
	Invis        uint8
	Sneaking     uint8
	PVP          uint8
	AnimType     uint8
	Light        uint8
	ANON         int8
	AFK          int8
	SummonedPC   int8
	LD           int8
	GM           int8
	FlyMode      int8
	BodyTexture  int8
	Helm         int8
	Face         uint8
	Equipment    []uint16
	GuildRank    int16
	Deity        int16
	TemporaryPet int8
	Name         string
	Surname      string
	Void2        uint8

	bPointer int
}

func (p *ZoneSpawn) EQType() EQType { return EQT_ZoneSpawn }
func (p *ZoneSpawn) bp() *int       { return &p.bPointer }

func (p *ZoneSpawn) Unmarshal(b []byte) (int, error) {
	p.bPointer = 4
	if err := EQReadLittleEndian(b, p, &p.Accel, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Heading, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.HeadingDelta, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Y, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.X, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Z, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.YDelta, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.ZDelta, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.XDelta, 0); err != nil {
		return 0, err
	}

	p.bPointer = 18
	if err := EQReadLittleEndian(b, p, &p.PetOwnerID, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Animation, 0); err != nil {
		return 0, err
	}

	p.bPointer = 27
	if err := EQReadLittleEndian(b, p, &p.Title, 0); err != nil {
		return 0, err
	}

	p.bPointer = 32
	if err := EQReadLittleEndian(b, p, &p.Walkspeed, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Runspeed, 0); err != nil {
		return 0, err
	}

	p.bPointer = 76
	if err := EQReadLittleEndian(b, p, &p.SpawnID, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.BodyType, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.HPCur, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.GuildID, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Race, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.NPC, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.ClassB, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Gender, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Level, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Invis, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Sneaking, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.PVP, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.AnimType, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Light, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.ANON, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.AFK, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.SummonedPC, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.LD, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.GM, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.FlyMode, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.BodyTexture, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Helm, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Face, 0); err != nil {
		return 0, err
	}

	p.bPointer = 122
	if err := EQReadLittleEndian(b, p, &p.GuildRank, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Deity, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.TemporaryPet, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Name, 64); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Surname, 32); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Void2, 32); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

type ZoneSpawns struct {
	Spawns []*ZoneSpawn // 004

	bPointer int
}

func (p *ZoneSpawns) EQType() EQType { return EQT_ZoneSpawns }
func (p *ZoneSpawns) bp() *int       { return &p.bPointer }

func (p *ZoneSpawns) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	p.Spawns = make([]*ZoneSpawn, 0)
	for len(b)-p.bPointer >= 224 {
		spawn := &ZoneSpawn{}
		bp, err := spawn.Unmarshal(b[p.bPointer:])
		if err != nil {
			return 0, err
		}
		p.bPointer += bp
		p.Spawns = append(p.Spawns, spawn)
	}

	return p.bPointer, nil
}

func (p *ZoneSpawn) Proto() *eqstruct.Spawn {
	return &eqstruct.Spawn{
		SpawnId:      uint32(p.SpawnID),
		Acceleration: int32(p.Accel),
		Heading:      uint32(p.Heading),
		HeadingDelta: int32(p.HeadingDelta),
		X:            int32(p.X),
		Y:            int32(p.Y),
		Z:            int32(p.Z),
		YDelta:       p.YDelta,
		XDelta:       p.XDelta,
		ZDelta:       p.ZDelta,
		PetOwner:     uint32(p.PetOwnerID),
		Animation:    uint32(p.Animation),
		HairColor:    uint32(p.Haircolor),
		BeardColor:   uint32(p.Beardcolor),
		EyeColor_1:   uint32(p.Eyecolor1),
		EyeColor_2:   uint32(p.Eyecolor2),
		HairStyle:    uint32(p.Hairstyle),
		Beard:        uint32(p.Beard),
		Title:        uint32(p.Title),
		Size:         p.Size,
		WalkSpeed:    p.Walkspeed,
		RunSpeed:     p.Runspeed,
		BodyType:     int32(p.BodyType),
		HpCurr:       int32(p.HPCur),
		GuildId:      int32(p.GuildID),
		Race:         int32(p.Race),
		Npc:          uint32(p.NPC),
		ClassB:       uint32(p.ClassB),
		Gender:       uint32(p.Gender),
		Level:        uint32(p.Level),
		Invis:        uint32(p.Invis),
		Sneaking:     uint32(p.Sneaking),
		Pvp:          uint32(p.PVP),
		AnimType:     uint32(p.AnimType),
		Light:        uint32(p.Light),
		Anon:         int32(p.ANON),
		Afk:          int32(p.AFK),
		SummonedPc:   int32(p.SummonedPC),
		Ld:           int32(p.LD),
		Gm:           int32(p.GM),
		FlyMode:      int32(p.FlyMode),
		BodyTexture:  int32(p.BodyTexture),
		Helm:         int32(p.Helm),
		Face:         uint32(p.Face),
		GuildRank:    int32(p.GuildRank),
		Diety:        int32(p.Deity),
		TempPet:      int32(p.TemporaryPet),
		Name:         p.Name,
		Surname:      p.Surname,
	}
}

func (p *ZoneSpawns) Proto() *eqstruct.Spawns {
	arr := make([]*eqstruct.Spawn, len(p.Spawns))
	for i := range arr {
		arr[i] = p.Spawns[i].Proto()
	}

	return &eqstruct.Spawns{
		Spawns: arr,
	}
}

func (p *ZoneSpawn) ProtoMess() proto.Message {
	return p.Proto()
}

func (p *ZoneSpawns) ProtoMess() proto.Message {
	return p.Proto()
}
