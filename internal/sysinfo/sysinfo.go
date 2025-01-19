// Package sysinfo provides functionality to retrieve system information
// from the host machine including OS details, architecture, and runtime information.
package sysinfo

import (
	"os"
	"runtime"
)

// SystemInfo holds various system-related information about the host machine.
type SystemInfo struct {
	OS string
	Arch string
	NumCPU int	
	Version string
	Hostname string 
	WorkingDir string
}

// GetSystemInfo retrieves current system information and returns it as a SystemInfo struct.
// It includes details about the operating system, architecture, CPU count, Go version,
// hostname, and current working directory.
func GetSystemInfo() SystemInfo {

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	workingDir, err := os.Getwd()
	if err != nil {
		workingDir = "unknown"
	}

	return SystemInfo{
		OS: runtime.GOOS,
		Arch: runtime.GOARCH,
		NumCPU: runtime.NumCPU(),
		Version: runtime.Version(),
		Hostname: hostname,
		WorkingDir: workingDir,
	}
}

// String implements the Stringer interface for pretty-printing SystemInfo.
// Returns a formatted multi-line string containing all system information.
func (si SystemInfo) String() string {
	return "OS: " + si.OS + "\n" +
		"Arch: " + si.Arch + "\n" +
		"NumCPU: " + string(rune(si.NumCPU)) + "\n" +
		"Version: " + si.Version + "\n" +
		"Hostname: " + si.Hostname + "\n" +
		"WorkingDir: " + si.WorkingDir
}