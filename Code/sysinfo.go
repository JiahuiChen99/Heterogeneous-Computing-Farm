package model

// SysInfo
// Wrapper struct of 'Utsname' returned by 'uname' system call
type SysInfo struct {
	SysName  string `json:"sys_name"`
	NodeName string `json:"node_name"`
	Release  string `json:"release"`
	Version  string `json:"version"`
	Machine  string `json:"machine"`
}