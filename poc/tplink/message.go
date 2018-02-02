package tplink

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
	"io/ioutil"
)

type Msg interface {
	Send(ip string, payload []byte) (resp []byte, err error)
}

type msg struct {
}

func NewMsg() Msg {
	return &msg{}
}

func (*msg) encrypt(payload []byte) []byte {
	n := len(payload)
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, uint32(n))
	encrypted := []byte(buf.Bytes())
	key := byte(0xAB)
	for i := 0; i != n; i++ {
		tmp := payload[i] ^ key
		key = tmp
		encrypted = append(encrypted, tmp)
	}
	return encrypted
}

func (*msg) decrypt(payload []byte) []byte {
	n := len(payload)
	decrypted := make([]byte, n)
	key := byte(0xAB)
	for i := 0; i != n; i++ {
		nextKey := payload[i]
		decrypted[i] = payload[i] ^ key
		key = nextKey
	}
	return decrypted
}

func (m *msg) Send(ip string, payload []byte) (resp []byte, err error) {
	conn, err := net.DialTimeout("tcp", ip+":9999", 10*time.Second)
	if err != nil {
		err = fmt.Errorf("cannot connnect to plug: %s", err.Error())
		return
	}
	encryptedPayload := m.encrypt(payload)
	_, err = conn.Write(encryptedPayload)
	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		err = fmt.Errorf("cannot read data from plug: %s", err.Error())
		return
	}
	resp = m.decrypt(buf)
	return
}
