package conn

import (
	"net"
)

type Connection struct {
	lock bool
	conn net.Conn
}

func (c Connection) Exec(cmd Command, data Data) {
}
