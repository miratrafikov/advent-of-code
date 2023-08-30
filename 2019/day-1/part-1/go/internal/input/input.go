package input

import "strconv"

func LinesToSliceOfInts(lines []string) ([]int, error) {
	s := make([]int, len(lines))
	for i, line := range lines {
		lineAsInt, err := strconv.Atoi(line)
		if err != nil {
			return s, err
		}
		s[i] = lineAsInt
	}
	return s, nil
}
