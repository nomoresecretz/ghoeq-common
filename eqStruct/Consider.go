package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type Consider struct {
	PlayerID uint16
	TargetID uint16
	Faction  uint32
	Level    uint32
	HPCur    int32
	HPMax    int32
	PVPCon   uint8

	bPointer int
}

func (p *Consider) EQType() EQType { return EQT_Consider }
func (p *Consider) bp() *int       { return &p.bPointer }

func (p *Consider) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.PlayerID, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.TargetID, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Faction, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Level, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.HPCur, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.HPMax, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.PVPCon, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *Consider) Proto() *eqstruct.Consider {
	return &eqstruct.Consider{
		PlayerId: uint32(p.PlayerID),
		TargetId: uint32(p.TargetID),
		Faction:  p.Faction,
		Level:    p.Level,
		CurrHp:   p.HPCur,
		MaxHp:    p.HPMax,
		Pvpcon:   uint32(p.PVPCon),
	}
}

func (p *Consider) ProtoMess() proto.Message {
	return p.Proto()
}
