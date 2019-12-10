package day01

func FuelRequirements(mass int) int {
	return mass/3 - 2
}

func RecFuelRequirements(mass int) int {
	total := 0
	f := mass
	for {
		r := f/3 - 2
		if r <= 0 {
			break
		}
		total += r
		f = r
	}
	return total
}

func FuelRequirementsSum(ms []int, reqFunc func(int) int) int {
	sum := 0
	for _, m := range ms {
		sum += reqFunc(m)
	}
	return sum
}
