package main

func Sum(nums []int) int {
	summed := 0
	for _, n := range nums {
		summed += n
	}
	return summed
}
