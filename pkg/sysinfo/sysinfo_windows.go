package sysinfo // import "github.com/dims/moby-sysinfo/pkg/sysinfo"

// New returns an empty SysInfo for windows for now.
func New(quiet bool) *SysInfo {
	sysInfo := &SysInfo{}
	return sysInfo
}
