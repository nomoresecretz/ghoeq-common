package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type ZoneChange struct {
	CharName   string // MAX64
	ZoneID     uint16
	ZoneReason uint16
	Success    int8
	Error      [3]uint8

	bPointer int
}

func (p *ZoneChange) EQType() EQType { return EQT_ZoneChange }
func (p *ZoneChange) bp() *int       { return &p.bPointer }

func (p *ZoneChange) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.CharName, 64); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.ZoneID, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.ZoneReason, 0); err != nil {
		return 0, err
	}

	p.bPointer = 72
	if err := EQRead(b, p, &p.Success, 0); err != nil {
		return 0, err
	}

	p.Error = [3]uint8{}
	for i := range p.Error {
		if err := EQRead(b, p, &p.Error[i], 0); err != nil {
			return 0, err
		}
	}

	return p.bPointer, nil
}

func (p *ZoneChange) Proto() *eqstruct.ZoneChange {
	err := make([]uint32, len(p.Error))
	for i, v := range p.Error {
		err[i] = uint32(v)
	}

	return &eqstruct.ZoneChange{
		CharName:   p.CharName,
		ZoneId:     uint32(p.ZoneID),
		ZoneReason: uint32(p.ZoneReason),
		Success:    uint32(p.Success),
		Error:      err,
	}
}

func (p *ZoneChange) ProtoMess() proto.Message {
	return p.Proto()
}
