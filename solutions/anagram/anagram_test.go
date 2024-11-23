package anagram

import "testing"

func TestCase1(ts *testing.T) {
	s := "anagram"
	t := "nagaram"

	testassertion(ts, s, t, true)
}
func TestCase2(ts *testing.T) {
	s := "car"
	t := "rat"

	testassertion(ts, s, t, false)
}

func testassertion(ts *testing.T, s string, t string, expected bool) {
	isAnagram := IsAnagram(s, t)

	if expected != isAnagram {
		ts.Fatalf("incorrect state")
	}
}
