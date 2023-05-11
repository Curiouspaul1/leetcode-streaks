package main

func sortedSquares(nums []int) []int {
	var result []int
	for _, num := range nums {
		result = append(result, num*num)
	}
	return result
}



