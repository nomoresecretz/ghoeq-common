package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type WearChange struct {
	SpawnID    uint16
	WearSlotID uint8
	Material   uint16
	Color      *TintStruct

	bPointer int
}

func (p *WearChange) EQType() EQType { return EQT_WearChange }
func (p *WearChange) bp() *int       { return &p.bPointer }

func (p *WearChange) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQReadLittleEndian(b, p, &p.SpawnID, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.WearSlotID, 0); err != nil {
		return 0, err
	}

	p.bPointer = 4
	if err := EQReadLittleEndian(b, p, &p.Material, 0); err != nil {
		return 0, err
	}

	p.Color = &TintStruct{}
	bp, err := p.Color.Unmarshal(b[p.bPointer:])
	if err != nil {
		return 0, err
	}
	p.bPointer += bp

	return p.bPointer, nil
}

func (p *WearChange) Proto() *eqstruct.WearChange {
	return &eqstruct.WearChange{
		SpawnId:    uint32(p.SpawnID),
		WearSlotId: uint32(p.WearSlotID),
		Material:   uint32(p.Material),
		Color:      p.Color.Proto(),
	}
}

func (p *WearChange) ProtoMess() proto.Message {
	return p.Proto()
}
