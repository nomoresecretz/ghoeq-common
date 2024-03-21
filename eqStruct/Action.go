package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type Action struct {
	Target        uint16 // 00
	Source        uint16 // 02
	Level         uint16 // 04
	TargetLevel   uint16 // 06
	InstrumentMod int32  // 08
	Force         float32
	Sequence      float32
	PushupAngle   float32
	Type          uint8
	TapAmount     int16
	Spell         uint16
	BuffUnknown   uint8

	bPointer int
}

func (p *Action) EQType() EQType   { return EQT_Action }
func (p *Action) bp() *int         { return &p.bPointer }
func (p *Action) SetPointer(i int) { p.bPointer = i }

func (p *Action) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.Target, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Source, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Level, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.TargetLevel, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.InstrumentMod, 0); err != nil {
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

	if err := EQRead(b, p, &p.Type, 0); err != nil {
		return 0, err
	}

	p.bPointer = 28
	if err := EQRead(b, p, &p.TapAmount, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Spell, 0); err != nil {
		return 0, err
	}

	p.bPointer = 33
	if err := EQRead(b, p, &p.BuffUnknown, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *Action) Proto() *eqstruct.Action {
	return &eqstruct.Action{
		Target:        uint32(p.Target),
		Source:        uint32(p.Source),
		Level:         uint32(p.Level),
		InstrumentMod: p.InstrumentMod,
		Force:         p.Force,
		Sequence:      p.Sequence,
		PushupAngle:   p.PushupAngle,
		Type:          uint32(p.Type),
		TapAmount:     int32(p.TapAmount),
		Spell:         uint32(p.Spell),
		BuffUnknown:   uint32(p.BuffUnknown),
	}
}

func (p *Action) ProtoMess() proto.Message {
	return p.Proto()
}
