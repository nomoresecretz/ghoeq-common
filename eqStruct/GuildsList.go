package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type GuildsList struct {
	Header []uint8
	Guilds []*GuildEntry

	bPointer int
}

const GuildEntryMax = 512

func (p *GuildsList) EQType() EQType { return EQT_GuildsList }
func (p *GuildsList) bp() *int       { return &p.bPointer }

func (p *GuildsList) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	p.Header = make([]uint8, 4)
	for i := range p.Header {
		if err := EQReadLittleEndian(b, p, &p.Header[i], 0); err != nil {
			return 0, err
		}
	}

	p.Guilds = make([]*GuildEntry, GuildEntryMax)
	for i := range p.Guilds {
		if p.bPointer == len(b) {
			return p.bPointer, nil
		}
		p.Guilds[i] = &GuildEntry{}
		bp, err := p.Guilds[i].Unmarshal(b[p.bPointer:])
		if err != nil {
			return 0, err
		}
		p.bPointer += bp
	}

	return p.bPointer, nil
}

func (p *GuildsList) Proto() *eqstruct.GuildList {
	return &eqstruct.GuildList{}
}

func (p *GuildsList) ProtoMess() proto.Message {
	return p.Proto()
}
