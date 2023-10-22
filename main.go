package main

import "fmt"

func main() {
	/*
		 cpu := NewCPU()
		// cpu := NewCPU(make([]uint16, 65535))

		regNames := []string{"ip", "acc", "r1", "r2", "r3", "r4", "r5", "r6", "r7", "r8"}

		for i, n := range regNames {
			// i *= 2
			cpu.setRegister(n, uint16(i))
			reg, _ := cpu.getRegister(n)
			fmt.Printf("cpu.getRegister(%v): %v\n", n, reg)
		}

		fmt.Printf("cpu: %+v\n", cpu)
		cpu.fetch16()
		cpu.fetch16()
		// cpu.execute(0x11)
		fmt.Printf("cpu: %+v\n", cpu)
	*/

	memory := CreateMemory(20) //256
	memory[0] = MOV_LIT_R1
	memory[1] = 0x1234
	memory[2] = MOV_LIT_R2
	memory[3] = 0xABCD
	memory[4] = ADD_REG_REG
	memory[5] = 2
	memory[6] = 3
	// fmt.Printf("memory: %v\n", memory)
	// fmt.Printf("memory: %0x\n", memory)

	cpu := NewCPU(memory)
	cpu.step()
	cpu.step()
	cpu.step()
	fmt.Printf("cpu: %+v\n", cpu)
	// fmt.Printf("cpu: %#v\n", cpu)
}
