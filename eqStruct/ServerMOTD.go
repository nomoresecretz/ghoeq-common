package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type ServerMOTD struct {
	Name    string // 000
	Message string

	bPointer int
}

func (p *ServerMOTD) EQType() EQType { return EQT_ServerMOTD }
func (p *ServerMOTD) bp() *int       { return &p.bPointer }

func (p *ServerMOTD) Unmarshal(b []byte) error {
	p.bPointer = 0

	if err := EQReadLittleEndian(b, p, &p.Name, 64); err != nil {
		return err
	}

	if err := EQReadLittleEndian(b, p, &p.Message, 512); err != nil {
		return err
	}

	return nil
}

func (p *ServerMOTD) Proto() *eqstruct.ServerMOTD {
	return &eqstruct.ServerMOTD{
		Name:    p.Name,
		Message: p.Message,
	}
}

func (p *ServerMOTD) ProtoMess() proto.Message {
	return p.Proto()
}
