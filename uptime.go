package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"time"
)

// getprocessuptime retrieves the process uptime using winapi
func getprocessuptime(pid uint32) (time.Duration, error) {
	hProcess, err := windows.OpenProcess(
		windows.PROCESS_QUERY_LIMITED_INFORMATION,
		false,
		pid,
	)

	if err != nil {
		return 0, fmt.Errorf("cant open a process (possibly requires elevation): %w", err)
	}
	defer func(handle windows.Handle) {
		_ = windows.CloseHandle(handle)
	}(hProcess)
	
	var creationTime, exitTime, kernelTime, userTime windows.Filetime
	err = windows.GetProcessTimes(
		hProcess,
		&creationTime,
		&exitTime,
		&kernelTime,
		&userTime,
	)

	if err != nil {
		return 0, fmt.Errorf("cant catch creationtime: %w", err)
	}
	creation := time.Unix(0, creationTime.Nanoseconds())
	uptime := time.Since(creation)

	return uptime, nil
}
