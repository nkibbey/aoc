package intcode

import "fmt"

func ParseOpcode(program []int) {
	for instr := 0; instr < len(program) && program[instr] != 99; {
		fmt.Println(program[instr])
		instr += 4
	}

}

/*


 */
