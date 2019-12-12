package day02_test

import (
	"testing"

	"github.com/kprav33n/aoc19.go/pkg/day02"
	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	assert := assert.New(t)
	code := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	err := day02.Execute(code)
	assert.NoError(err)
	assert.Equal([]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, code)
}

func TestExecuteInvalidOpcode(t *testing.T) {
	assert := assert.New(t)
	code := []int{3}
	err := day02.Execute(code)
	assert.Error(err)
}
