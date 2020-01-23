package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Packet struct {
	len      int16
	version  int8
	mainCmd  uint16
	subCmd   uint16
	clientId uint32
	data     []byte
}

func New(mainCmd, subCmd uint16, clientId uint32) *Packet {
	pk := &Packet{
		version:  1,
		mainCmd:  mainCmd,
		subCmd:   subCmd,
		clientId: clientId,
	}
	return pk
}

func NewWithData(data []byte) *Packet {
	pk := &Packet{}
	buf := bytes.NewBuffer(data)
	err := pk.readHead(buf)
	if err != nil {
		return nil
	}
	pk.data = buf.Bytes()

	return pk
}

func (p *Packet) writeHead(buf *bytes.Buffer) (err error) {
	if err = binary.Write(buf, binary.LittleEndian, p.len); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.LittleEndian, p.version); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.LittleEndian, p.mainCmd); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.LittleEndian, p.subCmd); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.LittleEndian, p.clientId); err != nil {
		return err
	}

	return err
}

func (p *Packet) readHead(buf *bytes.Buffer) (err error) {
	if err = binary.Read(buf, binary.LittleEndian, &p.len); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &p.version); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &p.mainCmd); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &p.subCmd); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &p.clientId); err != nil {
		return err
	}
	return err
}

// Encode 编码数据包
func (p *Packet) Encode(data []byte) error {
	buf := &bytes.Buffer{}
	p.len = int16(len(data)) + 2 + 1 + 2 + 2 + 4
	err := p.writeHead(buf)
	if err != nil {
		return err
	}

	if len(data) > 0 {
		var n int
		n, err = buf.Write(data)
		if err != nil {
			return err
		}
		if n < 0 {

		}
	}

	p.data = buf.Bytes()

	return nil
}

func (p Packet) Version() int8 {
	return p.version
}

func (p Packet) Len() int16 {
	return p.len
}

func (p Packet) MainCmd() uint16 {
	return p.mainCmd
}

func (p Packet) SubMain() uint16 {
	return p.subCmd
}

func (p Packet) ClientId() uint32 {
	return p.clientId
}

func (p Packet) Data() []byte {
	return p.data
}

func (p Packet) String() string {
	return fmt.Sprintf("{len:%d version:%d mainCmd:%d subCmd:%d clientId:%d datalen:%d}", p.len, p.version, p.mainCmd, p.subCmd, p.clientId, len(p.data))
}
