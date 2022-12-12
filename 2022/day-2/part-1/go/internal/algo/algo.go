package algo

var roundScores = map[string]int{
	"A X": 4,
	"A Y": 8,
	"A Z": 3,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 7,
	"C Y": 2,
	"C Z": 6,
}

func CalculateTotalResult(lines []string) int {
	total := 0
	for _, line := range lines {
		total += roundScores[line]
	}

	return total
}
