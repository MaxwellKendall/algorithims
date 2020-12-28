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
