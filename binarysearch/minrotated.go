package binarysearch

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func find(nums []int) int {
	l, r, res := 0, len(nums)-1, nums[0]

	for l <= r {
		if nums[l] < nums[r] {
			return min(nums[l], res)
		}
		var m int = (l + r) / 2
		res = min(nums[m], res)
		if nums[m] >= nums[l] {
			// Search right
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return res
}

func FindMin(nums []int) int {
	return find(nums)
}
