package day03

import (
	"fmt"
	"math"
	"strconv"
)

type Point struct {
	X int
	Y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func (p *Point) ManhattanDistance(other Point) int {
	return abs(p.X-other.X) + abs(p.Y-other.Y)
}

func (p *Point) Left() Point {
	return Point{p.X - 1, p.Y}
}

func (p *Point) Right() Point {
	return Point{p.X + 1, p.Y}
}

func (p *Point) Up() Point {
	return Point{p.X, p.Y + 1}
}

func (p *Point) Down() Point {
	return Point{p.X, p.Y - 1}
}

func tracePath(path []string) map[Point]int {
	points := make(map[Point]int)
	last := Point{0, 0}
	steps := 0

	for _, inst := range path {
		unit, err := strconv.Atoi(inst[1:])
		if err != nil {
			panic(fmt.Sprintf("failed to parse instruction %s: %v", inst, err))
		}

		var next func(p Point) Point

		switch string(inst[0]) {
		case "L":
			next = func(p Point) Point { return p.Left() }
		case "R":
			next = func(p Point) Point { return p.Right() }
		case "U":
			next = func(p Point) Point { return p.Up() }
		case "D":
			next = func(p Point) Point { return p.Down() }
		default:
			panic(fmt.Sprintf("invalid instruction: %s", inst))
		}

		for i := 1; i <= unit; i++ {
			last = next(last)
			steps++
			if _, ok := points[last]; !ok {
				points[last] = steps
			}
		}
	}

	return points
}

func wireIntersectionMetric(first []string, second []string,
	strategy func(p Point, steps int) int) int {
	firstPoints := tracePath(first)
	secondPoints := tracePath(second)
	minMetric := math.MaxInt32

	for fp, fs := range firstPoints {
		ss, ok := secondPoints[fp]
		if ok {
			metric := strategy(fp, fs+ss)
			if metric < minMetric {
				minMetric = metric
			}
		}
	}

	return minMetric
}

func WireIntersectionMinDist(first []string, second []string) int {
	return wireIntersectionMetric(first, second, func(p Point, _ int) int {
		return p.ManhattanDistance(Point{0, 0})
	})
}

func WireIntersectionMinDelay(first []string, second []string) int {
	return wireIntersectionMetric(first, second, func(_ Point, steps int) int {
		return steps
	})
}
