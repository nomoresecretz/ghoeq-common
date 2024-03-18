package eqStruct

import (
	"fmt"

	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type ClientUpdate struct {
	ClientID string

	bPointer int
}

func (p *ClientUpdate) EQType() EQType { return EQT_ClientUpdate }
func (p *ClientUpdate) bp() *int       { return &p.bPointer }

func Unmarshal(b []byte) (int, error) {
	return 0, fmt.Errorf("this is not a real bytestream packet")
}

func (p *ClientUpdate) Proto() *eqstruct.ClientUpdate {
	return &eqstruct.ClientUpdate{
		ClientId: p.ClientID,
	}
}

func (p *ClientUpdate) ProtoMess() proto.Message {
	return p.Proto()
}
