package main

import "fmt"

func FizzBuzz(n int) string {

	fizzMap := map[bool]string{
		true : "Fizz", 
		false : "",
	}
	buzzMap := map[bool]string{
		true : "Buzz", 
		false : "",
	}
	result := fmt.Sprintf("%s%s",fizzMap[n==3],buzzMap[n==5])
	resultMap := map[bool]string{
		true: fmt.Sprintf("%d",n),
		false: result,
	}
	return resultMap[result == ""]
}
