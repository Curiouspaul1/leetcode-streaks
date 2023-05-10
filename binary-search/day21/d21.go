package main

import "fmt"

func main() {
	nums := []int{1, 0, 3, 5, 9, 12}
	ans := search(nums, 9)
	fmt.Printf("ans: %v", ans)
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left <= right {
		mid := ((left + right) / 2) + 1
		fmt.Printf("left: %v , right: %v , mid: %v \n", left, right, mid)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}

	}

	return -1

}
