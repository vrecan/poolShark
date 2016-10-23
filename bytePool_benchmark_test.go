package pool

import (
	// "fmt"
	log "github.com/cihub/seelog"
	"math/rand"
	"sync"
	"testing"
)

const (
	oneMB       = 100000
	hundredKB   = 100000
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()-+1234567890"
)

//BENCH BytePool
func BenchmarkBytePoolSizehundredKBInflight100(b *testing.B) {
	defer log.Flush()
	p := NewBytePool(1000, hundredKB)
	c := makeCopySlice(hundredKB)
	for i := 0; i < b.N; i++ {
		benchBytePool(p, 100, c)
	}
}

func BenchmarkBytePoolSizeoneMBInflight100(b *testing.B) {
	defer log.Flush()
	p := NewBytePool(1000, oneMB)
	c := makeCopySlice(oneMB)
	for i := 0; i < b.N; i++ {
		benchBytePool(p, 100, c)
	}
}

func BenchmarkBytePoolSizeoneMBInflight10000(b *testing.B) {
	defer log.Flush()
	p := NewBytePool(1000, oneMB)
	c := makeCopySlice(oneMB)
	for i := 0; i < b.N; i++ {
		benchBytePool(p, 10000, c)
	}
}

func BenchmarkBytePool2Concurrent(b *testing.B) {
	defer log.Flush()
	p := NewBytePool(1000, oneMB)
	c := makeCopySlice(oneMB)
	for i := 0; i < b.N; i++ {
		runMany(2, func() {
			benchBytePool(p, 10000, c)
		})
	}
}

func BenchmarkBytePool5Concurrent(b *testing.B) {
	defer log.Flush()
	p := NewBytePool(1000, oneMB)
	c := makeCopySlice(oneMB)
	for i := 0; i < b.N; i++ {
		runMany(5, func() {
			benchBytePool(p, 10000, c)
		})
	}
}

func BenchmarkBytePool10Concurrent(b *testing.B) {
	defer log.Flush()
	p := NewBytePool(1000, oneMB)
	c := makeCopySlice(oneMB)
	for i := 0; i < b.N; i++ {
		runMany(10, func() {
			benchBytePool(p, 10000, c)
		})
	}
}

//BENCH ByteSyncPool

func BenchmarkByteSyncPoolSizehundredKBInflight100(b *testing.B) {
	defer log.Flush()
	p := NewByteSyncPool(hundredKB)
	c := makeCopySlice(hundredKB)
	for i := 0; i < b.N; i++ {
		benchByteSyncPool(p, 100, c)
	}
}

func BenchmarkByteSyncPoolSizeoneMBInflight100(b *testing.B) {
	defer log.Flush()
	p := NewByteSyncPool(oneMB)
	c := makeCopySlice(oneMB)
	for i := 0; i < b.N; i++ {
		benchByteSyncPool(p, 100, c)
	}
}

func BenchmarkByteSyncPoolSizeoneMBInflight10000(b *testing.B) {
	defer log.Flush()
	p := NewByteSyncPool(oneMB)
	c := makeCopySlice(oneMB)
	for i := 0; i < b.N; i++ {
		benchByteSyncPool(p, 10000, c)
	}
}

func BenchmarkByteSyncPool2Concurrent(b *testing.B) {
	defer log.Flush()
	p := NewByteSyncPool(oneMB)
	c := makeCopySlice(oneMB)
	for i := 0; i < b.N; i++ {
		runMany(2, func() {
			benchByteSyncPool(p, 10000, c)
		})
	}
}

func BenchmarkByteSyncPool5Concurrent(b *testing.B) {
	defer log.Flush()
	p := NewByteSyncPool(oneMB)
	c := makeCopySlice(oneMB)
	for i := 0; i < b.N; i++ {
		runMany(5, func() {
			benchByteSyncPool(p, 10000, c)
		})
	}
}

func BenchmarkByteSyncPool10Concurrent(b *testing.B) {
	defer log.Flush()
	p := NewByteSyncPool(oneMB)
	c := makeCopySlice(oneMB)
	for i := 0; i < b.N; i++ {
		runMany(10, func() {
			benchByteSyncPool(p, 10000, c)
		})
	}
}

//BENCH ByteChanPool
func BenchmarkByteChanPoolSizehundredKBInflight100(b *testing.B) {
	defer log.Flush()
	p := NewByteChanPool(1000, hundredKB)
	c := makeCopySlice(hundredKB)
	for i := 0; i < b.N; i++ {
		benchByteChanPool(p, 100, c)
	}
}

func BenchmarkByteChanPoolSizeoneMBInflight100(b *testing.B) {
	defer log.Flush()
	p := NewByteChanPool(1000, oneMB)
	c := makeCopySlice(oneMB)
	for i := 0; i < b.N; i++ {
		benchByteChanPool(p, 100, c)
	}
}

func BenchmarkByteChanPoolSizeoneMBInflight10000(b *testing.B) {
	defer log.Flush()
	p := NewByteChanPool(1000, oneMB)
	c := makeCopySlice(oneMB)
	for i := 0; i < b.N; i++ {
		benchByteChanPool(p, 10000, c)
	}
}

func BenchmarkByteChanPool2Concurrent(b *testing.B) {
	defer log.Flush()
	p := NewByteChanPool(1000, oneMB)
	c := makeCopySlice(oneMB)
	for i := 0; i < b.N; i++ {
		runMany(2, func() {
			benchByteChanPool(p, 10000, c)
		})
	}
}

func BenchmarkByteChanPool5Concurrent(b *testing.B) {
	defer log.Flush()
	p := NewByteChanPool(1000, oneMB)
	c := makeCopySlice(oneMB)
	for i := 0; i < b.N; i++ {
		runMany(5, func() {
			benchByteChanPool(p, 10000, c)
		})
	}
}

func BenchmarkByteChanPool10Concurrent(b *testing.B) {
	defer log.Flush()
	p := NewByteChanPool(1000, oneMB)
	c := makeCopySlice(oneMB)
	for i := 0; i < b.N; i++ {
		runMany(10, func() {
			benchByteChanPool(p, 10000, c)
		})
	}
}

func runMany(n int, f func()) {
	wg := &sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			f()
			wg.Done()
		}()
	}
	wg.Wait()

}

func makeCopySlice(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func benchBytePool(p *BytePool, inflight int, cp []byte) {
	list := make([][]byte, inflight)
	for i := 0; inflight < i; i++ {
		s := p.Get()
		copy(s, cp)
		list = append(list, s)
	}
	for i, _ := range list {
		p.Put(list[i])
	}
}

func benchByteChanPool(p *ByteChanPool, inflight int, cp []byte) {
	list := make([][]byte, inflight)
	for i := 0; inflight < i; i++ {
		s := p.Get()
		copy(s, cp)
		list = append(list, s)
	}
	for i, _ := range list {
		p.Put(list[i])
	}
}

func benchByteSyncPool(p *ByteSyncPool, inflight int, cp []byte) {
	list := make([][]byte, inflight)
	for i := 0; inflight < i; i++ {
		s := p.Get()
		copy(s, cp)
		list = append(list, s)
	}
	for i, _ := range list {
		p.Put(list[i])
	}

}
