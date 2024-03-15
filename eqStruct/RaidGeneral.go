package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type RaidGeneral struct {
	Action     uint32
	PlayerName string // MAX64
	LeaderName string // MAX64
	Parameter  uint32

	bPointer int
}

func (p *RaidGeneral) EQType() EQType { return EQT_RaidGeneral }
func (p *RaidGeneral) bp() *int       { return &p.bPointer }

func (p *RaidGeneral) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQReadLittleEndian(b, p, &p.Action, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.PlayerName, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.LeaderName, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Parameter, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *RaidGeneral) Proto() *eqstruct.RaidGeneral {
	return &eqstruct.RaidGeneral{}
}

func (p *RaidGeneral) ProtoMess() proto.Message {
	return p.Proto()
}
