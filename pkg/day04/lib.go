package day04

func IsValidPassword(password string) bool {
	last := '0'
	seenTwice := false

	for _, rv := range password {
		if rv < last {
			return false
		} else if rv == last {
			seenTwice = true
		}

		last = rv
	}

	return seenTwice
}

func IsValidPassword2(password string) bool {
	last := '0'
	inTwice := false
	seenTwice := make(map[rune]bool)

	for _, rv := range password {
		switch {
		case rv < last:
			return false

		case rv == last:
			if !inTwice {
				seenTwice[rv] = true
				inTwice = true
			} else {
				seenTwice[rv] = false
			}

		default:
			inTwice = false
		}

		last = rv
	}

	for _, st := range seenTwice {
		if st {
			return true
		}
	}

	return false
}
