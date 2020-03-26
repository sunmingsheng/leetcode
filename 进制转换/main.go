package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//十进制转二进制
	n := 13
	result := ""
	for ; n > 0; n /= 2 {
		result = strconv.Itoa(n%2) + result
	}
	fmt.Println(result)

	//十进制转8进制
	n = 13
	result = ""
	for ; n > 0; n /= 8 {
		result = strconv.Itoa(n%8) + result
	}
	fmt.Println(result)

	//十进制转16进制
	n = 13
	result = ""
	for ; n > 0; n /= 16 {
		temp := n % 16
		switch temp {
		case 10:
			result = "A" + result
		case 11:
			result = "B" + result
		case 12:
			result = "C" + result
		case 13:
			result = "D" + result
		case 14:
			result = "E" + result
		case 15:
			result = "F" + result
		default:
			result = strconv.Itoa(temp) + result
		}
	}
	fmt.Println(result)

	//二进制转十进制
	result = "1101"
	var value float64
	runeArr := []rune(result)
	for n := len(runeArr) - 1; n >= 0; n-- {
		index := len(result) - n - 1
		temp, _ := strconv.Atoi(string(result[n]))
		value += float64(temp) * math.Pow(2, float64(index))
	}
	fmt.Println(int32(value))

	//16进制转十进制
	result = "D"
	value = 0
	runeArr = []rune(result)
	for n := len(runeArr) - 1; n >= 0; n-- {
		index := len(result) - n - 1
		temp := 0
		switch string(result[n]) {
		case "A":
			temp = 10
		case "B":
			temp = 11
		case "C":
			temp = 12
		case "D":
			temp = 13
		case "E":
			temp = 14
		case "F":
			temp = 15
		default:
			temp, _ = strconv.Atoi(string(result[n]))
		}
		value += float64(temp) * math.Pow(16, float64(index))
	}
	fmt.Println(value)

}
