package main

import "os"

func parseNumber(s string) (int, bool) {
	if s == "" {
		return 0, false
	}
	negative := false
	if s[0] == '-' || s[0] == '+' {
		negative = s[0] == '-'
		s = s[1:]
	}
	if s == "" {
		return 0, false
	}
	max := uint(^uint(0) >> 1)
	limit := max
	if negative {
		limit++
	}
	var value uint
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		digit := uint(s[i] - '0')
		if value > (limit-digit)/10 {
			return 0, false
		}
		value = value*10 + digit
	}
	if negative {
		if value == max+1 {
			return -int(max) - 1, true
		}
		return -int(value), true
	}
	return int(value), true
}

func printNumber(n int) {
	var buffer [21]byte
	i := len(buffer)
	negative := n < 0
	var value uint
	if negative {
		value = uint(-(n + 1)) + 1
	} else {
		value = uint(n)
	}
	for {
		i--
		buffer[i] = byte(value%10) + '0'
		value /= 10
		if value == 0 {
			break
		}
	}
	if negative {
		i--
		buffer[i] = '-'
	}
	os.Stdout.Write(buffer[i:])
	os.Stdout.WriteString("\n")
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	a, okA := parseNumber(os.Args[1])
	b, okB := parseNumber(os.Args[3])
	if !okA || !okB {
		return
	}
	max := int(^uint(0) >> 1)
	min := -max - 1
	var result int
	switch os.Args[2] {
	case "+":
		if b > 0 && a > max-b || b < 0 && a < min-b {
			return
		}
		result = a + b
	case "-":
		if b < 0 && a > max+b || b > 0 && a < min+b {
			return
		}
		result = a - b
	case "*":
		if (a == min && b == -1) || (b == min && a == -1) {
			return
		}
		result = a * b
		if b != 0 && result/b != a {
			return
		}
	case "/":
		if b == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		if a == min && b == -1 {
			return
		}
		result = a / b
	case "%":
		if b == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		result = a % b
	default:
		return
	}
	printNumber(result)
}
