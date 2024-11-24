package main

import (
	"fmt"
	twosums "leetcode-go/solutions/two-sums"
)

// test_assertion(t, []int{2, 7, 11, 15}, 9, []int{0, 1})
// test_assertion(t, []int{3, 2, 4}, 6, []int{1, 2})
// test_assertion(t, []int{3, 3}, 7, []int{0, 1})
func main() {
	a := []int{2, 7, 11, 15}
	v := twosums.TwoSum(a, 9)
	fmt.Printf("%v, %v\n", a, v)

	a1 := []int{3, 2, 4}
	v1 := twosums.TwoSum(a1, 6)
	fmt.Printf("%v, %v\n", a1, v1)

	a2 := []int{3, 3}
	v2 := twosums.TwoSum(a2, 6)
	fmt.Printf("%v, %v\n", a2, v2)
}
