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
		mods := util.ReadStdinToIntArray("\n")
		result := day01.FuelRequirementsSum(mods, day01.FuelRequirements)
		fmt.Println(result)

	case "day01b":
		mods := util.ReadStdinToIntArray("\n")
		result := day01.FuelRequirementsSum(mods, day01.RecFuelRequirements)
		fmt.Println(result)

	case "day02a":
		code := util.ReadStdinToIntArray(",")
		code[1] = 12
		code[2] = 2
		err := day02.Execute(code)
		if err != nil {
			panic(err)
		}
		fmt.Println(code[0])

	default:
		panic(fmt.Sprintf("unhandled argument: %s", os.Args[1]))
	}
}
