package twosums

import (
	"errors"
	"math"
)

func withinThreshold(target int) bool {
	if target < -1000000000 || target > 1000000000 {
		return false
	}
	return true
}

func TwoSum(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen < 2 || numsLen > 10000 {
		return []int{}
	}

	if !withinThreshold(target) {

		return []int{}
	}

	seen := map[int]int{}

	for i, num := range nums {
		needs := target - num
		value, ok := seen[needs]
		// fmt.Printf("[%d] need %d, is %d, found? %t\n", i, needs, num, ok)
		if ok {
			return []int{value, i}
		}
		seen[num] = i
	}

	return []int{}
}

func Find(nums []int, target int) (int, error) {
	left := 0
	right := len(nums) - 1

	for ok := true; ok; ok = left <= right {
		middle := int(math.Floor((float64(left) + float64(right)) / 2.0))
		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			return middle, nil
		}
	}

	return 0, errors.New("could not find")
}

func IsSorted(nums []int) bool {

	if len(nums) <= 1 {
		return true
	}
	prev := math.MinInt32

	for _, num := range nums {
		if num >= prev {
			prev = num
			continue
		}
		return false
	}

	return true
}
