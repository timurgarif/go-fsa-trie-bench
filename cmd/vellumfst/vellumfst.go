package main

import (
	"bench"
	"fmt"
)

func main() {
	bench.PrepareEnLemPosSample()
	fmt.Println("Mem stat before build")
	bench.PrintMemUsage()

	fst := bench.BuildFstVellum()
	fmt.Println("Mem stat after build")
	bench.PrintMemUsage()
	_, _, _ = fst.Get([]byte("prevent"))
}
