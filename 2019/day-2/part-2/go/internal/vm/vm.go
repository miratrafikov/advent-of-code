package vm

func Run(program []int, seeds []int) int {
	program[1], program[2] = seeds[0], seeds[1]
	for i := 0; i < len(program); i += 4 {
		switch program[i] {
		case 1:
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
			break
		case 2:
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
			break
		case 99:
			return program[0]
		}
	}
	return program[0]
}
