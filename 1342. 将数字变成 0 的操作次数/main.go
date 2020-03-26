package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(8))
}
var count = 0

func numberOfSteps (num int) int {
	deal(num)
	return count
}

func deal(num int) {
	if num <= 0 {
		return
	}
	count ++
	if num % 2 == 0 {
		num = num / 2
	} else {
		num = num - 1
	}
	deal(num)
}