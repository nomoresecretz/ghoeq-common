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

	if err := EQReadLittleEndian(b, p, &p.SpawnID, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Type, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Parameter, 0); err != nil {
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

const (
	AT_Die         = 0
	AT_WhoLevel    = 1
	AT_MaxHP       = 2
	AT_Invis       = 3
	AT_PVP         = 4
	AT_Light       = 5
	AT_Anim        = 14
	AT_Sneak       = 15
	AT_SpawnID     = 16
	AT_HP          = 17
	AT_LinkDead    = 18
	AT_Levitate    = 19
	AT_GM          = 20
	AT_Anon        = 21
	AT_GuildID     = 22
	AT_GuildRank   = 23
	AT_AFK         = 24
	AT_Pet         = 25
	AT_SummonedPC  = 27
	AT_Split       = 28
	AT_Size        = 29
	AT_NPC         = 30
	AT_NPCName     = 31
	AT_DamageState = 44
	AT_Trader      = 300
)
