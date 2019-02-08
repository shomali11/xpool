# xpool [![Go Report Card](https://goreportcard.com/badge/github.com/shomali11/xpool)](https://goreportcard.com/report/github.com/shomali11/xpool) [![GoDoc](https://godoc.org/github.com/shomali11/xpool?status.svg)](https://godoc.org/github.com/shomali11/xpool) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A pool to facilitate the creation and reuse of objects

# Examples

## Example 1

Using `NewPool` to create a pool with capacity of 10.

```go
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
```
