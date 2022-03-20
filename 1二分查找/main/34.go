package main

func SearchRange(nums []int, target int) []int {
	res := getMid(nums, target)
	if res == -1 {
		// 注意return格式
		return []int{-1, -1}
	} else {
		low := res
		high := res
		i := res
		// for 语句如何写
		for i < len(nums) {
			if nums[i] == target {
				high = i
				i++
			} else {
				break
			}

		}
		j := res
		for j > -1 {
			if nums[j] == target {
				low = j
				j--
			} else {
				break
			}
		}
		return []int{low, high}
	}
}

// 函数可以写内部吗？
func getMid(nums []int, target int) int {
	var (
		left  int = 0
		right int = len(nums)
	)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}
