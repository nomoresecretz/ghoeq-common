package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type LogServer struct {
	ShortName string // 032 Max32
	Address   string // 112 Max16
	Port      uint32 // 176

	bPointer int
}

func (p *LogServer) EQType() EQType   { return EQT_LogServer }
func (p *LogServer) bp() *int         { return &p.bPointer }
func (p *LogServer) SetPointer(i int) { p.bPointer = i }

func (p *LogServer) Unmarshal(b []byte) (int, error) {
	p.bPointer = 32
	if err := EQRead(b, p, &p.ShortName, 32); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Address, 32); err != nil {
		return 0, err
	}

	if err := EQReadBigEndian(b, p, &p.Port, 32); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *LogServer) Proto() *eqstruct.LogServer {
	return &eqstruct.LogServer{
		ShortName: p.ShortName,
		Address:   p.Address,
		Port:      p.Port,
	}
}

func (p *LogServer) ProtoMess() proto.Message {
	return p.Proto()
}
