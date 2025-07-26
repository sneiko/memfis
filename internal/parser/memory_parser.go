package parser

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"memfis/internal/models"
)

// ParseMemoryData parses memory profiling data from a file
func ParseMemoryData(filename string) (*models.MemoryData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var stackTraces []models.StackTrace
	var memStats models.MemoryStats

	stackID := 0
	inMemStats := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "# runtime.MemStats") {
			inMemStats = true
			continue
		}

		if inMemStats {
			parseMemStatsLine(line, &memStats)
			continue
		}

		// Parse stack traces
		if strings.Contains(line, ":") && strings.Contains(line, "[") && strings.Contains(line, "]") && strings.Contains(line, "@") {
			stackTrace := parseStackTraceLine(line, stackID)
			if stackTrace != nil {
				// Read function lines
				for scanner.Scan() {
					funcLine := strings.TrimSpace(scanner.Text())
					if strings.HasPrefix(funcLine, "#\t") {
						funcInfo := strings.TrimPrefix(funcLine, "#\t")
						stackTrace.Functions = append(stackTrace.Functions, funcInfo)
					} else {
						break
					}
				}
				stackTraces = append(stackTraces, *stackTrace)
				stackID++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &models.MemoryData{
		MemStats:    memStats,
		StackTraces: stackTraces,
		FileName:    filename,
	}, nil
}

func parseMemStatsLine(line string, memStats *models.MemoryStats) {
	switch {
	case strings.HasPrefix(line, "# Alloc = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# Alloc = "), 10, 64)
		memStats.Alloc = val
	case strings.HasPrefix(line, "# TotalAlloc = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# TotalAlloc = "), 10, 64)
		memStats.TotalAlloc = val
	case strings.HasPrefix(line, "# Sys = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# Sys = "), 10, 64)
		memStats.Sys = val
	case strings.HasPrefix(line, "# Lookups = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# Lookups = "), 10, 64)
		memStats.Lookups = val
	case strings.HasPrefix(line, "# Mallocs = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# Mallocs = "), 10, 64)
		memStats.Mallocs = val
	case strings.HasPrefix(line, "# Frees = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# Frees = "), 10, 64)
		memStats.Frees = val
	case strings.HasPrefix(line, "# HeapAlloc = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# HeapAlloc = "), 10, 64)
		memStats.HeapAlloc = val
	case strings.HasPrefix(line, "# HeapSys = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# HeapSys = "), 10, 64)
		memStats.HeapSys = val
	case strings.HasPrefix(line, "# HeapIdle = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# HeapIdle = "), 10, 64)
		memStats.HeapIdle = val
	case strings.HasPrefix(line, "# HeapInuse = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# HeapInuse = "), 10, 64)
		memStats.HeapInuse = val
	case strings.HasPrefix(line, "# HeapReleased = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# HeapReleased = "), 10, 64)
		memStats.HeapReleased = val
	case strings.HasPrefix(line, "# HeapObjects = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# HeapObjects = "), 10, 64)
		memStats.HeapObjects = val
	case strings.HasPrefix(line, "# Stack = "):
		memStats.Stack = strings.TrimPrefix(line, "# Stack = ")
	case strings.HasPrefix(line, "# MSpan = "):
		memStats.MSpan = strings.TrimPrefix(line, "# MSpan = ")
	case strings.HasPrefix(line, "# MCache = "):
		memStats.MCache = strings.TrimPrefix(line, "# MCache = ")
	case strings.HasPrefix(line, "# BuckHashSys = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# BuckHashSys = "), 10, 64)
		memStats.BuckHashSys = val
	case strings.HasPrefix(line, "# GCSys = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# GCSys = "), 10, 64)
		memStats.GCSys = val
	case strings.HasPrefix(line, "# OtherSys = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# OtherSys = "), 10, 64)
		memStats.OtherSys = val
	case strings.HasPrefix(line, "# NextGC = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# NextGC = "), 10, 64)
		memStats.NextGC = val
	case strings.HasPrefix(line, "# LastGC = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# LastGC = "), 10, 64)
		memStats.LastGC = val
	case strings.HasPrefix(line, "# NumGC = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# NumGC = "), 10, 32)
		memStats.NumGC = uint32(val)
	case strings.HasPrefix(line, "# NumForcedGC = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# NumForcedGC = "), 10, 32)
		memStats.NumForcedGC = uint32(val)
	case strings.HasPrefix(line, "# GCCPUFraction = "):
		val, _ := strconv.ParseFloat(strings.TrimPrefix(line, "# GCCPUFraction = "), 64)
		memStats.GCCPUFraction = val
	case strings.HasPrefix(line, "# DebugGC = "):
		val, _ := strconv.ParseBool(strings.TrimPrefix(line, "# DebugGC = "))
		memStats.DebugGC = val
	case strings.HasPrefix(line, "# MaxRSS = "):
		val, _ := strconv.ParseUint(strings.TrimPrefix(line, "# MaxRSS = "), 10, 64)
		memStats.MaxRSS = val
	case strings.HasPrefix(line, "# PauseNs = "):
		pauseStr := strings.TrimPrefix(line, "# PauseNs = ")
		pauseStr = strings.Trim(pauseStr, "[]")
		pauseValues := strings.Fields(pauseStr)
		for _, pv := range pauseValues {
			if val, err := strconv.ParseUint(pv, 10, 64); err == nil {
				memStats.PauseNs = append(memStats.PauseNs, val)
			}
		}
	case strings.HasPrefix(line, "# PauseEnd = "):
		pauseStr := strings.TrimPrefix(line, "# PauseEnd = ")
		pauseStr = strings.Trim(pauseStr, "[]")
		pauseValues := strings.Fields(pauseStr)
		for _, pv := range pauseValues {
			if val, err := strconv.ParseUint(pv, 10, 64); err == nil {
				memStats.PauseEnd = append(memStats.PauseEnd, val)
			}
		}
	}
}

func parseStackTraceLine(line string, stackID int) *models.StackTrace {
	re := regexp.MustCompile(`(\d+): (\d+) \[(\d+): (\d+)\] @ (.+)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) != 6 {
		return nil
	}

	inUseBytes, _ := strconv.Atoi(matches[2])
	allocCount, _ := strconv.Atoi(matches[3])
	allocBytes, _ := strconv.Atoi(matches[4])
	addresses := strings.Fields(matches[5])

	return &models.StackTrace{
		ID:         stackID,
		AllocBytes: allocBytes,
		AllocCount: allocCount,
		InUseBytes: inUseBytes,
		InUseCount: 1,
		Addresses:  addresses,
		Functions:  []string{},
	}
}
