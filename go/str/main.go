package main

import (
	"strconv"
	"strings"
	"fmt"
)

// IsUnique takes a string and returns a boolean
func IsUnique(s string) bool {
	for i, v := range s {
		for _, v2 := range s[:i] {
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

// CompressString takes a string and returns a string that represents repeated characters like: aa --> a2
func CompressString(s string) string {
	arr := strings.Split(s, "")
	rtrn := ""
	var prevVal string
	consecutiveChars := 1
	for i, v := range arr {
		if prevVal == v {
			consecutiveChars++
		} else if consecutiveChars > 1 {
			rtrn = rtrn + strconv.Itoa(consecutiveChars) + v
			consecutiveChars = 1
		} else {
			rtrn = rtrn + v
		}

		if i == len(arr) - 1 && prevVal == v {
			rtrn = rtrn + strconv.Itoa(consecutiveChars)
		}

		prevVal = v
	}
	return rtrn
}

func main() {
	fmt.Println("IsUnique: ", IsUnique("123451"))
	fmt.Println("IsPermutation: ", IsPermutation("hell", "hellz"))
	fmt.Println("URLify: ", URLify("google.com/ hello world"))
	fmt.Println("PalindromePermutation: ", PalindromePermutation("a plan a canal man a panama"))
	fmt.Println("OneAway: ", OneAway("pale", "bale"))
	fmt.Println("CompressString: ", CompressString("mississippi"))
}
