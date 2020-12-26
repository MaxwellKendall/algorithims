package main

import (
	"strings"
	"fmt"
)

// IsUnique takes a string and returns a boolean
func IsUnique(s string) bool {
	arr := strings.Split(s, "")
	for i, v := range arr {
		for _, v2 := range arr[:i] {
			if v2 == v {
				fmt.Println(v)
				return false
			}
		}
	}
	return true
}

// IsPermutation tells us if a string is a substring of another
func IsPermutation(s1 string, s2 string) bool {
	if (strings.Contains(s2, s1) || strings.Contains(s1, s2)) {
		return true
	}
	return false
}

// URLify returns a string with urlEncoded spaces
func URLify(s string) string {
	arr := strings.Split(s, "")

	rtrn := ""

	for _, v := range arr {
		if (v == " ") {
			rtrn = rtrn + "%20"
		} else {
			rtrn = rtrn + v
		}
	}

	return rtrn
}

// PalindromePermutation returns true if string is palindrome or has potential for palindrome status
func PalindromePermutation(s string) bool {
	var occurrencesByChar = make(map[string]int)
	arr := strings.Split(s, "")
	for _, v := range arr {
		if (v != " ") {
			occurrencesByChar[v]++
		}
	}

	for _, v := range occurrencesByChar {
		if (v % 2 != 0 && v != 1) {
			return false
		}	
	}

	return true
}

// OneAway returns true if input is one char edit (delete/replace/add) away from equivalency
func OneAway(s1 string, s2 string) bool {
	len1 := len(s1)
	len2 := len(s2)
	if (len1 + 1 == len2 || len1 - 1 == len2 || len1 == len2) {
		var occurrencesByChar = make(map[string]int)
		allStrs := append(strings.Split(s1, ""), strings.Split(s2, "")...)
		for _, v := range allStrs {
			occurrencesByChar[v]++
		}
		charsNotEven := 0
		for _, v := range occurrencesByChar {
			if (v % 2 != 0) {
				charsNotEven++
			}
		}

		if (len1 == len2) {
			return charsNotEven <= 2
		}

		return charsNotEven <= 1
	}
	return false
}

func main() {
	fmt.Println("IsUnique: ", IsUnique("123451"))
	fmt.Println("IsPermutation: ", IsPermutation("hell", "hellz"))
	fmt.Println("URLify: ", URLify("google.com/ hello world"))
	fmt.Println("PalindromePermutation: ", PalindromePermutation("a plan a canal man a panama"))
	fmt.Println("OneAway: ", OneAway("pale", "bale"))
}
