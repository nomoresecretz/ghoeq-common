package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type PlayRequest struct {
	IP       string // 000
	bPointer int
}

func (p *PlayRequest) EQType() EQType { return EQT_PlayRequest }
func (p *PlayRequest) bp() *int       { return &p.bPointer }

func (p *PlayRequest) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.IP, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *PlayRequest) Proto() *eqstruct.PlayRequest {
	return &eqstruct.PlayRequest{
		IP: p.IP,
	}
}

func (p *PlayRequest) ProtoMess() proto.Message {
	return p.Proto()
}
