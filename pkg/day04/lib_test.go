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

func TestIsValidPassword2(t *testing.T) {
	assert := assert.New(t)
	assert.True(day04.IsValidPassword2("112233"))
	assert.False(day04.IsValidPassword2("123444"))
	assert.True(day04.IsValidPassword2("111122"))
	assert.True(day04.IsValidPassword2("112222"))
	assert.False(day04.IsValidPassword2("223450"))
}
