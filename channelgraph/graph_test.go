package channelgraph

import (
	"math"
	"testing"
)

func TestLongGraph(t *testing.T) {
	points := make([]float64, 0)

	for i := 0; i < 10_000; i++ {
		points = append(points, 10*math.Sin(float64(i)))
	}

	printGraph( "test more", points)

}
