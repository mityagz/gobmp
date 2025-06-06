package dumper

import (
	"log"
	"os"

	"github.com/sbezverk/gobmp/pkg/pub"
)

type msgOut struct {
	MsgType  int    `json:"msg_type,omitempty"`
	RouterID string `json:"msg_rid,omitempty"`
	MsgHash  string `json:"msg_hash,omitempty"`
	Msg      string `json:"msg_data,omitempty"`
}

type pubwriter struct {
	output *log.Logger
}

func (p *pubwriter) PublishMessage(rid string, msgType int, msgHash []byte, msg []byte) error {
	m := msgOut{
		MsgType:  msgType,
		RouterID: rid,
		MsgHash:  string(msgHash),
		Msg:      string(msg),
	}

	p.output.Printf("%+v", m)

	return nil
}

func (p *pubwriter) Stop() {
	p.output.Printf("gobmp is stopping...")
}

// NewDumper returns a new instance of standard out  dumper
func NewDumper() (pub.Publisher, error) {
	pw := pubwriter{
		output: log.New(os.Stdout, "gobmp: ", log.Lmicroseconds),
	}

	return &pw, nil
}
