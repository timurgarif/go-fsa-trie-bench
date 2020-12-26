package bench

import (
	"fmt"
	"runtime"
)

func PrintMemUsage() {
	runtime.GC()
	b2kb := func(b uint64) uint64 {
		return b / 1024
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v Kb", b2kb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v Kb", b2kb(m.TotalAlloc))
	fmt.Printf("\tSys = %v Kb\n", b2kb(m.Sys))
}
