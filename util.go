package main

import (
	"log"
	"runtime"
	"time"
)

var (
	lastAlloc uint64
)

func PrintMemoryUsage(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	diff := m.Alloc - lastAlloc
	lastAlloc = m.Alloc

	log.Printf("%s - Diff: %v MiB, Alloc: %v MiB, TotalAlloc: %v MiB, Sys: %v MiB, NumGC: %v",
		label,
		bToMb(diff),
		bToMb(m.Alloc),
		bToMb(m.TotalAlloc),
		bToMb(m.Sys),
		m.NumGC,
	)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func PrintElapsedTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
