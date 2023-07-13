package binarysearch

func findPivot(nums []int) (int, int) {
	l, r, res, minIdx := 0, len(nums)-1, nums[0], 0

	for l <= r {
		if nums[l] < nums[r] {
			if nums[l] < res {
				return nums[l], l
			} else {
				return res, minIdx
			}
		}
		m := (l + r) / 2

		if nums[m] < res {
			minIdx = m
			res = nums[m]
		}
		if nums[m] >= nums[l] {
			// Search right
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return res, minIdx
}

func search(nums []int, target int) int {
	// Find pivot
	_, pivIdx := findPivot(nums)

	// Construct a corrected slice
	sortedSlice := append(nums[pivIdx:], nums[:pivIdx]...)

	// Binary search in sorted slice
	l, r := 0, len(sortedSlice)
	for l < r {
		m := (l + r) / 2
		if (m == l || m == r) && sortedSlice[m] != target {
			return -1
		}
		if sortedSlice[m] == target {
			// m corresponds to which idx in original nums?
			// add pivIdx, and modulus with len(nums)
			return (m + pivIdx) % len(nums)
		}
		if sortedSlice[m] < target {
			// Search right
			l = m
		} else {
			r = m
		}
	}
	return -1
}

func FindTarget(nums []int, target int) int {
	return search(nums, target)
}

func FindMin(nums []int) int {
	min, _ := findPivot(nums)
	return min
}
