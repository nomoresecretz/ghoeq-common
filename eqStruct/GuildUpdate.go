package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type GuildUpdate struct {
	GuildID uint32
	Guild   *GuildEntry

	bPointer int
}

func (p *GuildUpdate) EQType() EQType   { return EQT_GuildUpdate }
func (p *GuildUpdate) bp() *int         { return &p.bPointer }
func (p *GuildUpdate) SetPointer(i int) { p.bPointer = i }

func (p *GuildUpdate) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.GuildID, 0); err != nil {
		return 0, err
	}

	p.Guild = &GuildEntry{}
	bp, err := p.Guild.Unmarshal(b[p.bPointer:])
	if err != nil {
		return 0, err
	}
	p.bPointer += bp

	return p.bPointer, nil
}

func (p *GuildUpdate) Proto() *eqstruct.GuildUpdate {
	return &eqstruct.GuildUpdate{
		GuildId: p.GuildID,
		Guild:   p.Guild.Proto(),
	}
}

func (p *GuildUpdate) ProtoMess() proto.Message {
	return p.Proto()
}
