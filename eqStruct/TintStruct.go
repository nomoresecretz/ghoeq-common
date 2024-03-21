package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type TintStruct struct {
	Blue, Green, Red uint8
	Use              uint8

	bPointer int
}

func (p *TintStruct) EQType() EQType   { return EQT_TintStruct }
func (p *TintStruct) bp() *int         { return &p.bPointer }
func (p *TintStruct) SetPointer(i int) { p.bPointer = i }

func (p *TintStruct) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.Blue, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Green, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Red, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Use, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *TintStruct) Proto() *eqstruct.TintStruct {
	return &eqstruct.TintStruct{
		Red:   uint32(p.Red),
		Blue:  uint32(p.Blue),
		Green: uint32(p.Green),
		Use:   uint32(p.Use),
	}
}

func (p *TintStruct) ProtoMess() proto.Message {
	return p.Proto()
}
