package main

import "testing"

func TestIsUnique(t *testing.T) {
	test := IsUnique("111111")
	if (test == true) {
		t.Error("isUnique is wrong", test)
	}
}

func TestIsPermutation(t *testing.T) {
	test := IsPermutation("he", "hello")
	if (test == false) {
		t.Error("IsPermutation is wrong", test)
	}
}

func TestURLify(t *testing.T) {
	test := URLify("confessionalchristianity.com/confessions/reformed baptist catechism")
	if (test != "confessionalchristianity.com/confessions/reformed%20baptist%20catechism") {
		t.Error("IsPermutation is wrong", test)
	}
}

func TestPalindromePermutation(t *testing.T) {
	test := PalindromePermutation("taco cat")
	if (test == false) {
		t.Error("PalindromePermutation is wrong", test)
	}
}

func TestOneAway(t *testing.T) {
	test := OneAway("heck", "deck")
	if (test != true) {
		t.Error("OneAway is wrong", test)
	}
}

func TestCompressString(t *testing.T) {
	test := CompressString("mississippi")
	if (test != "mis2is2ip2i") {
		t.Error("CompressString is wrong", test)
	}
}

func TestStringRotation(t *testing.T) {
	test := StringRotation("mississippi", "ppississiim")
	if (!test) {
		t.Error("StringRotation is wrong", test)
	}
}

func TestZeroMatrix(t *testing.T) {
	input := [][]int{
		{5, 1, 2, 0},
		{4, 5, 6, 7},
	}
	output := [][]int{
		{0, 0, 0, 0},
		{4, 5, 6, 0},
	}
	test := ZeroMatrix(input)
	for i, row := range test {
		for i2, v2 := range row {
			if v2 != output[i][i2] {
				t.Error("ZeroMatrix is wrong", test)
			}
		}
	}
}
