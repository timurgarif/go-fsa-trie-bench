package main

import (
	"bench"
	"fmt"
)

func main() {
	bench.PrepareEnLemPosSample()
	fmt.Println("Mem statbefore build")
	bench.PrintMemUsage()

	m := bench.BuildMap()
	fmt.Println("Mem stat after build")
	bench.PrintMemUsage()
	_ = m["prevent"]
}
