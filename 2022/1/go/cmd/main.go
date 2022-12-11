package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := getFilePathFromArgs()
	if filePath == "" {
		log.Fatal("input file path must be provided")
	}

	lines, err := readLinesFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	result := largestSumOfConsecutiveNumbers(lines)
	fmt.Println(result)
}

func getFilePathFromArgs() string {
	flagHelp, flagFile := readFlags()
	if flagHelp {
		fmt.Println("path to input file provided using -f flag")
		fmt.Println("help printed using -h flag")
		os.Exit(0)
	}

	return flagFile
}

func readFlags() (bool, string) {
	flagHelp := flag.Bool("h", false, "show help")
	flagFile := flag.String("f", "", "path to file containing the problem's input")
	flag.Parse()

	return *flagHelp, *flagFile
}

func readLinesFromFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return []string{}, fmt.Errorf("could not read from file, %s", err)
	}

	return strings.Split(string(data), "\n"), nil
}

func largestSumOfConsecutiveNumbers(lines []string) int {
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
