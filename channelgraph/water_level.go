package channelgraph

import (
	"fmt"
	"time"

	"github.com/guptarohit/asciigraph"
)

var Period = 100 * time.Millisecond

func MonitorSize[T any](chl chan T, close chan any) {
	points := make([]float64, 0)
	// goterm.Clear()
	// goterm.MoveCursor(0, 0)
	ticker := time.NewTicker(Period)
	if cap(chl) == 0 {
		return
	}
	max := float64(cap(chl))
	points = append(points, float64(len(chl)*10)/max)

LOG_LOOP:
	for {
		select {
		case <-close:
			ticker.Stop()
			break LOG_LOOP

		case <-ticker.C:
			points = append(points, float64(len(chl)*10)/max)
		}
	}
	fmt.Println(asciigraph.Plot(points))
}
