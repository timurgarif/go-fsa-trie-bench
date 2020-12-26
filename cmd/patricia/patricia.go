package main

import (
	"bench"
	"fmt"
)

func main() {
	bench.PrepareEnLemPosSample()
	fmt.Println("Mem statbefore build")
	bench.PrintMemUsage()

	trie := bench.BuildPatricia()
	fmt.Println("Mem stat after build")
	bench.PrintMemUsage()
	_ = trie.Get([]byte("prevent"))
}
