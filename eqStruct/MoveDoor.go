package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type MoveDoor struct {
	DoorID int8
	Action int8

	bPointer int
}

func (p *MoveDoor) EQType() EQType { return EQT_MoveDoor }
func (p *MoveDoor) bp() *int       { return &p.bPointer }

func (p *MoveDoor) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.DoorID, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Action, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *MoveDoor) Proto() *eqstruct.MoveDoor {
	return &eqstruct.MoveDoor{}
}

func (p *MoveDoor) ProtoMess() proto.Message {
	return p.Proto()
}
