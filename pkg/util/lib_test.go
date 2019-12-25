package util_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/kprav33n/aoc19.go/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestReadIntInputsLineSeparated(t *testing.T) {
	assert := assert.New(t)
	var input bytes.Buffer
	input.Write([]byte(" 1\n2\n-3\n0\n "))
	items := util.ReadIntInputs(&input, "\n")
	assert.Equal([]int{1, 2, -3, 0}, items)
}

func TestReadIntInputsCommaSeparated(t *testing.T) {
	assert := assert.New(t)
	var input bytes.Buffer
	input.Write([]byte("-1, 2,3,993429"))
	items := util.ReadIntInputs(&input, ",")
	assert.Equal([]int{-1, 2, 3, 993429}, items)
}

type ErroringReader struct{}

func (r *ErroringReader) Read([]byte) (int, error) {
	return 0, fmt.Errorf("forced error")
}

func TestReadIntInputsReadingPanic(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() { util.ReadIntInputs(&ErroringReader{}, ",") })
}

func TestReadIntInputsParsingPanic(t *testing.T) {
	assert := assert.New(t)
	var input bytes.Buffer
	input.Write([]byte("-1, deadbeef,3,993429"))
	assert.Panics(func() { util.ReadIntInputs(&input, ",") })
}
