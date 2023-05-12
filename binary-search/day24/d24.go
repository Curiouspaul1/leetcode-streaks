package main

import "fmt"

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}

	ans := searchMatrix(matrix, 3)
	fmt.Print(ans)

}

func searchMatrix(matrix [][]int, target int) bool {

	Top := len(matrix[0])
	Bottom := 0

	for Bottom <= Top {
		mid := (Top + Bottom) / 2
		finalAnswer := search(matrix[mid], target)
		

	}
	return false
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left <= right {
		mid := ((left + right) / 2) + 1
		//fmt.Print(left, right, mid)
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
