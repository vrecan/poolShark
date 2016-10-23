package pool

import (
	// "fmt"
	log "github.com/cihub/seelog"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestByteSyncPool(t *testing.T) {

	defer log.Flush()
	Convey("use pool validate they aren't the same", t, func() {

		pool := NewByteSyncPool(30)
		bytes := pool.Get()
		bytes2 := pool.Get()
		So(bytes, ShouldNotEqual, bytes2)
	})

	Convey("modify pool, put back on and validate change is not visible with get", t, func() {

		pool := NewByteSyncPool(30)
		bytes := pool.Get()
		bytes[0] = byte(1)
		bytes[1] = byte(1)
		pool.Put(bytes[:2])
		bytes2 := pool.Get()
		So(bytes2, ShouldResemble, make([]byte, 30))
	})

}
