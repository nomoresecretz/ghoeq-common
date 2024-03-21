package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type GuildEntry struct {
	GuildID uint32
	Name    string // MAX64
	Exists  uint16

	bPointer int
}

func (p *GuildEntry) EQType() EQType   { return EQT_GuildEntry }
func (p *GuildEntry) bp() *int         { return &p.bPointer }
func (p *GuildEntry) SetPointer(i int) { p.bPointer = i }

func (p *GuildEntry) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.GuildID, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Name, 64); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Exists, 0); err != nil {
		return 0, err
	}

	p.bPointer = 96
	return p.bPointer, nil
}

func (p *GuildEntry) Proto() *eqstruct.GuildEntry {
	return &eqstruct.GuildEntry{
		GuildId: p.GuildID,
		Name:    p.Name,
		Exists:  uint32(p.Exists),
	}
}

func (p *GuildEntry) ProtoMess() proto.Message {
	return p.Proto()
}
