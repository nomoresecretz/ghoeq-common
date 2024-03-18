package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type ZonePoints struct {
	Count  uint32
	Points []*ZonePoint

	bPointer int
}

func (p *ZonePoints) EQType() EQType { return EQT_ZonePoints }
func (p *ZonePoints) bp() *int       { return &p.bPointer }

func (p *ZonePoints) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQReadLittleEndian(b, p, &p.Count, 0); err != nil {
		return 0, err
	}

	p.Points = make([]*ZonePoint, p.Count)
	for i := range p.Points {
		p.Points[i] = &ZonePoint{}
		bp, err := p.Points[i].Unmarshal(b[p.bPointer:])
		if err != nil {
			return 0, err
		}
		p.bPointer += bp
	}

	return p.bPointer, nil
}

func (p *ZonePoints) Proto() *eqstruct.ZonePoints {
	points := make([]*eqstruct.ZonePoint, len(p.Points))
	for i, v := range p.Points {
		points[i] = v.Proto()
	}

	return &eqstruct.ZonePoints{
		ZonePoints: points,
	}
}

func (p *ZonePoints) ProtoMess() proto.Message {
	return p.Proto()
}

type ZonePoint struct {
	Iterator uint32
	Y        float32
	X        float32
	Z        float32
	Heading  float32
	ZoneID   uint16

	bPointer int
}

func (p *ZonePoint) EQType() EQType { return EQT_ZonePoint }
func (p *ZonePoint) bp() *int       { return &p.bPointer }

func (p *ZonePoint) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQReadLittleEndian(b, p, &p.Iterator, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Y, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.X, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Z, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Heading, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.ZoneID, 0); err != nil {
		return 0, err
	}

	p.bPointer = 24
	return p.bPointer, nil
}

func (p *ZonePoint) Proto() *eqstruct.ZonePoint {
	return &eqstruct.ZonePoint{
		Iterator: p.Iterator,
		Y:        p.Y,
		X:        p.X,
		Z:        p.Z,
		Heading:  p.Heading,
		ZoneId:   uint32(p.ZoneID),
	}
}

func (p *ZonePoint) ProtoMess() proto.Message {
	return p.Proto()
}
