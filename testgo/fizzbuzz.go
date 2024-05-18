package main

import "fmt"

var fizzMap map[bool]string = map[bool]string{true: "Fizz", false: ""}
var buzzMap map[bool]string = map[bool]string{true: "Buzz", false: ""}

func FizzBuzz(n int) string {
	result := ""
	result += fizzMap[n%3 == 0]
	result += buzzMap[n%5 == 0]
	if result == "" {
		return fmt.Sprintf("%d", n)
	}
	return result
}
