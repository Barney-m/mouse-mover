package main

import (
	"flag"
	"time"

	"github.com/go-vgo/robotgo"
)

var (
	sleepDuration = flag.Int("Mouse Sleep Duration", 100, "Sleep Duration in Second")
	moveDistance  = flag.Int("Move Distance", 15, "Mouse move distance for X and Y axis")
)

func main() {
	flag.Parse()
	distance := *moveDistance

	for {
		pos1x, pos1y := robotgo.Location()
		time.Sleep(1 * time.Second)
		pos2x, pos2y := robotgo.Location()

		if checkMousePosDiff(pos1x+pos1y, pos2x+pos2y) {
			time.Sleep(time.Duration(*sleepDuration) * time.Second)
			continue
		}

		x := pos2x
		y := pos2y

		robotgo.MoveSmooth(x, y+distance) // Move up
		time.Sleep(50 * time.Millisecond)

		robotgo.MoveSmooth(x+distance, y) // Move right
		time.Sleep(50 * time.Millisecond)

		robotgo.MoveSmooth(x, y-distance) // Move down
		time.Sleep(50 * time.Millisecond)

		robotgo.MoveSmooth(x-distance, y) // Move left
		time.Sleep(50 * time.Millisecond)

		time.Sleep(time.Duration(*sleepDuration) * time.Second)
	}
}

// Return true if different
func checkMousePosDiff(sumPos1 int, sumPos2 int) bool {
	return sumPos1 != sumPos2
}
