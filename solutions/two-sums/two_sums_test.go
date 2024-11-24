package twosums

import (
	"slices"
	"testing"
)

func TestFind(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	index, _ := Find(nums, 11)
	if index != 2 {
		t.Fatalf("incorrect")
	}
}

func TestIsSorted(t *testing.T) {
	n0 := []int{-1, 0}
	n1 := []int{1, 0}
	n2 := []int{1, 2, 3}
	n3 := []int{3, 2, 1}
	n4 := []int{0}
	if IsSorted(n0) != true {
		t.Fatalf("incorrected")
	}
	if IsSorted(n1) != false {
		t.Fatalf("incorrected")
	}
	if IsSorted(n2) != true {
		t.Fatalf("incorrected")
	}
	if IsSorted(n3) != false {
		t.Fatalf("incorrected")
	}
	if IsSorted(n4) != true {
		t.Fatalf("incorrected")
	}
}

func TestCase1(t *testing.T) {
	test_assertion(t, []int{2, 7, 11, 15}, 9, []int{0, 1})
}
func TestCase2(t *testing.T) {
	test_assertion(t, []int{3, 2, 4}, 6, []int{1, 2})
}
func TestCase3(t *testing.T) {
	test_assertion(t, []int{3, 3}, 6, []int{0, 1})
}

func test_assertion(t *testing.T, nums []int, target int, expected []int) {

	s := TwoSum(nums, target)
	eq := slices.Equal(s, expected)
	if !eq {
		t.Fatalf("incorrect")
	}
}
