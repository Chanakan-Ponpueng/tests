package main

import "fmt"

func FizzBuzz(n int) string {
	results := map[bool]string{
		true : "Fizz", 
		false : fmt.Sprintf("%d",n),
	}
	return results[ n == 3 ] 
}
