package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type ZoneChangeReq struct {
	ZoneID  uint32
	Y       float32
	X       float32
	Z       float32
	Heading float32
	Type    uint32

	bPointer int
}

func (p *ZoneChangeReq) EQType() EQType { return EQT_ZoneChangeReq }
func (p *ZoneChangeReq) bp() *int       { return &p.bPointer }

func (p *ZoneChangeReq) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.ZoneID, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Y, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.X, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Z, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Heading, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Type, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *ZoneChangeReq) Proto() *eqstruct.ZoneChangeReq {
	return &eqstruct.ZoneChangeReq{
		ZoneId:  p.ZoneID,
		X:       p.X,
		Y:       p.Y,
		Z:       p.Z,
		Heading: p.Heading,
		Type:    p.Type,
	}
}

func (p *ZoneChangeReq) ProtoMess() proto.Message {
	return p.Proto()
}
