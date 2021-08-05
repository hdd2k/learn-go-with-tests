package main

func Sum(nums []int) int {
	summed := 0
	for _, n := range nums {
		summed += n
	}
	return summed
}

func SumAll(arrs ...[]int) (sums []int) {
	var arrResults []int
	for _, numbers := range arrs {
		arrResults = append(arrResults, Sum(numbers))
	}
	return arrResults
}

func SumAllTail(arrs ...[]int) (sumTails []int) {
	var arrResults []int
	for _, numbers := range arrs {
		numsLen := len(numbers)
		var tailNums []int
		if numsLen == 0 {
			tailNums = numbers
		} else {
			tailNums = numbers[1:]
		}
		arrResults = append(arrResults, Sum(tailNums))
	}
	return arrResults
}
