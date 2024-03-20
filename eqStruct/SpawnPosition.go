package eqStruct

import (
	pb "github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type SpawnPositionUpdate struct {
	SpawnID      uint16 // 000
	AnimType     int8   // 002
	Heading      uint8  // 003
	HeadingDelta int8   // 004
	Y            int16  // 005
	X            int16  // 007
	Z            int16  // 009
	XYZDelta     uint32 // 011

	bPointer int
}

func (p *SpawnPositionUpdate) EQType() EQType { return EQT_SpawnPositionUpdate }
func (p *SpawnPositionUpdate) bp() *int       { return &p.bPointer }

func (p *SpawnPositionUpdate) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0
	if err := EQRead(b, p, &p.SpawnID, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.AnimType, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Heading, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.HeadingDelta, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Y, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.X, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Z, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.XYZDelta, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

type SpawnPositionUpdates struct {
	Count   uint32                 // 000
	Updates []*SpawnPositionUpdate // 004

	bPointer int
}

func (p *SpawnPositionUpdates) EQType() EQType { return EQT_SpawnPositionUpdates }
func (p *SpawnPositionUpdates) bp() *int       { return &p.bPointer }

func (p *SpawnPositionUpdates) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.Count, 0); err != nil {
		return 0, err
	}

	p.Updates = make([]*SpawnPositionUpdate, p.Count)
	for i := range p.Count {
		upd := &SpawnPositionUpdate{}
		bp, err := upd.Unmarshal(b[p.bPointer:])
		if err != nil {
			return 0, err
		}
		p.bPointer += bp
		p.Updates[i] = upd
	}

	return p.bPointer, nil
}

func (p *SpawnPositionUpdate) Proto() *pb.SpawnPosition {
	return &pb.SpawnPosition{
		SpawnId:      uint32(p.SpawnID),
		AnimType:     int32(p.AnimType),
		Heading:      uint32(p.Heading),
		HeadingDelta: int32(p.HeadingDelta),
		Y:            int32(p.Y),
		X:            int32(p.X),
		Z:            int32(p.Z),
		Delta:        p.XYZDelta,
	}
}

func (p *SpawnPositionUpdate) ProtoMess() proto.Message {
	return p.Proto()
}

func (p *SpawnPositionUpdates) Proto() *pb.SpawnPositions {
	positions := make([]*pb.SpawnPosition, len(p.Updates))
	for i, v := range p.Updates {
		positions[i] = v.Proto()
	}

	return &pb.SpawnPositions{
		Positions: positions,
	}
}

func (p *SpawnPositionUpdates) ProtoMess() proto.Message {
	return p.Proto()
}
