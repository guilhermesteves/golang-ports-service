package core

import (
	"fmt"
	"math"
	"runtime"
	"testing"
)

func TestStreamJson(t *testing.T) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	printStats(m)
}

func printStats(m runtime.MemStats) {
	fmt.Printf("Alloc = %v MB\n", toMB(m.Alloc))
	fmt.Printf("TotalAlloc = %v MB", toMB(m.TotalAlloc))
	fmt.Printf("Sys = %v MB", toMB(m.Sys))
	fmt.Printf("NumGC = %v\n", m.NumGC)
}

func toMB(b uint64) uint64 {
	return b / uint64(math.Pow(1024, 2))
}
