package netio

import (
	"errors"
	"io"
	"net"
)

type OnData func([]byte, net.Addr, net.Addr, error)

type NetReader struct {
	reader     interface{}
	LocalAddr  net.Addr
	RemoteAddr net.Addr
	host       string
	netType    string
}

func NewReader(host string, netType string) *NetReader {
	n := new(NetReader)
	n.host = host
	n.netType = netType
	return n
}

func (n *NetReader) Connect() error {
	Log("Listen:", n.host)
	switch n.netType {
	case "tcp", "tcp4", "tcp6", "unix", "unixpacket":
		{
			conn, err := net.Listen(n.netType, n.host)
			if err != nil {
				return err
			}
			cc, err := conn.Accept()
			if err != nil {
				conn.Close()
				return err
			}
			n.reader = cc
			n.LocalAddr = cc.LocalAddr()
			n.RemoteAddr = cc.RemoteAddr()
		}
	case "udp", "udp4", "udp6", "ip", "ip4", "ip6", "unixgram":
		{
			conn, err := net.ListenPacket(n.netType, n.host)
			if err != nil {
				return err
			}
			n.reader = conn
			n.LocalAddr = conn.LocalAddr()
		}
	default:
		return errors.New("unknown net type: '" + n.netType + "'")
	}
	return nil
}

func (n *NetReader) WriteTo(w io.Writer) (int64, error) {
	var err error
	var nb int64
	buf := make([]byte, 1024)

	for {
		var nr int
		var er error
		switch reader := n.reader.(type) {
		case net.Conn:
			{
				nr, er = reader.Read(buf)
			}
		case net.PacketConn:
			{
				nr, n.RemoteAddr, er = reader.ReadFrom(buf)
			}
		default:
			err = errors.New("unknown connect type")
		}
		if nr > 0 {
			nw, ew := w.Write(buf[0:nr])
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

func (n *NetReader) Close() error {
	if conn, ok := n.reader.(io.Closer); ok {
		return conn.Close()
	}
	return nil
}
