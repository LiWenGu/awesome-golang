package main

import (
	"fmt"
	"publicChain/part1-Basic-prototype/BLC"
)


func main() {
	block := BLC.NewBlock("Genenis Block", 1, []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	block.SetHash()
	fmt.Println(block)
}
