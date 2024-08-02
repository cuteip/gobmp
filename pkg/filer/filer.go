package filer

import (
	"bufio"
	"os"
	"sync"

	"github.com/goccy/go-json"

	"github.com/sbezverk/gobmp/pkg/pub"
)

// MsgOut defines structure of the message stored in the file.
type MsgOut struct {
	Type  int    `json:"type,omitempty"`
	Key   []byte `json:"key,omitempty"`
	Value []byte `json:"value,omitempty"`
}

type pubfiler struct {
	writer *bufio.Writer
	m      sync.Mutex
}

func (p *pubfiler) PublishMessage(msgType int, msgHash []byte, msg []byte) error {
	m := MsgOut{
		Type:  msgType,
		Key:   msgHash,
		Value: msg,
	}
	b, err := json.Marshal(&m)
	if err != nil {
		return err
	}
	b = append(b, '\n')
	p.m.Lock()
	defer p.m.Unlock()
	_, err = p.writer.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func (p *pubfiler) Stop() {
	// ignore error
	_ = p.writer.Flush()
}

// NewFiler returns a new instance of message filer
func NewFiler(file string, bufSize int) (pub.Publisher, error) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	pw := pubfiler{
		writer: bufio.NewWriterSize(f, bufSize),
	}

	return &pw, nil
}
