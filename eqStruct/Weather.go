package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type Weather struct {
	Type      uint32
	Intensity uint32

	bPointer int
}

func (p *Weather) EQType() EQType { return EQT_Weather }
func (p *Weather) bp() *int       { return &p.bPointer }

func (p *Weather) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.Type, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Intensity, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *Weather) Proto() *eqstruct.Weather {
	return &eqstruct.Weather{
		Type:      p.Type,
		Intensity: p.Intensity,
	}
}

func (p *Weather) ProtoMess() proto.Message {
	return p.Proto()
}
