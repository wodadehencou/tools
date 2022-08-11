package main

import (
	"time"

	"github.com/wodadehencou/tools/channelgraph"
)

func main() {
	TestMonitorSize()
}

func TestMonitorSize() {

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
				var ok bool
				v, ok = <-chl
				if !ok {
					return
				}
			}
			time.Sleep(time.Duration(v) * time.Millisecond)
		}
	}()

	channelgraph.MonitorSize(chl, stop)
	time.Sleep(time.Second)
}
