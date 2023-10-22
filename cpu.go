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

func main() {
	cpu := NewCPU(make([]uint16, 10))
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
}

func NewCPU(memory []uint16) CPU {
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
	nextInstructionAddress, _ := c.getRegister("ip")
	instruction := c.memory[nextInstructionAddress]
	c.setRegister("ip", nextInstructionAddress+1)
	return instruction
}

func (c *CPU) fetch16() uint16 {
	// nextInstructionAddress, _ := c.getRegister("ip")
	// instruction := c.memory[nextInstructionAddress]
	// c.setRegister("ip", nextInstructionAddress+1)
	// return instruction
	return c.fetch()
}

func (c *CPU) execute(instruction uint16) {
	switch instruction {
	// Move literal into r1 register
	case 0x10:
		{
			literal := (*c).fetch16()
			(*c).setRegister("r1", literal)
		}
	// Move literal into r2 register
	case 0x11:
		{
			literal := (*c).fetch16()
			(*c).setRegister("r2", literal)
		}
		// Add register to the register
	case 0x12:
		{
			r1 := c.fetch()
			r2 := c.fetch()

			registerValue1 := c.registers[r1]
			registerValue2 := c.registers[r2]

			(*c).setRegister("acc", registerValue1+registerValue2)
			// Assuming you want to add registerValue2 to registerValue1 and store the result in register r1
			// newValue := registerValue1 +
		}
	}
}

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
// 	nextInstructionAddress, _ := c.getRegister("ip")

// 	instruction := uint8(c.memory[nextInstructionAddress])
// 	(*c).setRegister("ip", nextInstructionAddress+1)
// 	return instruction
// }

// func (c *CPU) fetch16() uint16 {
// 	nextInstructionAddress, _ := c.getRegister("ip")

// 	// instruction := uint16(c.memory[nextInstructionAddress])
// 	// instruction := uint16(c.memory[nextInstructionAddress]<<8 | c.memory[nextInstructionAddress] + 1)
// 	instruction := uint16(c.memory[nextInstructionAddress] | c.memory[nextInstructionAddress] + 1)
// 	(*c).setRegister("ip", nextInstructionAddress+2)
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
