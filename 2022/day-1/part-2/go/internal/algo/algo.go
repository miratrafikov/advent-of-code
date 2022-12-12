package algo

import (
	"sort"
	"strconv"
)

func SumOfThreeLargestSumsOfConsecutiveNumbers(lines []string) int {
	maxSums := []int{0, 0, 0}
	currentSum := 0
	for _, line := range lines {
		if line != "" {
			number, _ := strconv.Atoi(line)
			currentSum += number
		} else {
			if currentSum > maxSums[0] {
				pushNewMaxSum(currentSum, &maxSums)
			}

			currentSum = 0
		}
	}

	return reduceToSum(maxSums)
}

func pushNewMaxSum(num int, nums *[]int) {
	concatenatedNumbers := append(*nums, num)
	sort.Ints(concatenatedNumbers)
	concatenatedNumbers = concatenatedNumbers[1:4]
	*nums = concatenatedNumbers
}

func reduceToSum(nums []int) int {
	sum := 0
	for _, v := range(nums) {
		sum += v
	}

	return sum
}