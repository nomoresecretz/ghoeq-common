package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type Time struct {
	Hour   uint8
	Minute uint8
	Day    uint8
	Month  uint8
	Year   uint8

	bPointer int
}

func (p *Time) EQType() EQType { return EQT_Time }
func (p *Time) bp() *int       { return &p.bPointer }

func (p *Time) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQReadLittleEndian(b, p, &p.Hour, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Minute, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Day, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Month, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Year, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *Time) Proto() *eqstruct.Time {
	return &eqstruct.Time{
		Hour:   uint32(p.Hour),
		Minute: uint32(p.Minute),
		Day:    uint32(p.Day),
		Month:  uint32(p.Month),
		Year:   uint32(p.Year),
	}
}

func (p *Time) ProtoMess() proto.Message {
	return p.Proto()
}
