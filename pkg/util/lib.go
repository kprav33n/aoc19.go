package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadStdinToIntArray(sep string) []int {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(fmt.Sprintf("failed to read from stdin: %v", err))
	}

	lines := strings.Split(strings.TrimSpace(string(bytes)), sep)
	mods := make([]int, len(lines), len(lines))
	for i, line := range lines {
		val, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Sprintf("failed to convert %s to int: %v", line, err))
		}
		mods[i] = val
	}
	return mods
}
