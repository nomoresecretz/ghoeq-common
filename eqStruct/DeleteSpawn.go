package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type DeleteSpawn struct {
	SpawnID uint16 // 000

	bPointer int
}

func (p *DeleteSpawn) EQType() EQType { return EQT_DeleteSpawn }
func (p *DeleteSpawn) bp() *int       { return &p.bPointer }

func (p *DeleteSpawn) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQReadLittleEndian(b, p, &p.SpawnID, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *DeleteSpawn) Proto() *eqstruct.DeleteSpawn {
	return &eqstruct.DeleteSpawn{
		SpawnId: uint32(p.SpawnID),
	}
}

func (p *DeleteSpawn) ProtoMess() proto.Message {
	return p.Proto()
}
