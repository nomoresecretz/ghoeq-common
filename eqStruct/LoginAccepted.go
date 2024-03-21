package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type LoginAccepted struct {
	Account string // 00 Max 10

	bPointer int
}

func (p *LoginAccepted) EQType() EQType   { return EQT_LoginAccepted }
func (p *LoginAccepted) bp() *int         { return &p.bPointer }
func (p *LoginAccepted) SetPointer(i int) { p.bPointer = i }

func (p *LoginAccepted) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0
	if err := EQRead(b, p, &p.Account, 10); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *LoginAccepted) Proto() *eqstruct.LoginInfo {
	return &eqstruct.LoginInfo{
		Account: p.Account,
	}
}

func (p *LoginAccepted) ProtoMess() proto.Message {
	return p.Proto()
}
