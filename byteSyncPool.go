package pool

import (
	"sync"
)

//BytePool is a simple pool of []byte
type ByteSyncPool struct {
	sliceSize int
	pool      *sync.Pool
}

//NewBytePool...
func NewByteSyncPool(sliceSize int) *ByteSyncPool {
	return &ByteSyncPool{
		sliceSize: sliceSize,
		pool: &sync.Pool{
			New: func() interface{} {
				return make([]byte, sliceSize)
			},
		},
	}
}

//Get returns a cleared []byte
func (b *ByteSyncPool) Get() (value []byte) {
	v := b.pool.Get().([]byte)
	//clear the slice before returning it
	for i, _ := range v {
		v[i] = 0
	}
	//resize the slice in case we have been returned a partial slice
	//if the []byte given doesn't have the correct capacity we will panic!
	v = v[:b.sliceSize]
	return v
}

//Put adds a slice to the pool
func (b ByteSyncPool) Put(values []byte) {
	b.pool.Put(values)
}
