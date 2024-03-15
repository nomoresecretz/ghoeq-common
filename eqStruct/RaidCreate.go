package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type RaidCreate struct {
	Action     uint32
	LeaderName string // MAX64
	LeaderID   uint32

	bPointer int
}

func (p *RaidCreate) EQType() EQType { return EQT_RaidCreate }
func (p *RaidCreate) bp() *int       { return &p.bPointer }

func (p *RaidCreate) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQReadLittleEndian(b, p, &p.Action, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.LeaderName, 64); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.LeaderID, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *RaidCreate) Proto() *eqstruct.RaidCreate {
	return &eqstruct.RaidCreate{
		Action:     p.Action,
		LeaderName: p.LeaderName,
		LeaderId:   p.LeaderID,
	}
}

func (p *RaidCreate) ProtoMess() proto.Message {
	return p.Proto()
}
