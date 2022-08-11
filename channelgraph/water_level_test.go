package channelgraph_test

import (
	"testing"
	"time"

	"github.com/wodadehencou/tools/channelgraph"
)

func TestMonitorSize(t *testing.T) {

	chl := make(chan int, 100)
	stop := make(chan any)

	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(time.Millisecond)
			chl <- i
		}
		close(chl)
		close(stop)
	}()

	go func() {
		for {
			v := 0
			for i := 0; i < 100; i++ {
				time.Sleep(time.Duration(v) * time.Microsecond)
				var ok bool
				v, ok = <-chl
				if !ok {
					return
				}
			}
			time.Sleep(time.Duration(v) * time.Millisecond)
		}
	}()

	channelgraph.MonitorSize("test", chl, stop)
	time.Sleep(time.Second)
}
