package main

import (
	"bench"
	"fmt"
)

func main() {
	bench.PrepareEnLemPosSample()
	fmt.Println("Mem statbefore build")
	bench.PrintMemUsage()

	dawg := bench.BuildSmhanovDawgNoVal()
	fmt.Println("Mem stat after build")
	bench.PrintMemUsage()
	_ = dawg.IndexOf("prevent")
}
