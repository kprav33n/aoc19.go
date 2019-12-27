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
