package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type HPUpdate struct {
	SpawnID uint32 // 000
	HPCur   int32
	HPMax   int32

	bPointer int
}

func (p *HPUpdate) EQType() EQType { return EQT_HPUpdate }
func (p *HPUpdate) bp() *int       { return &p.bPointer }

func (p *HPUpdate) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQReadLittleEndian(b, p, &p.SpawnID, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.HPCur, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.HPMax, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *HPUpdate) Proto() *eqstruct.HPUpdate {
	return &eqstruct.HPUpdate{
		SpawnId: uint32(p.SpawnID),
		HpCur:   p.HPCur,
		HpMax:   p.HPMax,
	}
}

func (p *HPUpdate) ProtoMess() proto.Message {
	return p.Proto()
}
