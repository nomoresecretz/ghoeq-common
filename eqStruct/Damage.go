package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type Damage struct {
	Target      uint16
	Source      uint16
	Type        uint16
	SpellID     uint16
	Damage      int32
	Force       float32
	Sequence    float32
	PushupAngle float32

	bPointer int
}

func (p *Damage) EQType() EQType { return EQT_Damage }
func (p *Damage) bp() *int       { return &p.bPointer }

func (p *Damage) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.Target, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Source, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Type, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.SpellID, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Damage, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Force, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Sequence, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.PushupAngle, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *Damage) Proto() *eqstruct.Damage {
	return &eqstruct.Damage{
		Target:      uint32(p.Target),
		Source:      uint32(p.Source),
		Type:        uint32(p.Type),
		SpellId:     uint32(p.SpellID),
		Damage:      p.Damage,
		Force:       p.Force,
		Sequence:    p.Sequence,
		PushupAngle: p.PushupAngle,
	}
}

func (p *Damage) ProtoMess() proto.Message {
	return p.Proto()
}
