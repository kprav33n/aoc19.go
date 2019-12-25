package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadIntInputs(input io.Reader, sep string) []int {
	bytes, err := ioutil.ReadAll(input)
	if err != nil {
		panic(fmt.Sprintf("failed to read from stdin: %v", err))
	}

	lines := strings.Split(strings.TrimSpace(string(bytes)), sep)
	mods := make([]int, len(lines))

	for i, line := range lines {
		val, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			panic(fmt.Sprintf("failed to convert %s to int: %v", line, err))
		}

		mods[i] = val
	}

	return mods
}

func ReadStringInputs(input io.Reader, sep string) []string {
	bytes, err := ioutil.ReadAll(input)
	if err != nil {
		panic(fmt.Sprintf("failed to read from stdin: %v", err))
	}

	return strings.Split(strings.TrimSpace(string(bytes)), sep)
}
