package main

import (
	"flag"
	"time"

	"github.com/go-vgo/robotgo"
)

var (
	sleepDuration = flag.Int("interval", 120, "Sleep Duration in Second")            // Default 2 minutes
	moveDistance  = flag.Int("distance", 15, "Mouse move distance for X and Y axis") // Distance default to 15 x and y axis
)

func main() {
	flag.Parse()
	distance := *moveDistance

	for {
		pos1x, pos1y := robotgo.Location()
		time.Sleep(1 * time.Second)
		pos2x, pos2y := robotgo.Location()

		// If mouse active then ignore current instance
		if checkMousePosDiff(pos1x+pos1y, pos2x+pos2y) {
			time.Sleep(time.Duration(*sleepDuration) * time.Second)
			continue
		}

		x := pos2x
		y := pos2y

		robotgo.MoveSmooth(x, y+distance) // Move up
		robotgo.MoveSmooth(x+distance, y) // Move right
		robotgo.MoveSmooth(x, y-distance) // Move down
		robotgo.MoveSmooth(x-distance, y) // Move left

		time.Sleep(time.Duration(*sleepDuration) * time.Second)
	}
}

// Check mouse position of previous and current see whether it is different
//
// Return true if different
func checkMousePosDiff(sumPos1 int, sumPos2 int) bool {
	return sumPos1 != sumPos2
}
