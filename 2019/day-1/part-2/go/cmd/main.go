package main

import (
	"day-one-part-two/internal/algo"
	"day-one-part-two/internal/input"
	"flag"
	"fmt"
	"log"
	"os"
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
	linesAsInts, err := input.LinesToSliceOfInts(lines)
	if err != nil {
		log.Fatal(err)
	}
	result := algo.CalculateTotalFuelNeeded(linesAsInts)
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
