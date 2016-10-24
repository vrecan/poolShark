package pool

//ByteChanPool is a simple pool of []byte using a channel
type ByteChanPool struct {
	sliceSize int
	poolSize  int
	pool      chan []byte
}

//NewByteChanPool...
func NewByteChanPool(poolSize, sliceSize int) *ByteChanPool {
	return &ByteChanPool{
		sliceSize: sliceSize,
		poolSize:  poolSize,
		pool:      make(chan []byte, poolSize),
	}
}

//Get returns a cleared []byte
func (b *ByteChanPool) Get() (value []byte) {
	select {
	case value = <-b.pool:
		//this will be optimied into a memclr call by the compiler
		for i, _ := range value {
			value[i] = 0
		}
		//resize the slice, this will panic if the slice doesn't have the capacity
		value = value[:b.sliceSize]
		return value
	default:
		return make([]byte, b.sliceSize)
	}
}

//Put adds a slice to the pool
func (b *ByteChanPool) Put(value []byte) {
	select {
	case b.pool <- value:
		//put on pool
	default:
		//drop value
	}
	return
}

//Size...
func (b ByteChanPool) Size() int {
	return len(b.pool)
}
