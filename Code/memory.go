package model

// Memory
// Representation of the main memory, the unit is in kB
type Memory struct {
	Total     uint64 `json:"total"`    // "MemTotal" system installed RAM memory
	Free      uint64 `json:"free"`     // "MemFree" system unused RAM memory
	TotalSwap uint64 `json:"swap"`     // "SwapTotal"
	FreeSwap  uint64 `json:"freeSwap"` // "SwapFree"
}