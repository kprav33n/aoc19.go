package day03_test

import (
	"testing"

	"github.com/kprav33n/aoc19.go/pkg/day03"
	"github.com/stretchr/testify/assert"
)

func TestWireIntersection(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(6, day03.WireIntersectionMinDist(
		[]string{"R8", "U5", "L5", "D3"},
		[]string{"U7", "R6", "D4", "L4"},
	))
	assert.Equal(159, day03.WireIntersectionMinDist(
		[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
		[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
	))
	assert.Equal(135, day03.WireIntersectionMinDist(
		[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
		[]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
	))
}

func TestWireIntersectionMinDistUnitError(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() {
		day03.WireIntersectionMinDist(
			[]string{"Rd4", "U5", "L5", "D3"},
			[]string{"U7", "R6", "D4", "L4"},
		)
	})
}

func TestWireIntersectionMinDistOpError(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() {
		day03.WireIntersectionMinDist(
			[]string{"R4", "U5", "L5", "D3"},
			[]string{"S7", "R6", "D4", "L4"},
		)
	})
}
