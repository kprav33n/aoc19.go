package day02

import "fmt"

func Execute(code []int) error {
	for i := 0; ; i += 4 {
		switch code[i] {
		case 1:
			code[code[i+3]] = code[code[i+2]] + code[code[i+1]]
		case 2:
			code[code[i+3]] = code[code[i+2]] * code[code[i+1]]
		case 99:
			return nil
		default:
			return fmt.Errorf("invalid op code encountered: %d", code[i])
		}
	}
}
