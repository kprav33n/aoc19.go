package main

import (
	"fmt"
	"os"

	"github.com/kprav33n/aoc19.go/pkg/day01"
	"github.com/kprav33n/aoc19.go/pkg/util"
)

func main() {
	switch os.Args[1] {
	case "day01a":
		mods := util.ReadStdinToIntArray()
		result := day01.FuelRequirementsSum(mods, day01.FuelRequirements)
		fmt.Println(result)

	case "day01b":
		mods := util.ReadStdinToIntArray()
		result := day01.FuelRequirementsSum(mods, day01.RecFuelRequirements)
		fmt.Println(result)
	}
}
