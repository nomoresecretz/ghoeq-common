package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type Death struct {
	SpawnID     uint16 // 000
	KillerID    uint16
	CorpseID    uint16
	SpawnLevel  uint8
	SpellID     uint16
	AttackSkill uint8
	Damage      int32
	IsPC        uint8

	bPointer int
}

func (p *Death) EQType() EQType   { return EQT_Death }
func (p *Death) bp() *int         { return &p.bPointer }
func (p *Death) SetPointer(i int) { p.bPointer = i }

func (p *Death) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.SpawnID, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.KillerID, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.CorpseID, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.SpawnLevel, 0); err != nil {
		return 0, err
	}

	p.bPointer = 8
	if err := EQRead(b, p, &p.SpellID, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.AttackSkill, 0); err != nil {
		return 0, err
	}

	p.bPointer = 12
	if err := EQRead(b, p, &p.Damage, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.IsPC, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *Death) Proto() *eqstruct.Death {
	return &eqstruct.Death{
		SpawnId:     uint32(p.SpawnID),
		KillerId:    uint32(p.KillerID),
		CorpseId:    uint32(p.CorpseID),
		SpawnLevel:  uint32(p.SpawnLevel),
		SpellId:     uint32(p.SpawnID),
		AttackSkill: uint32(p.AttackSkill),
		Damage:      uint32(p.Damage),
		IsPc:        uint32(p.IsPC),
	}
}

func (p *Death) ProtoMess() proto.Message {
	return p.Proto()
}
