package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type ManaUpdate struct {
	SpawnID  uint16 // 000
	ManaCurr uint16

	bPointer int
}

func (p *ManaUpdate) EQType() EQType { return EQT_ManaUpdate }
func (p *ManaUpdate) bp() *int       { return &p.bPointer }

func (p *ManaUpdate) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.SpawnID, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.ManaCurr, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *ManaUpdate) Proto() *eqstruct.ManaUpdate {
	return &eqstruct.ManaUpdate{
		SpawnId:  uint32(p.SpawnID),
		ManaCurr: uint32(p.ManaCurr),
	}
}

func (p *ManaUpdate) ProtoMess() proto.Message {
	return p.Proto()
}
