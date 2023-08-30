package input

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(path string) ([]int, error) {
	line, err := readFile(path)
	if err != nil {
		return []int{}, err
	}
	lineAsIntSlice, err := lineToSliceOfInts(line)
	if err != nil {
		return []int{}, err
	}
	return lineAsIntSlice, err
}

func readFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("could not read from file, %s", err)
	}
	return string(data), nil
}

func lineToSliceOfInts(line string) ([]int, error) {
	var err error
	trimmedLine := strings.TrimSuffix(line, "\n")
	strings := strings.Split(trimmedLine, ",")
	ints := make([]int, len(strings))
	for i := 0; i < len(strings); i++ {
		ints[i], err = strconv.Atoi(strings[i])
		if err != nil {
			return []int{}, err
		}
	}
	return ints, nil
}
