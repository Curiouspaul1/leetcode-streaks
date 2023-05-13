package main

import (
	"fmt"
	"math"
)

func main() {
	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums2 := []int{3,4,5,1,2}
	fmt.Print(findMin(nums2))

}

func findMin(nums []int) int {
	left, right, min := 0, len(nums)-1, nums[0]
	for left < right {
		mid := (left + right) / 2
		if (nums[left] > nums[right]) && (nums[mid] > nums[left]) && (nums[mid] > nums[right]) {
			left = mid + 1
			curMin := math.Min(float64(min), float64(nums[left]))
			min = int(curMin)
			return min
			

			//fmt.Printf("Rotated: and min is in the second half %v \n", min)
		} else {
			right = mid - 1
			curMin := math.Min(float64(min), float64(nums[left]))
			min = int(curMin)
			return min
			//fmt.Printf("Not Rotated %v", min)
		}
	}
	return min
}
