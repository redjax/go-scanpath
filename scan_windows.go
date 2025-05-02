//go:build windows

package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func scanDirectory(path string, limit int) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	fmt.Printf("%-30s %10s  %20s  %20s  %-10s  %s\n", "Name", "Size", "Created", "Modified", "Owner", "Permissions")
	count := 0
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		ctime := getWindowsCreationTime(info)
		owner := "N/A" // To implement: use Windows API for real owner

		fmt.Printf("%-30s %10d  %20s  %20s  %-10s  %s\n",
			info.Name(),
			info.Size(),
			ctime,
			info.ModTime().Format("2006-01-02 15:04:05"),
			owner,
			info.Mode(),
		)
		count++
		if limit > 0 && count >= limit {
			break
		}
	}
	return nil
}

func getWindowsCreationTime(info os.FileInfo) string {
	if stat, ok := info.Sys().(*syscall.Win32FileAttributeData); ok {
		t := time.Unix(0, stat.CreationTime.Nanoseconds())
		return t.Format("2006-01-02 15:04:05")
	}
	return "N/A"
}
