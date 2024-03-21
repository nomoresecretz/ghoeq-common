package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type LoginInfo struct {
	LoginInfo []byte // 000
	Account   string
	Password  string
	Zoning    uint8 // 192

	bPointer int
}

func (p *LoginInfo) EQType() EQType   { return EQT_LoginInfo }
func (p *LoginInfo) bp() *int         { return &p.bPointer }
func (p *LoginInfo) SetPointer(i int) { p.bPointer = i }

func (p *LoginInfo) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.Account, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Password, 0); err != nil {
		return 0, err
	}

	p.bPointer = 192
	if err := EQRead(b, p, &p.Zoning, 4); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *LoginInfo) Proto() *eqstruct.LoginInfo {
	return &eqstruct.LoginInfo{
		LoginInfo: p.LoginInfo,
		Account:   p.Account,
		Password:  p.Password,
		Zoning:    uint32(p.Zoning),
	}
}

func (p *LoginInfo) ProtoMess() proto.Message {
	return p.Proto()
}
