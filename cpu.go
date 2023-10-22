package main

import (
	"errors"
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

	// cpu.registers = make([]uint16, len(cpu.registerNames))
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
	// instructionIndex, _ := c.getRegister("ip")
	// instruction := c.memory[instructionIndex]
	// c.setRegister("ip", instructionIndex+1)
	// return instruction
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
		// Add register to the register
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
}

func (c *CPU) step() {
	instruction := (*c).fetch()
	(*c).execute(instruction)
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

// type CPU struct {
// 	memory        memory
// 	registerNames []string
// 	registers     []uint16
// 	registerMap   map[string]int
// 	// registers     []byte
// }

// func main() {
// 	cpu := NewCPU(make([]uint16, 20))

// 	regNames := []string{"ip", "acc", "r1", "r2", "r3", "r4", "r5", "r6", "r7", "r8"}

// 	for i, n := range regNames {
// 		i *= 32
// 		cpu.setRegister(n, uint16(i))
// 		reg, _ := cpu.getRegister(n)
// 		fmt.Printf("cpu.getRegister(%v): %v\n", n, reg)
// 	}
// 	ac, _ := cpu.getRegister("acc")
// 	fmt.Printf("ac: %v\n", ac)

// 	fmt.Printf("cpu: %+v\n", cpu)
// 	cpu.fetch16()
// 	cpu.execute(0x11)
// 	fmt.Printf("cpu: %+v\n", cpu)
// }

// func NewCPU(memory memory) CPU {
// 	cpu := CPU{
// 		memory:        memory,
// 		registerNames: []string{"ip", "acc", "r1", "r2", "r3", "r4", "r5", "r6", "r7", "r8"},
// 		registerMap:   make(map[string]int),
// 	}

// 	cpu.registers = CreateMemory(len(cpu.registerNames) * 2)

// 	for i, name := range cpu.registerNames {
// 		cpu.registerMap[name] = i * 2
// 	}
// 	return cpu
// }

// func (c CPU) getRegister(name string) (uint16, error) {
// 	var registerAsUint16 uint16

// 	for _, regName := range c.registerNames {
// 		if regName == name {
// 			registerAsUint16 = uint16(c.registers[c.registerMap[name]] | c.registers[c.registerMap[name]+1])
// 			// fmt.Printf("registerAsUint16: %v\n", registerAsUint16)
// 			return registerAsUint16, nil
// 		}
// 	}
// 	return 0, errors.New("register name not found")
// }

// func (c *CPU) setRegister(name string, value uint16) error /* uint16 */ {

// 	for _, regName := range (*c).registerNames {
// 		if regName == name {
// 			(*c).registers[uint16(c.registerMap[name])|uint16(c.registers[c.registerMap[name]+1])] = value
// 			return nil
// 		}
// 	}
// 	return errors.New("register name not found")
// 	// return registerAsUint
// }

// func (c *CPU) fetch() uint8 {
// 	instructionIndex, _ := c.getRegister("ip")

// 	instruction := uint8(c.memory[instructionIndex])
// 	(*c).setRegister("ip", instructionIndex+1)
// 	return instruction
// }

// func (c *CPU) fetch16() uint16 {
// 	instructionIndex, _ := c.getRegister("ip")

// 	// instruction := uint16(c.memory[instructionIndex])
// 	// instruction := uint16(c.memory[instructionIndex]<<8 | c.memory[instructionIndex] + 1)
// 	instruction := uint16(c.memory[instructionIndex] | c.memory[instructionIndex] + 1)
// 	(*c).setRegister("ip", instructionIndex+2)
// 	return instruction
// }

// func (c *CPU) execute(instruction uint16) {
// 	switch instruction {
// 	// Move literal into r1 register
// 	case 0x10:
// 		{
// 			literal := (*c).fetch16()
// 			(*c).setRegister("r1", literal)
// 		}
// 	// Move literal into r2 register
// 	case 0x11:
// 		{
// 			literal := (*c).fetch16()
// 			(*c).setRegister("r2", literal)
// 		}
// 	// Add register to the register
// 	case 0x12:
// 		{
// 			r1 := (*c).fetch()
// 			r2 := (*c).fetch()
// 			registerValue1 := uint16((*c).registers[r1*2]<<8 | (*c).registers[r1*2] + 1)
// 			registerValue2 := uint16((*c).registers[r2*2]<<8 | (*c).registers[r2*2] + 1)

// 		}
// 	}
// }

/*

func (c *CPU) setRegister(name string, value uint16) {
    for _, regName := range c.registerNames {
        if regName == name {
            // using big-endian byte order
            c.registers[c.registerMap[name]] = byte(value >> 8)  // high byte
            c.registers[c.registerMap[name]+1] = byte(value)     // low byte
            break
        }
    }
}

func (c *CPU) getRegister(name string) uint16 {
    var registerAsUint uint16
    for _, regName := range c.registerNames {
        if regName == name {
            // using big-endian byte order
            registerAsUint = uint16(c.registers[c.registerMap[name]])<<8 | uint16(c.registers[c.registerMap[name]+1])
            break
        }
    }
    return registerAsUint
}*/
