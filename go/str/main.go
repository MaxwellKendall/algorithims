package main

import (
	"strconv"
	"strings"
	"fmt"
)

// IsUnique takes a string and returns a boolean
func IsUnique(s string) bool {
	// complexity upper bound: O(n)
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
	// complexity upper bound: O(2n) or dropping constants to O(n)
	if (strings.Contains(s2, s1) || strings.Contains(s1, s2)) {
		return true
	}
	return false
}

// URLify returns a string with urlEncoded spaces
func URLify(s string) string {
	// complexity upper bound: O(n)
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
	// complexity upper bound: 0(2n) dropping constants --> O(n)
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
	// complexity upper bound: O(n)
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
	// complexity upper bound: O(n)
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

// RotateMatrix takes a matrix (2d array representing an image in pixels) and rotates it 90 degrees (❓ confused on expected return for this... ❓)
func RotateMatrix(m [][]int) [][]int {
	// complexity upper bound: O(n^2)
	rtrn := make([][]int, len(m))
	
	for row, v := range m {
		// O(n^2) unavoidable in this case?
		newRow := make([]int, len(m[row]))
		for i, v2 := range v {
			newVal := v2 + 90
			if (newVal > 360) {
				newVal = newVal - 360
			}
			newRow[i] = newVal
		}
		rtrn[row] = newRow
	}

	return rtrn
}

// ZeroMatrix takes a matrix (2d array) and returns a matrix where any index in a row w/ a zero, along with that entire row, is zeroed out
func ZeroMatrix(m [][]int) [][]int {
	// complexity upper bound: O(n^2)
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

// StringRotation takes two strings and returns a boolean indicating whether inputs are contain the same substrings
func StringRotation(s1 string, s2 string) bool {
	// complexity upper bound: O(3n) --> dropping constants to --> O(n)
	if len(s1) != len(s2) {
		return false
	}
	// step 1: find substr in s2 containing first char of s1
	firstChar := string(s1[0])
	// needs to be 2d array where at the end we select the longest item
	subStrIndices := make([]int, 0)
	charsAfterFirstChar := 0

	for i, v := range s2 {
		if string(v) == firstChar {
			subStrIndices = append(subStrIndices, i)
			charsAfterFirstChar++
		} else if charsAfterFirstChar >= 1 {
			nextChar := string(s1[charsAfterFirstChar])
			if string(v) == nextChar {
				subStrIndices = append(subStrIndices, i)
				if i == len(s2) - 1 {
					charsAfterFirstChar = len(s2) - 1
				} else {
					charsAfterFirstChar++
				}
			}
		}
	}

	if len(subStrIndices) <= 1 {
		return false
	}

	newS2 := ""

	for i, v := range s2 {
		if i == subStrIndices[0] || i == subStrIndices[1] {
			continue
		} else {
			newS2 = newS2 + string(v)
		}
	}

	fmt.Println(subStrIndices)
	// fmt.Println(s1[0:charsAfterFirstChar], newS2)

	return IsPermutation(s1[:subStrIndices[1]], newS2)
}

func main() {
	// matrixInput := [][]int{
	// 	{5, 1, 2, 0},
	// 	{4, 5, 6, 7},
	// }
	// fmt.Println("IsUnique: ", IsUnique("123451"))
	// fmt.Println("IsPermutation: ", IsPermutation("hell", "hellz"))
	// fmt.Println("URLify: ", URLify("google.com/ hello world"))
	// fmt.Println("PalindromePermutation: ", PalindromePermutation("a plan a canal man a panama"))
	// fmt.Println("OneAway: ", OneAway("pale", "bale"))
	// fmt.Println("CompressString: ", CompressString("mississippi"))
	// fmt.Println("RotateMatrix: ", RotateMatrix(matrixInput))
	// fmt.Println("ZeroMatrix: ", ZeroMatrix(matrixInput))
	fmt.Println("StringRotation: ", StringRotation("basketball", "babasktell"))
}
