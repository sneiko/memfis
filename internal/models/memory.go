package models

import "fmt"

// MemoryStats represents runtime.MemStats data
type MemoryStats struct {
	Alloc         uint64    `json:"alloc"`
	TotalAlloc    uint64    `json:"totalAlloc"`
	Sys           uint64    `json:"sys"`
	Lookups       uint64    `json:"lookups"`
	Mallocs       uint64    `json:"mallocs"`
	Frees         uint64    `json:"frees"`
	HeapAlloc     uint64    `json:"heapAlloc"`
	HeapSys       uint64    `json:"heapSys"`
	HeapIdle      uint64    `json:"heapIdle"`
	HeapInuse     uint64    `json:"heapInuse"`
	HeapReleased  uint64    `json:"heapReleased"`
	HeapObjects   uint64    `json:"heapObjects"`
	Stack         string    `json:"stack"`
	MSpan         string    `json:"mspan"`
	MCache        string    `json:"mcache"`
	BuckHashSys   uint64    `json:"buckHashSys"`
	GCSys         uint64    `json:"gcSys"`
	OtherSys      uint64    `json:"otherSys"`
	NextGC        uint64    `json:"nextGC"`
	LastGC        uint64    `json:"lastGC"`
	PauseNs       []uint64  `json:"pauseNs"`
	PauseEnd      []uint64  `json:"pauseEnd"`
	NumGC         uint32    `json:"numGC"`
	NumForcedGC   uint32    `json:"numForcedGC"`
	GCCPUFraction float64   `json:"gcCPUFraction"`
	DebugGC       bool      `json:"debugGC"`
	MaxRSS        uint64    `json:"maxRSS"`
}

// StackTrace represents a memory allocation stack trace
type StackTrace struct {
	ID          int      `json:"id"`
	AllocBytes  int      `json:"allocBytes"`
	AllocCount  int      `json:"allocCount"`
	InUseBytes  int      `json:"inUseBytes"`
	InUseCount  int      `json:"inUseCount"`
	Addresses   []string `json:"addresses"`
	Functions   []string `json:"functions"`
}

// MemoryData contains all parsed memory profiling data
type MemoryData struct {
	MemStats    MemoryStats  `json:"memStats"`
	StackTraces []StackTrace `json:"stackTraces"`
	FileName    string       `json:"fileName"`
}

// FormatBytes converts bytes to human readable format
func FormatBytes(bytes uint64) string {
	if bytes == 0 {
		return "0 B"
	}
	const unit = 1024
	sizes := []string{"B", "KB", "MB", "GB", "TB"}
	
	i := 0
	size := float64(bytes)
	for size >= unit && i < len(sizes)-1 {
		size /= unit
		i++
	}
	
	if i == 0 {
		return fmt.Sprintf("%.0f %s", size, sizes[i])
	}
	return fmt.Sprintf("%.2f %s", size, sizes[i])
}

// FormatNumber adds thousand separators to numbers
func FormatNumber(num uint64) string {
	str := fmt.Sprintf("%d", num)
	n := len(str)
	if n <= 3 {
		return str
	}
	
	result := ""
	for i, char := range str {
		if i > 0 && (n-i)%3 == 0 {
			result += ","
		}
		result += string(char)
	}
	return result
}

