package algo

import "strconv"

func LargestSumOfConsecutiveNumbers(lines []string) int {
	maxSum := 0
	currentSum := 0
	for _, line := range lines {
		if line == "" {
			currentSum = 0
		} else {
			number, _ := strconv.Atoi(line)
			currentSum += number
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}

	return maxSum
}
