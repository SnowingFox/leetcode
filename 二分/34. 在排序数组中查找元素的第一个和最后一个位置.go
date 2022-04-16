package main

import (
	"fmt"
)

func searchRange(nums []int, target int, flag bool) int {
	left, ans := 0, -1
	right := len(nums) - 1

	for left <= right {
		if mid := (left + right) / 2; nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			if flag {
				right = mid - 1
			} else {
				left = mid + 1
			}
			ans = mid
		}
	}

	return ans
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 4, 4, 4, 8}
	target := 7
	left := searchRange(arr, target, true)
	right := -1
	if left != -1 {
		right = searchRange(arr, target, false)
	}
	fmt.Println([]int{left, right})
}
