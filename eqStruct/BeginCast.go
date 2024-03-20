package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type BeginCast struct {
	CasterID uint16
	SpellID  uint16
	CastTime uint16

	bPointer int
}

func (p *BeginCast) EQType() EQType { return EQT_BeginCast }
func (p *BeginCast) bp() *int       { return &p.bPointer }

func (p *BeginCast) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.CasterID, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.SpellID, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.CastTime, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *BeginCast) Proto() *eqstruct.BeginCast {
	return &eqstruct.BeginCast{
		CasterId: uint32(p.CasterID),
		SpellId:  uint32(p.SpellID),
		CastTime: uint32(p.CastTime),
	}
}

func (p *BeginCast) ProtoMess() proto.Message {
	return p.Proto()
}
