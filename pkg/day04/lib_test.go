package day04_test

import (
	"testing"

	"github.com/kprav33n/aoc19.go/pkg/day04"
	"github.com/stretchr/testify/assert"
)

func TestIsValidPassword(t *testing.T) {
	assert := assert.New(t)
	assert.True(day04.IsValidPassword("111111"))
	assert.False(day04.IsValidPassword("223450"))
	assert.False(day04.IsValidPassword("123789"))
}
