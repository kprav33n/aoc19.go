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

func tracePath(path []string) map[Point]struct{} {
	points := make(map[Point]struct{})
	last := Point{0, 0}

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
			points[last] = struct{}{}
		}
	}

	return points
}

func WireIntersectionMinDist(first []string, second []string) int {
	firstPoints := tracePath(first)
	secondPoints := tracePath(second)
	minDist := math.MaxInt32

	for fp := range firstPoints {
		_, ok := secondPoints[fp]
		if ok {
			dist := fp.ManhattanDistance(Point{0, 0})
			if dist < minDist {
				minDist = dist
			}
		}
	}

	return minDist
}
