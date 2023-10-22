package main

import "fmt"

func main() {
	memory := CreateMemory(20) // or 256
	memory[0] = MOV_LIT_R1
	memory[1] = 0x1234
	memory[2] = MOV_LIT_R2
	memory[3] = 0xABCD
	memory[4] = ADD_REG_REG
	memory[5] = 2
	memory[6] = 3

	cpu := NewCPU(memory)
	fmt.Printf("Full cpu state before executions: %+v\n\n", cpu)

	cpu.debug()

	fmt.Println() // line space

	cpu.step()
	cpu.debug()

	fmt.Println() // line space

	cpu.step()
	cpu.debug()

	fmt.Println() // line space

	cpu.step()
	cpu.debug()

	fmt.Println() // line space

	fmt.Printf("Full cpu state after executions: %+v\n\n", cpu)

}
