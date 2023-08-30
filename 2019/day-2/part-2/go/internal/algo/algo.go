package algo

import "day-two-part-two/internal/vm"

func Run(program []int) int {
	target := 19690720
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			seeds := []int{i, j}
			result := vm.Run(copyProgram(program), seeds)
			if result == target {
				return 100*seeds[0] + seeds[1]
			}
		}
	}
	panic("bad program or target")
}

func copyProgram(program []int) []int {
	programCopy := make([]int, len(program))
	copy(programCopy, program)
	return programCopy
}
