package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type EnterWorld struct {
	Name string // 000

	bPointer int
}

func (p *EnterWorld) EQType() EQType   { return EQT_EnterWorld }
func (p *EnterWorld) bp() *int         { return &p.bPointer }
func (p *EnterWorld) SetPointer(i int) { p.bPointer = i }

func (p *EnterWorld) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.Name, 64); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *EnterWorld) Proto() *eqstruct.EnterWorld {
	return &eqstruct.EnterWorld{
		Name: p.Name,
	}
}

func (p *EnterWorld) ProtoMess() proto.Message {
	return p.Proto()
}
