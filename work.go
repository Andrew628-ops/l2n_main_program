package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello Welcome")
}

func params(n int) bool {
	if n == 2 {
		return true
	} else if n >= 2 {
		fmt.Println("Too many arguement")
		return false
	} else {
		fmt.Println("Missing file")
		return false
	}
}

/*
func hexDecimal(hexStr string) string {
	value, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(value)
	return hexStr
}*/

func hexaDecimal(hexStr string) (int64, error) {
	value, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func binary(binStr string) (int64, error) {
	value, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func lower(s string) string {
	words := []rune(s)
	for char := range words {
		if words[char] >= 'A' && words[char] <= 'Z' {
			words[char] = words[char] + 32
		}
	}
	return string(words)
}

func upper(s string) string {
	words := []rune(s)
	for char := range words {
		if words[char] >= 'a' && words[char] <= 'z' {
			words[char] -= 32
		}
	}
	return string(words)
}

func punctuation(s string) bool {
	return s == "." || s == ":" || s == ";" || s == "," || s == "?"
}
