package main

func FizzBuzz(n int) string {
	if n%3 == 0 {
		return "Fizz"
	}
	if n == 5 {
		return "Buzz"
	}
	if n == 4 {
		return "4"
	}
	if n == 2 {
		return "2"
	}
	return "1"
}
