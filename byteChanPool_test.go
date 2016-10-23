package pool

import (
	// "fmt"
	log "github.com/cihub/seelog"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestByteChanPool(t *testing.T) {

	defer log.Flush()
	Convey("use pool validate they aren't the same", t, func() {

		pool := NewByteChanPool(5, 30)
		bytes := pool.Get()
		bytes2 := pool.Get()
		So(bytes, ShouldNotEqual, bytes2)
	})

	Convey("modify pool, put back on and validate change is not visible with get", t, func() {

		pool := NewByteChanPool(5, 30)
		bytes := pool.Get()
		bytes[0] = byte(1)
		bytes[1] = byte(1)
		pool.Put(bytes[:2])
		bytes2 := pool.Get()
		So(bytes2, ShouldResemble, make([]byte, 30))
	})

	Convey("exceed put buffer validate length", t, func() {
		pool := NewByteChanPool(1, 30)
		So(pool.Size(), ShouldEqual, 0)
		bytes := pool.Get()
		bytes2 := pool.Get()
		pool.Put(bytes)
		So(pool.Size(), ShouldEqual, 1)
		pool.Put(bytes2)
		So(pool.Size(), ShouldEqual, 1)
		bytes3 := pool.Get()
		So(bytes3, ShouldResemble, make([]byte, 30))
		So(pool.Size(), ShouldEqual, 0)
	})

}
