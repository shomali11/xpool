package xpool

// Pool contains items
type Pool struct {
	pool    chan interface{}
	factory func() interface{}
}

// NewPool creates a new pool of items
func NewPool(max int, factory func() interface{}) *Pool {
	return &Pool{
		pool:    make(chan interface{}, max),
		factory: factory,
	}
}

// Get an item from the pool.
func (p *Pool) Get() interface{} {
	select {
	case item := <-p.pool:
		return item
	default:
		return p.factory()
	}
}

// Return an item to the pool.
func (p *Pool) Return(item interface{}) {
	select {
	case p.pool <- item:
	default:
		// Discard item
	}
}
