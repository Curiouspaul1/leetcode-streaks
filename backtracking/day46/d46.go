package main

import (
	"log"
)

func main()  {
	nums := []int{1,2,3}
	ans:=  subsets(nums)
	log.Println(ans)
}

func subsets(nums []int) [][]int {
	ps := [][]int{}
	ps = append(ps, []int{})
	for _, num := range nums {
		ss := make([][]int, len(ps), len(ps))
		copy(ss, ps)
		for i := range ss {
			ss[i] = append(ss[i], num)
		}
		for _,s := range ss {
			ps = append(ps, s)
			
		}		
	}
	return ps
}

