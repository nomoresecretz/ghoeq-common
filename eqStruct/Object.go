package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type Object struct {
	LinkedList   []uint32
	Unknown008   uint32
	DropId       uint32
	ZoneId       uint16
	ZoneInstance uint16
	Heading      float32
	Z            float32
	X            float32
	Y            float32
	Name         string
	ObjectType   uint32
	SpawnId      uint32

	bPointer int
}

func (p *Object) EQType() EQType { return EQT_Object }
func (p *Object) bp() *int       { return &p.bPointer }

func (p *Object) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	p.LinkedList = make([]uint32, 2)
	if err := EQReadLittleEndian(b, p, &p.LinkedList[0], 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.LinkedList[1], 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Unknown008, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.DropId, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.ZoneId, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.ZoneInstance, 0); err != nil {
		return 0, err
	}

	p.bPointer = 28
	if err := EQReadLittleEndian(b, p, &p.Heading, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Z, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.X, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Y, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Name, 16); err != nil {
		return 0, err
	}

	p.bPointer = 80
	if err := EQReadLittleEndian(b, p, &p.ObjectType, 0); err != nil {
		return 0, err
	}

	p.bPointer = 88
	if err := EQReadLittleEndian(b, p, &p.SpawnId, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *Object) Proto() *eqstruct.Object {
	return &eqstruct.Object{
		Prev:         p.LinkedList[0],
		Next:         p.LinkedList[1],
		Unknown_008:  p.Unknown008,
		DropId:       p.DropId,
		ZoneId:       uint32(p.ZoneId),
		ZoneInstance: uint32(p.ZoneInstance),
		Heading:      p.Heading,
		Z:            p.Z,
		X:            p.X,
		Y:            p.Y,
		Name:         p.Name,
		ObjectType:   p.ObjectType,
		SpawnId:      p.SpawnId,
	}
}

func (p *Object) ProtoMess() proto.Message {
	return p.Proto()
}
