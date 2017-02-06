package netio

import (
	"io"
	"net"
)

type NetWriter struct {
	host    string
	netType string
	net.Conn
}

func NewWriter(host string, netType string) *NetWriter {
	n := new(NetWriter)
	n.host = host
	n.netType = netType
	return n
}

func (n *NetWriter) Connect() error {
	Log("Connect:", n.host)
	conn, err := net.Dial(n.netType, n.host)
	n.Conn = conn
	if err != nil {
		return err
	}
	return nil
}

func (n *NetWriter) ReadFrom(r io.Reader) (int64, error) {
	var err error
	var nb int64
	buf := make([]byte, 1024)
	for {
		nr, er := r.Read(buf)
		if nr > 0 {
			nw, ew := n.Write(buf[0:nr])
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
			nb += int64(nw)
		}
		if er == io.EOF {
			break
		}
		if er != nil {
			err = er
			break
		}
	}
	return nb, err
}
