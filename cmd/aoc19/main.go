package main

import (
	"fmt"
	"os"

	"github.com/kprav33n/aoc19.go/pkg/day01"
	"github.com/kprav33n/aoc19.go/pkg/day02"
	"github.com/kprav33n/aoc19.go/pkg/util"
)

func main() {
	switch os.Args[1] {
	case "day01a":
		mods := util.ReadIntInputs(os.Stdin, "\n")
		result := day01.FuelRequirementsSum(mods, day01.FuelRequirements)
		fmt.Println(result)

	case "day01b":
		mods := util.ReadIntInputs(os.Stdin, "\n")
		result := day01.FuelRequirementsSum(mods, day01.RecFuelRequirements)
		fmt.Println(result)

	case "day02a":
		code := util.ReadIntInputs(os.Stdin, ",")
		code[1] = 12
		code[2] = 2

		err := day02.Execute(code)
		if err != nil {
			panic(err)
		}

		fmt.Println(code[0])

	case "day02b":
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

	default:
		panic(fmt.Sprintf("unhandled argument: %s", os.Args[1]))
	}
}
