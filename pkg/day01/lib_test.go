package day01_test

import (
	"testing"

	"github.com/kprav33n/aoc19.go/pkg/day01"
	"github.com/stretchr/testify/assert"
)

func TestFuelRequirements(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2, day01.FuelRequirements(12))
	assert.Equal(2, day01.FuelRequirements(14))
	assert.Equal(654, day01.FuelRequirements(1969))
	assert.Equal(33583, day01.FuelRequirements(100756))
}

func TestRecFuelRequirements(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2, day01.RecFuelRequirements(14))
	assert.Equal(966, day01.RecFuelRequirements(1969))
	assert.Equal(50346, day01.RecFuelRequirements(100756))
}

func TestFuelRequirementsSum(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(34241, day01.FuelRequirementsSum([]int{12, 14, 1969, 100756},
		day01.FuelRequirements))
	assert.Equal(51314, day01.FuelRequirementsSum([]int{14, 1969, 100756},
		day01.RecFuelRequirements))
}
