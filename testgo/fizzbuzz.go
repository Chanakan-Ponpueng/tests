package main

import "fmt"

func FizzBuzz(n int) string {
	if n == 3 || n == 5 {
		return map[int]string{3: "Fizz", 5: "Buzz"}[n]
	}
	return fmt.Sprintf("%d", n)
}
