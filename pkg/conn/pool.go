package conn

type Client struct {
	IdlePoolCh chan *Connection
}

func New(total int) *Client {

	pch := make(chan *Connection, total)
	p := Client{
		IdlePoolCh: pch,
	}

	for i := 0; i < total; i++ {
		pch <- new(Connection)
	}

	return &p
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
