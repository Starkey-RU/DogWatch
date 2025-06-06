package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"time"
)

// //[BACKUP LEGACY]
//func uptime(pid uint32) string {
//	duration, err := getprocessuptime(pid)
//	if err != nil {
//		return fmt.Sprintf("Ошибка: %v", err)
//	}
//	fmt.Printf("Процесс работает: %v\n", duration)
//	return duration.String()
//}

// getprocessuptime - буквально шаманит с winapi для подачи
func getprocessuptime(pid uint32) (time.Duration, error) {
	hProcess, err := windows.OpenProcess(
		windows.PROCESS_QUERY_LIMITED_INFORMATION,
		false,
		pid,
	)
	//я так понял, некоторые приложения под elevation premission могут не хвататься.
	if err != nil {
		return 0, fmt.Errorf("cant open a process: %w", err)
	}
	defer func(handle windows.Handle) {
		err := windows.CloseHandle(handle)
		if err != nil {
		}
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
