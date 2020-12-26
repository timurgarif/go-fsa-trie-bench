package main

import (
	"bench"
	"fmt"
)

func main() {
	bench.PrepareEnLemPosSample()
	fmt.Println("Mem stat before build")
	bench.PrintMemUsage()

	t := bench.BuildArmonRadixTrie()
	fmt.Println("Mem stat after build")
	bench.PrintMemUsage()
	_, _ = t.Get("prevent")
}
