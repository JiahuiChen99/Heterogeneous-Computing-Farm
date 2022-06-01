package model

// Core
// Representation of a CPU core
type Core struct {
	Processor uint64 `json:"processor"` // "processor" field
	CoreID    uint64 `json:"coreID"`    // "core id" field
}

// Cpu
// Representation of a processor modeled after /proc/cpuinfo
type Cpu struct {
	CpuName  string `json:"cpuName"`  // CPU model name
	CpuCores uint64 `json:"cpuCores"` // Number of CPU cores
	Socket   uint64 `json:"socket"`   // "physical id" for multiprocessor systems
	Cores    []Core `json:"cores"`
}