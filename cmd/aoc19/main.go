package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kprav33n/aoc19.go/pkg/day01"
	"github.com/kprav33n/aoc19.go/pkg/day02"
	"github.com/kprav33n/aoc19.go/pkg/day03"
	"github.com/kprav33n/aoc19.go/pkg/day04"
	"github.com/kprav33n/aoc19.go/pkg/util"
)

func main() {
	dispatcher := map[string]func(){
		"day01a": day01a,
		"day01b": day01b,
		"day02a": day02a,
		"day02b": day02b,
		"day03a": day03a,
		"day03b": day03b,
		"day04a": day04a,
	}

	f, ok := dispatcher[os.Args[1]]
	if !ok {
		f = unimplemented
	}

	f()
}

func unimplemented() {
	panic("nothing to see here!")
}

func day01a() {
	mods := util.ReadIntInputs(os.Stdin, "\n")
	result := day01.FuelRequirementsSum(mods, day01.FuelRequirements)
	fmt.Println(result)
}

func day01b() {
	mods := util.ReadIntInputs(os.Stdin, "\n")
	result := day01.FuelRequirementsSum(mods, day01.RecFuelRequirements)
	fmt.Println(result)
}

func day02a() {
	code := util.ReadIntInputs(os.Stdin, ",")
	code[1] = 12
	code[2] = 2

	err := day02.Execute(code)
	if err != nil {
		panic(err)
	}

	fmt.Println(code[0])
}

func day02b() {
	code := util.ReadIntInputs(os.Stdin, ",")

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			inst := make([]int, len(code))
			copy(inst, code)
			inst[1] = noun
			inst[2] = verb

			err := day02.Execute(inst)
			if err != nil {
				panic(err)
			}

			if inst[0] == 19690720 {
				fmt.Println(100*noun + verb)
				return
			}
		}
	}

	panic("unable to arrive at a solution")
}

func day03a() {
	lines := util.ReadStringInputs(os.Stdin, "\n")
	first := strings.Split(strings.TrimSpace(lines[0]), ",")
	second := strings.Split(strings.TrimSpace(lines[1]), ",")
	result := day03.WireIntersectionMinDist(first, second)
	fmt.Println(result)
}

func day03b() {
	lines := util.ReadStringInputs(os.Stdin, "\n")
	first := strings.Split(strings.TrimSpace(lines[0]), ",")
	second := strings.Split(strings.TrimSpace(lines[1]), ",")
	result := day03.WireIntersectionMinDelay(first, second)
	fmt.Println(result)
}

func day04a() {
	items := util.ReadIntInputs(os.Stdin, "-")
	numValid := 0
	for i := items[0]; i <= items[1]; i++ {
		password := strconv.FormatInt(int64(i), 10)
		if day04.IsValidPassword(password) {
			numValid++
		}
	}
	fmt.Println(numValid)
}
