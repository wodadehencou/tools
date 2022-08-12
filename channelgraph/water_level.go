package channelgraph

import (
	"fmt"
	"time"

	"github.com/guptarohit/asciigraph"
)

var Period = 100 * time.Millisecond
var GraphWidth = 80

func MonitorSize[T any](name string, chl chan T, close chan any) {
	points := make([]float64, 0)
	// goterm.Clear()
	// goterm.MoveCursor(0, 0)
	ticker := time.NewTicker(Period)
	if cap(chl) == 0 {
		return
	}
	max := float64(cap(chl))
	points = append(points, float64(len(chl)*10)/max)
	readyToFinish := false

LOG_LOOP:
	for {
		select {
		case <-close:
			readyToFinish = true

		case <-ticker.C:
			size := len(chl)
			points = append(points, float64(size*10)/max)
			if readyToFinish && (size == 0) {
				break LOG_LOOP
			}
		}
	}

	// fmt.Println(len(points))
	printGraph(name, points)
}

func printGraph(name string, points []float64) {
	graphData := make([]float64, 0, GraphWidth)
	step := len(points) / GraphWidth
	moreStep := len(points) - GraphWidth*step

	graphData = append(graphData, points[0])
	cur := 0
	added := 0
	for cur < len(points)-1 {
		// fmt.Printf("cur = %d, step = %d, added = %d \n", cur, step, added)
		graphData = append(graphData, points[cur])
		cur += step
		added++

		if added < moreStep {
			cur++
		}
	}

	// fmt.Println(len(graphData))
	fmt.Println(asciigraph.Plot(graphData, asciigraph.Caption(name), asciigraph.Precision(0)))
}
