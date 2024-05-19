package main

import "fmt"

func FizzBuzz(n int) string {
	if n%3 == 0 || n%5 == 0 {
		fizzMap := map[bool]string{true: "Fizz", false: ""}
		buzzMap := map[bool]string{true: "Buzz", false: ""}
		return fizzMap[n%3 == 0] + buzzMap[n%5 == 0]
	}
	return fmt.Sprintf("%d", n)
}
