package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type RaidAddMember struct {
	RaidGeneral     RaidGeneral
	Class           uint8
	Level           uint8
	GroupLeaderFlag uint8

	bPointer int
}

func (p *RaidAddMember) EQType() EQType { return EQT_RaidAddMember }
func (p *RaidAddMember) bp() *int       { return &p.bPointer }

func (p *RaidAddMember) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	p.RaidGeneral = RaidGeneral{}
	bp, err := p.RaidGeneral.Unmarshal(b)
	if err != nil {
		return 0, err
	}
	p.bPointer += bp

	p.bPointer = 136
	if err := EQReadLittleEndian(b, p, &p.Class, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Level, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.GroupLeaderFlag, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *RaidAddMember) Proto() *eqstruct.RaidAddMember {
	return &eqstruct.RaidAddMember{
		RaidGeneral:     p.RaidGeneral.Proto(),
		Class:           uint32(p.Class),
		Level:           uint32(p.Level),
		GroupLeaderFlag: uint32(p.GroupLeaderFlag),
	}
}

func (p *RaidAddMember) ProtoMess() proto.Message {
	return p.Proto()
}
