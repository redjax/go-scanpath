//go:build windows

package scan

import (
	"fmt"
	"os"
	"scanpath/internal/tbl"
	"syscall"
	"time"
)

func ScanDirectory(path string, limit int, sortColumn, sortOrder string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	var results [][]string
	count := 0
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		ctime := getWindowsCreationTime(info)
		owner := "N/A" // To implement: use Windows API for real owner

		row := []string{
			info.Name(),
			fmt.Sprintf("%d", info.Size()),
			ctime,
			info.ModTime().Format("2006-01-02 15:04:05"),
			owner,
			info.Mode().String(),
		}
		results = append(results, row)
		count++
		if limit > 0 && count >= limit {
			break
		}
	}

	// Sort the table before printing
	tbl.SortResults(results, sortColumn, sortOrder)
	tbl.PrintScanResultsTable(results)

	return nil
}

func getWindowsCreationTime(info os.FileInfo) string {
	if stat, ok := info.Sys().(*syscall.Win32FileAttributeData); ok {
		t := time.Unix(0, stat.CreationTime.Nanoseconds())
		return t.Format("2006-01-02 15:04:05")
	}
	return "N/A"
}
