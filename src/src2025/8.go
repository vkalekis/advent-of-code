package src2025

import (
	"slices"
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

type jbox struct {
	x, y, z   int
	circuitId int
}

type distInfo struct {
	box1, box2 *jbox
	dist       int
}

func resolveDistances(boxes []jbox) []distInfo {
	distances := make([]distInfo, 0)
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			if i != j {
				dx := boxes[i].x - boxes[j].x
				dy := boxes[i].y - boxes[j].y
				dz := boxes[i].z - boxes[j].z

				distances = append(distances, distInfo{
					box1: &boxes[i],
					box2: &boxes[j],
					dist: dx*dx + dy*dy + dz*dz,
				})
			}
		}
	}
	slices.SortFunc(distances, func(dist1, dist2 distInfo) int {
		return dist1.dist - dist2.dist
	})
	return distances
}

func (s *Solver) Day_08(part int, reader utils.Reader) int {

	boxes := make([]jbox, 0)
	for line := range reader.Stream() {
		coords := strings.Split(line, ",")
		if len(coords) != 3 {
			logger.Fatalf("Invalid length for %s", line)
		}
		boxes = append(boxes, jbox{x: utils.ToInt(coords[0]), y: utils.ToInt(coords[1]), z: utils.ToInt(coords[2])})
	}

	logger.Debugf("Jboxes: %v", boxes)
	// logger.Debugf("Distances: %v", resolveDistances(boxes)[:10])

	iter := func(d distInfo, circuitCounter *int) {
		logger.Debugf("(bef) distance: %+v %+v %d", *d.box1, *d.box2, d.dist)
		if d.box1.circuitId == 0 && d.box2.circuitId == 0 {
			// they don't have a circuit, connect them
			*circuitCounter++
			d.box1.circuitId, d.box2.circuitId = *circuitCounter, *circuitCounter
		} else if d.box1.circuitId != 0 && d.box2.circuitId == 0 {
			// connect box2 to circuit of box1
			d.box2.circuitId = d.box1.circuitId
		} else if d.box1.circuitId == 0 && d.box2.circuitId != 0 {
			// connect box1 to circuit of box2
			d.box1.circuitId = d.box2.circuitId
		} else if d.box1.circuitId == d.box2.circuitId {
			// they are in the same circuit, pass
		} else {
			// they belong to different circuits, connect them (will skip an id but who cares)
			originalId := d.box1.circuitId
			targetId := d.box2.circuitId
			for i := range boxes {
				if boxes[i].circuitId == originalId {
					boxes[i].circuitId = targetId
				}
			}
		}
		logger.Debugf("(after) distance: %+v %+v %d", *d.box1, *d.box2, d.dist)
	}

	switch part {
	case 1:
		var circuitCounter int
		for _, d := range resolveDistances(boxes)[:1000] {
			iter(d, &circuitCounter)
		}

		logger.Debugf("Jboxes: %v", boxes)

		circuitStats := make(map[int]int)
		for _, box := range boxes {
			circuitStats[box.circuitId]++
		}

		sizes := make([]int, 0, len(circuitStats))
		for id, count := range circuitStats {
			if id != 0 {
				sizes = append(sizes, count)
			}
		}

		slices.SortFunc(sizes, func(a, b int) int {
			return b - a
		})

		top3prod := 1
		for _, v := range sizes[:3] {
			top3prod *= v
		}
		return top3prod
	case 2:
		var circuitCounter int
		var last2Xprod int
		for _, d := range resolveDistances(boxes) {
			iter(d, &circuitCounter)

			var cont bool
			for _, box := range boxes {
				if box.circuitId == 0 {
					cont = true
					break
				}
			}
			if !cont {
				logger.Infof("stopped!: %+v %+v", *d.box1, *d.box2)
				last2Xprod = d.box1.x * d.box2.x
				break
			}
		}
		return last2Xprod
	}

	return -1
}
