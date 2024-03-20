package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type ChannelMessage struct {
	TargetName    string // MAX64
	Sender        string // MAX64
	Language      uint16
	ChannelNum    uint16
	LanguageSkill uint16
	Message       string

	bPointer int
}

func (p *ChannelMessage) EQType() EQType { return EQT_ChannelMessage }
func (p *ChannelMessage) bp() *int       { return &p.bPointer }

func (p *ChannelMessage) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.TargetName, 64); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Sender, 64); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Language, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.ChannelNum, 0); err != nil {
		return 0, err
	}

	p.bPointer = 134
	if err := EQRead(b, p, &p.LanguageSkill, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Message, 999999999); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *ChannelMessage) Proto() *eqstruct.ChannelMessage {
	return &eqstruct.ChannelMessage{
		TargetName:    p.TargetName,
		Sender:        p.Sender,
		Language:      uint32(p.Language),
		ChannelNumber: uint32(p.ChannelNum),
		LanguageSkill: uint32(p.LanguageSkill),
		Message:       p.Message,
	}
}

func (p *ChannelMessage) ProtoMess() proto.Message {
	return p.Proto()
}
