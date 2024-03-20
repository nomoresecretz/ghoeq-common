package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type WorldApprove struct {
	Response uint8
	Name     string // MAX64

	bPointer int
}

func (p *WorldApprove) EQType() EQType { return EQT_WorldApprove }
func (p *WorldApprove) bp() *int       { return &p.bPointer }

func (p *WorldApprove) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.Response, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Name, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *WorldApprove) Proto() *eqstruct.WorldApprove {
	return &eqstruct.WorldApprove{
		Response: uint32(p.Response),
		Name:     p.Name,
	}
}

func (p *WorldApprove) ProtoMess() proto.Message {
	return p.Proto()
}
