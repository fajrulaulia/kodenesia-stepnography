package stepnography

////disable it if not run using go run
//package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decimalToBinary(nilai int64) int {
	result, err := strconv.Atoi(strconv.FormatInt(int64(nilai), 2)[len(strconv.FormatInt(int64(nilai), 2))-1:])
	if err != nil {
		fmt.Println(result)
	}
	return result
}

func binertoChar(input string) string {
	b := make([]byte, 0)
	for _, s := range strings.Fields(input) {
		n, _ := strconv.ParseUint(s, 2, 8)
		b = append(b, byte(n))
	}
	return string(b)
}

func generateStepnography(input []int) string {
	var gen []int //intialize to write result from DecimalToBinary()
	for _, decimal := range input {
		n := decimalToBinary(int64(decimal))
		gen = append(gen, n) //append gen from func DecimalToBinary()
	}
	return binertoChar(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(gen)), ""), "[]")) //result character
}

////disable it if not run using go run
// func main() {
// 	result := generateStepnography([]int{1, 1, 0, 0, 0, 0, 1})
// 	// result := generateStepnography([]int{31, 31, 20, 20, 20, 21, 8})
// 	// result := generateStepnography([]int{51, 61, 10, 0, 0, 15, 17})
// 	log.Println("result", result)
// }
