package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type SpawnDoor struct {
	Name        string // MAX16
	Y           float32
	X           float32
	Z           float32
	Heading     float32
	Incline     uint16
	Size        uint16
	Unknown_036 uint16
	DoorID      uint8
	OpenType    uint8
	OpenFlag    uint8
	Inverted    uint8
	Parameter   uint16

	bPointer int
}

func (p *SpawnDoor) EQType() EQType { return EQT_SpawnDoor }
func (p *SpawnDoor) bp() *int       { return &p.bPointer }

func (p *SpawnDoor) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQReadLittleEndian(b, p, &p.Name, 16); err != nil {
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

	if err := EQReadLittleEndian(b, p, &p.Incline, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Size, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Unknown_036, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.DoorID, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.OpenType, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.OpenFlag, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Inverted, 0); err != nil {
		return 0, err
	}

	if err := EQReadLittleEndian(b, p, &p.Parameter, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *SpawnDoor) Proto() *eqstruct.SpawnDoor {
	return &eqstruct.SpawnDoor{
		Name:      p.Name,
		Y:         p.Y,
		X:         p.X,
		Z:         p.Z,
		Heading:   p.Heading,
		Incline:   uint32(p.Incline),
		Size:      uint32(p.Size),
		DoorId:    uint32(p.DoorID),
		OpenType:  uint32(p.OpenType),
		OpenFlag:  uint32(p.OpenFlag),
		Inverted:  uint32(p.Inverted),
		Parameter: uint32(p.Parameter),
	}
}

func (p *SpawnDoor) ProtoMess() proto.Message {
	return p.Proto()
}

type SpawnDoors struct {
	Doors []*SpawnDoor // 004

	bPointer int
}

func (p *SpawnDoors) EQType() EQType { return EQT_SpawnDoors }
func (p *SpawnDoors) bp() *int       { return &p.bPointer }

func (p *SpawnDoors) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	p.Doors = make([]*SpawnDoor, 0)
	for len(b)-p.bPointer >= 224 {
		spawn := &SpawnDoor{}
		bp, err := spawn.Unmarshal(b[p.bPointer:])
		if err != nil {
			return 0, err
		}
		p.bPointer += bp
		p.Doors = append(p.Doors, spawn)
	}

	return p.bPointer, nil
}

func (p *SpawnDoors) Proto() *eqstruct.SpawnDoors {
	arr := make([]*eqstruct.SpawnDoor, len(p.Doors))
	for i := range arr {
		arr[i] = p.Doors[i].Proto()
	}

	return &eqstruct.SpawnDoors{
		Doors: arr,
	}
}

func (p *SpawnDoors) ProtoMess() proto.Message {
	return p.Proto()
}
