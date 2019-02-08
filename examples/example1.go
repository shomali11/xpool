package main

import (
	"github.com/shomali11/xpool"
	"time"
)

type Object struct{}

func main() {
	factory := func() interface{} {
		return &Object{}
	}

	pool := xpool.NewPool(10, factory)

	for i := 0; i < 10; i++ {
		go func() {
			object := pool.Borrow().(*Object)
			pool.Return(object)
		}()
	}

	time.Sleep(time.Second)
}
