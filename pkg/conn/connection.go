package conn

import (
	"encoding/json"
	"net"
)

type Connection struct {
	lock bool
	conn net.Conn
}

type Request struct {
	Command Command `json:"command"`
	Data    Data    `json:"data"`
}

func (c Connection) Exec(cmd Command, data Data) error {
	req := Request{
		Command: cmd,
		Data:    data,
	}
	err := json.NewEncoder(c.conn).Encode(req)
	if err != nil {
		return err
	}

	return nil
}
