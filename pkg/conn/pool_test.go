package conn

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientExec(t *testing.T) {

	var wg sync.WaitGroup
	c := New(10)
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
