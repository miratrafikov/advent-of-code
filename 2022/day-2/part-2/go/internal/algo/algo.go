package algo

var roundScores = map[string]int{
	"A X": 3,
	"A Y": 4,
	"A Z": 8,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 2,
	"C Y": 6,
	"C Z": 7,
}

func CalculateTotalResult(lines []string) int {
	total := 0
	for _, line := range lines {
		total += roundScores[line]
	}

	return total
}
