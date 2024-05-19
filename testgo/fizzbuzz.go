package main

import "fmt"

func FizzBuzz(n int) string {
	if n%3 == 0 || n%5 == 0 {
		fizzMap := map[bool]string{true: "Fizz", false: "Buzz"}
		return fizzMap[n%3 == 0]
	}
	return fmt.Sprintf("%d", n)
}
