package main

import (
	"errors"
	"fmt"
)

type CPU struct {
	memory        []uint16
	registerNames []string
	registers     []uint16
	registerMap   map[string]int
}

func NewCPU(memory Memory) CPU {
	cpu := CPU{
		memory:        memory,
		registerNames: []string{"ip", "acc", "r1", "r2", "r3", "r4", "r5", "r6", "r7", "r8"},
		registerMap:   make(map[string]int),
	}

	cpu.registers = CreateMemory(len(cpu.registerNames))

	for i, name := range cpu.registerNames {
		cpu.registerMap[name] = i
	}
	return cpu
}

func (c CPU) getRegister(name string) (uint16, error) {
	for _, regName := range c.registerNames {
		if regName == name {
			return c.registers[c.registerMap[name]], nil
		}
	}
	return 0, errors.New("register name not found")
}

func (c *CPU) setRegister(name string, value uint16) error {
	for _, regName := range c.registerNames {
		if regName == name {
			c.registers[c.registerMap[name]] = value
			return nil
		}
	}
	return errors.New("register name not found")
}

func (c *CPU) fetch() uint16 {
	instructionIndex, _ := c.getRegister("ip")
	instruction := c.memory[instructionIndex]
	c.setRegister("ip", instructionIndex+1)
	return instruction
}

func (c *CPU) fetch16() uint16 {
	return c.fetch()
}

func (c *CPU) execute(instruction uint16) {
	switch instruction {
	// Move literal into r1 register
	case MOV_LIT_R1:
		{
			literal := (*c).fetch16()
			(*c).setRegister("r1", literal)
		}
	// Move literal into r2 register
	case MOV_LIT_R2:
		{
			literal := (*c).fetch16()
			(*c).setRegister("r2", literal)
		}
		// Add register to the register (we add values in r1 and r2 and save in acc register)
	case ADD_REG_REG:
		{
			r1 := (*c).fetch()
			r2 := (*c).fetch()

			registerValue1 := c.registers[r1]
			registerValue2 := c.registers[r2]

			(*c).setRegister("acc", registerValue1+registerValue2)
			return
		}
	}
}

func (c *CPU) step() {
	instruction := (*c).fetch()
	(*c).execute(instruction)
}

func (c *CPU) debug() {
	for _, name := range c.registerNames {
		registerValue, _ := c.getRegister(name)
		formattedRegisterValue := fmt.Sprintf("0x%04x", registerValue)
		fmt.Printf("%v: %v\n", name, formattedRegisterValue)
	}
}

/*
	EXECUTION STEP {
		we have 3 instructions at the time of writing this:
		MOV_LIT_R1  = 0x10
		MOV_LIT_R2  = 0x11
		ADD_REG_REG = 0x12,
		This will be the progression of the state,
		when we call step() 3 times in `main.go` to execute all 3 implemented instruction this time
		ip=5, mem[5]=2, acc=0 => ip=6, mem[6]=3, acc=0 => ip=7, mem[6]=3, acc=0x1234+0xabcd

			case ADD_REG_REG:
					{
						r1 := (*c).fetch() //2
						r2 := (*c).fetch() //3

						registerValue1 := c.registers[r1] // 0x1234
						registerValue2 := c.registers[r2] // 0xabcd

						(*c).setRegister("acc", registerValue1+registerValue2)
						return
					}
	}


*/
