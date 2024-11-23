package dupes

import (
	"testing"
)

func TestCase1(t *testing.T) {

	nums := []int{1, 2, 3, 1}
	testassertion(t, nums, true)
}

func TestCase2(t *testing.T) {

	nums := []int{1, 2, 3, 4}
	testassertion(t, nums, false)

}

func TestCase3(t *testing.T) {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	testassertion(t, nums, true)
}

func testassertion(t *testing.T, nums []int, expect bool) {
	answer := ContainsDuplicate(nums)
	if answer != expect {
		t.Fatalf("incorrect")
	}
}
