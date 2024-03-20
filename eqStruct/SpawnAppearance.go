package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type SpawnAppearance struct {
	SpawnID   uint16
	Type      uint16
	Parameter uint32

	bPointer int
}

func (p *SpawnAppearance) EQType() EQType { return EQT_SpawnAppearance }
func (p *SpawnAppearance) bp() *int       { return &p.bPointer }

func (p *SpawnAppearance) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.SpawnID, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Type, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Parameter, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *SpawnAppearance) Proto() *eqstruct.SpawnAppearance {
	return &eqstruct.SpawnAppearance{
		SpawnId:   uint32(p.SpawnID),
		Type:      uint32(p.Type),
		Parameter: p.Parameter,
	}
}

func (p *SpawnAppearance) ProtoMess() proto.Message {
	return p.Proto()
}

//go:generate enumer -type=EQAppearanceType

type EQAppearanceType uint16

const (
	AT_Die         EQAppearanceType = 0
	AT_WhoLevel    EQAppearanceType = 1
	AT_MaxHP       EQAppearanceType = 2
	AT_Invis       EQAppearanceType = 3
	AT_PVP         EQAppearanceType = 4
	AT_Light       EQAppearanceType = 5
	AT_Anim        EQAppearanceType = 14
	AT_Sneak       EQAppearanceType = 15
	AT_SpawnID     EQAppearanceType = 16
	AT_HP          EQAppearanceType = 17
	AT_LinkDead    EQAppearanceType = 18
	AT_Levitate    EQAppearanceType = 19
	AT_GM          EQAppearanceType = 20
	AT_Anon        EQAppearanceType = 21
	AT_GuildID     EQAppearanceType = 22
	AT_GuildRank   EQAppearanceType = 23
	AT_AFK         EQAppearanceType = 24
	AT_Pet         EQAppearanceType = 25
	AT_SummonedPC  EQAppearanceType = 27
	AT_Split       EQAppearanceType = 28
	AT_Size        EQAppearanceType = 29
	AT_NPC         EQAppearanceType = 30
	AT_NPCName     EQAppearanceType = 31
	AT_DamageState EQAppearanceType = 44
	AT_Trader      EQAppearanceType = 300
)
