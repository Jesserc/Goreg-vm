package main

// type memory []byte
type memory []uint16

// func CreateMemory(length int) *memory {
// 	// see this function later
// 	mem := make(memory, length)
// 	return &mem
// }

func CreateMemory(length int) memory {
	return make(memory, length)
}
