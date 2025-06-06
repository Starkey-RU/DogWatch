package main

import (
	"encoding/csv"
	"fmt"
	wapi "github.com/codehardt/go-win64api"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func DogWatch(ch <-chan string) {
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	outputDir := "output"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Println("Ошибка создания директории:", err)
		return
	}
	filename := filepath.Join(outputDir, fmt.Sprintf("process_%s.csv", timestamp))
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	// хоть ide говорит, что так лучше не делать - я потом с файлом же не работают
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// титульник csv
	writer.Write([]string{"PID", "Username", "Executable", "Fullpath", "Uptime"})

	for user := range ch {
		pr, err := wapi.ProcessList()
		if err != nil {
			fmt.Printf("Ошибка получения списка процессов: %v\n", err)
			continue
		}
		for _, p := range pr {
			if !strings.EqualFold(p.Username, user) {
				continue
			}
			if strings.Contains(p.Fullpath, "C:\\Windows\\System32") {
				continue
			}
			pid := uint32(p.Pid)
			uptimeDur, err := getprocessuptime(pid)
			if err != nil {
				continue
			}
			record := []string{
				fmt.Sprintf("%d", pid),
				p.Username,
				p.Executable,
				p.Fullpath,
				uptimeDur.String(),
			}
			if err := writer.Write(record); err != nil {
				fmt.Println("Ошибка записи в CSV:", err)
			}
		}
	}
	//dbg
	fmt.Println("CSV создан:", filename)
}
