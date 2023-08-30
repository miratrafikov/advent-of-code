package vm

func Run(program []int) int {
	program[1] = 12
	program[2] = 2
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
