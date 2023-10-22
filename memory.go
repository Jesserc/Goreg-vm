package main

// type Memory []byte
type Memory []uint16

// func CreateMemory(length int) *Memory {
// 	// see this function later
// 	mem := make(Memory, length)
// 	return &mem
// }

func CreateMemory(length int) Memory {
	return make(Memory, length)
}
