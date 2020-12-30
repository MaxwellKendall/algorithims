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

// StringRotation takes two strings and returns a boolean indicating whether inputs are the same chars rotated in different orders
func StringRotation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	// first index is s1 & second is s2 
	occurrencesPerStringByChar := make(map[string][]int)
	for _, v := range strings.Split(s1, "") {
		if count, ok := occurrencesPerStringByChar[v]; ok {
			occurrencesPerStringByChar[v][0] = count[0] + 1
		} else {
			occurrencesPerStringByChar[v] = []int{1, 0}
		}
	}
	for _, v := range strings.Split(s2, "") {
		if count, ok := occurrencesPerStringByChar[v]; ok {
			occurrencesPerStringByChar[v][1] = count[1] + 1
		} else {
			// char only exists in s2
			return false
		}
	}

	for _, value := range occurrencesPerStringByChar {
		if value[0] != value[1] {
			return false
		}
	}

	return true
}

// ZeroMatrix takes a matrix (2d array) and returns a matrix where any index in a row w/ a zero, along with that entire row, is zeroed out
func ZeroMatrix(m [][]int) [][]int {
	rowAndIndexOfZero := make([]int, 0)
	
	for row, v := range m {
		// O(n^2) unavoidable in this case?
		for index, v2 := range v {
			if len(rowAndIndexOfZero) > 0 {
				break
			}
			if v2 == 0 {
				rowAndIndexOfZero = append(rowAndIndexOfZero, row)
				rowAndIndexOfZero = append(rowAndIndexOfZero, index)
			}
		}
	}

	if len(rowAndIndexOfZero) > 0 {
		allZeros := make([]int, len(m[0]))
		for i := range allZeros {
			allZeros[i] = 0
		}
		// Mutating fn input 👇 (not good)
		m[rowAndIndexOfZero[0]] = allZeros
		for row := range m {
			m[row][rowAndIndexOfZero[1]] = 0
		}
		return m
	}
	return m
}

func main() {
	matrixInput := [][]int{
		{5, 1, 2, 0},
		{4, 5, 6, 7},
	}
	fmt.Println("IsUnique: ", IsUnique("123451"))
	fmt.Println("IsPermutation: ", IsPermutation("hell", "hellz"))
	fmt.Println("URLify: ", URLify("google.com/ hello world"))
	fmt.Println("PalindromePermutation: ", PalindromePermutation("a plan a canal man a panama"))
	fmt.Println("OneAway: ", OneAway("pale", "bale"))
	fmt.Println("CompressString: ", CompressString("mississippi"))
	fmt.Println("StringRotation: ", StringRotation("waterbottle", "erbottlewat"))
	fmt.Println("ZeroMatrix: ", ZeroMatrix(matrixInput))
}
