package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(input string) bool {
	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))

	panjang := len(input)

	for i := 0; i < panjang/2; i++ {
		if input[i] != input[panjang-i-1] {
			return false
		}
	}

	return true
}

func FindMaxValue(arr [5]int) int {

	maxValue := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}

	return maxValue
}

func MakeStar(input int) {

	for i := 0; i < input; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Factorial(input int) int {
	return recursive(1, input)
}

func recursive(hasil, index int) int {
	if index == 1 {
		return hasil
	} else {
		return recursive(hasil*index, index-1)
	}
}

func main() {
	fmt.Println("== Output Answer ==")
	fmt.Println("== Palindrom ==")
	fmt.Println(IsPalindrome("radar")) // True
	fmt.Println(IsPalindrome("hello")) // False
	fmt.Println("== Find value max ==")
	array := [5]int{3, 5, 1, 9, 2}
	maxValue := FindMaxValue(array)
	fmt.Println("Max:", maxValue) // 9
	fmt.Println("== Karakter Star ==")
	MakeStar(5)
	fmt.Println("== Faktorial ==")
	resultFactorial := Factorial(5)
	fmt.Println(resultFactorial) // karena 5! = 5 x 4 x 3 x 2 x 1 = 120
}
