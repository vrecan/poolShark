package pool

import (
	"sync"
)

//BytePool is a simple pool of []byte
type BytePool struct {
	sliceSize int
	poolSize  int

	poolSlice [][]byte
	mutex     *sync.Mutex
}

//NewBytePool...
func NewBytePool(poolSize, sliceSize int) *BytePool {
	return &BytePool{
		sliceSize: sliceSize,
		poolSize:  poolSize,
		poolSlice: make([][]byte, 0),
		mutex:     &sync.Mutex{},
	}
}

//Get returns a cleared []byte
func (b *BytePool) Get() (value []byte) {
	b.mutex.Lock()
	if len(b.poolSlice) > 0 {
		//pop item off of slice reducing the size of the slice
		value, b.poolSlice = b.poolSlice[len(b.poolSlice)-1], b.poolSlice[:len(b.poolSlice)-1]
		b.mutex.Unlock()
		for i, _ := range value {
			value[i] = 0
		}
		//resize slice to slice size, this will panic if you give it a byte slice smaller then set size
		value = value[:b.sliceSize]

		return value
	}
	b.mutex.Unlock()
	return make([]byte, b.sliceSize)
}

//Put adds a slice to the pool
func (b *BytePool) Put(value []byte) {
	b.mutex.Lock()
	if len(b.poolSlice) < b.poolSize {
		b.poolSlice = append(b.poolSlice, value)
	}
	b.mutex.Unlock()
}

func (b BytePool) Size() int {
	b.mutex.Lock()
	s := len(b.poolSlice)
	b.mutex.Unlock()
	return s
}
