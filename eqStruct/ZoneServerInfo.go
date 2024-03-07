package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type ZoneServerInfo struct {
	IP       string // 000
	Port     uint16 // 128
	bPointer int
}

func (p *ZoneServerInfo) EQType() EQType { return EQT_ZoneServerInfo }
func (p *ZoneServerInfo) bp() *int       { return &p.bPointer }

func (p *ZoneServerInfo) Unmarshal(b []byte) error {
	p.bPointer = 0

	if err := EQRead(b, p, &p.IP, 128); err != nil {
		return err
	}

	if err := EQRead(b, p, &p.Port, 0); err != nil {
		return err
	}

	return nil
}

func (p *ZoneServerInfo) Proto() *eqstruct.ZoneServerInfo {
	return &eqstruct.ZoneServerInfo{
		Ip:   p.IP,
		Port: uint32(p.Port),
	}
}

func (p *ZoneServerInfo) ProtoMess() proto.Message {
	return p.Proto()
}
