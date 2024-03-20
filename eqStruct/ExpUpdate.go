package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type ExpUpdate struct {
	Exp uint32 // 000

	bPointer int
}

func (p *ExpUpdate) EQType() EQType { return EQT_ExpUpdate }
func (p *ExpUpdate) bp() *int       { return &p.bPointer }

func (p *ExpUpdate) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.Exp, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *ExpUpdate) Proto() *eqstruct.ExpUpdate {
	return &eqstruct.ExpUpdate{
		Exp: p.Exp,
	}
}

func (p *ExpUpdate) ProtoMess() proto.Message {
	return p.Proto()
}
