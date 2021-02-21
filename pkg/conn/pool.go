package conn

import (
	"net"
	"time"
)

type Client struct {
	IdlePoolCh chan *Connection
}

func New(addr string, timeout time.Duration, total int) (*Client, error) {

	pch := make(chan *Connection, total)
	p := Client{
		IdlePoolCh: pch,
	}

	for i := 0; i < total; i++ {
		conn, err := net.DialTimeout("tcp", addr, timeout)
		if err != nil {
			return nil, err
		}
		pch <- &Connection{
			conn: conn,
		}
	}

	return &p, nil
}

func (p Client) Exec(cmd Command, data Data) error {
	for {
		select {
		case c := <-p.IdlePoolCh:
			c.Exec(cmd, data)
			p.IdlePoolCh <- c
			return nil
		}
	}
}
