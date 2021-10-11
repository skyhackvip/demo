package rolling

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	Convey("when adding values to slicewindow", t, func() {
		sw := NewSliceWindow()
		for _, x := range []float64{3, 5, 8, 1, 6, 7, 9, 2} {
			sw.Incr(x)
			time.Sleep(1 * time.Second)
		}
		Convey("the sum value should be", func() {
			So(sw.Sum(time.Now()), ShouldEqual, 25)
		})
	})
}
