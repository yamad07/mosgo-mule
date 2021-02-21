package conn

import (
	"net"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClientExec(t *testing.T) {
	_, err := net.Listen("tcp", ":3000")
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	c, err := New("localhost:3000", time.Second*10, 10)
	assert.NoError(t, err)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(data Data) {
			defer wg.Done()
			err := c.Exec(Set, data)
			assert.NoError(t, err)
		}(i)
	}
	wg.Wait()
}
