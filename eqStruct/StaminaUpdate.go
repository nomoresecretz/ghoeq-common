package eqStruct

import (
	pb "github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type StaminaUpdate struct {
	Food    int16
	Water   int16
	Fatigue int8

	bPointer int
}

func (p *StaminaUpdate) EQType() EQType { return EQT_StaminaUpdate }
func (p *StaminaUpdate) bp() *int       { return &p.bPointer }

func (p *StaminaUpdate) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0
	if err := EQReadLittleEndian(b, p, &p.Food, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Water, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Fatigue, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *StaminaUpdate) Proto() *pb.StaminaUpdate {
	return &pb.StaminaUpdate{
		Food:    int32(p.Food),
		Water:   int32(p.Water),
		Fatigue: int32(p.Fatigue),
	}
}

func (p *StaminaUpdate) ProtoMess() proto.Message {
	return p.Proto()
}
