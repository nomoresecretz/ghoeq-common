package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type Target struct {
	TargetID uint16 // 000

	bPointer int
}

func (p *Target) EQType() EQType { return EQT_Target }
func (p *Target) bp() *int       { return &p.bPointer }

func (p *Target) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQReadLittleEndian(b, p, &p.TargetID, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *Target) Proto() *eqstruct.Target {
	return &eqstruct.Target{
		TargetId: uint32(p.TargetID),
	}
}

func (p *Target) ProtoMess() proto.Message {
	return p.Proto()
}
