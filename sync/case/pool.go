package _case

import (
	"math/rand"
	"sync"
)

func PoolCase() {

}

const (
	ON  = 1
	OFF = 0
)

type Conn struct {
	ID     int64
	Target string
	Status int
}

func NewConn(target string) *Conn {
	return &Conn{
		ID:     rand.Int63(),
		Target: target,
		Status: ON,
	}
}

func (c *Conn) GetStatus() int {
	return c.Status
}

// 内嵌pool方法
type ConnPool struct {
	sync.Pool
}

func GetPool(target string) (*ConnPool, error) {
	return &ConnPool{
		Pool: sync.Pool{
			New: func() any {
				return NewConn(target)
			},
		},
	}, nil
}
func (c *ConnPool) Get() *Conn {
	conn := c.Pool.Get().(*Conn)
	if conn.GetStatus() == OFF {
		conn = c.Pool.New().(*Conn)
	}
	return conn
}
func (c *ConnPool) Put(conn *Conn) {
	if conn.GetStatus() == OFF {
		return
	}
	c.Pool.Put(conn)
}
