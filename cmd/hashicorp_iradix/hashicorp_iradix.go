package main

import (
	"bench"
	"fmt"
)

func main() {
	bench.PrepareEnLemPosSample()
	fmt.Println("Mem stat before build")
	bench.PrintMemUsage()

	t := bench.BuildHashicorpIradix()
	fmt.Println("Mem stat after build")
	bench.PrintMemUsage()
	_, _ = t.Get([]byte("prevent"))
}
