package dupes

func ContainsDuplicate(nums []int) bool {
	arrLen := len(nums)
	if arrLen < 1 && arrLen > 10000 {
		return true
	}
	for i := 0; i < arrLen; i++ {
		if nums[i] < -1000000000 || nums[i] > 1000000000 {
			return true
		}
	}

	set := map[int]bool{}
	for i := 0; i < arrLen; i++ {
		value := nums[i]
		set[value] = true
	}

	return len(nums) > len(set)
}
