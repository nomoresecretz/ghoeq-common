package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type LFG struct {
	Name  string // MAX64
	Value int32

	bPointer int
}

func (p *LFG) EQType() EQType { return EQT_LFG }
func (p *LFG) bp() *int       { return &p.bPointer }

func (p *LFG) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQReadLittleEndian(b, p, &p.Name, 64); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Value, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *LFG) Proto() *eqstruct.LFG {
	return &eqstruct.LFG{
		Name:  p.Name,
		Value: uint32(p.Value),
	}
}

func (p *LFG) ProtoMess() proto.Message {
	return p.Proto()
}

type LFGAppearance struct {
	EntityID int16
	Unknown  int16
	Value    int32

	bPointer int
}

func (p *LFGAppearance) EQType() EQType { return EQT_LFGAppearance }
func (p *LFGAppearance) bp() *int       { return &p.bPointer }

func (p *LFGAppearance) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQReadLittleEndian(b, p, &p.EntityID, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Unknown, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Value, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *LFGAppearance) Proto() *eqstruct.LFGAppearance {
	return &eqstruct.LFGAppearance{
		EntityId: int32(p.EntityID),
		Value:    p.Value,
		Unknown:  int32(p.Unknown),
	}
}

func (p *LFGAppearance) ProtoMess() proto.Message {
	return p.Proto()
}
