package model

// Gpu
// Abstraction of a GPU
// TODO: Get CUDA Cores
type Gpu struct {
	GpuName string `json:"gpuName"` // GPU model name
	GpuID   string `json:"gpuID"`   // "GPU UUID"
	BusID   string `json:"busID"`   // "Bus Location" PCIe bus ID
	IRQ     uint64 `json:"IRQ"`     // "IRQ" GPU Interrupt lane
	Major   uint64 `json:"major"`   //
	Minor   uint64 `json:"minor"`   // "Device Minor" for /dev/nvidia<minor> character device
}