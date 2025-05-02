//go:build windows

package scan

import (
	"fmt"
	"os"
	"scanpath/internal/tbl"
	"syscall"
	"time"
)

func ScanDirectory(path string, limit int, sortColumn, sortOrder string, filterString string) error {
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
		// To implement: use Windows API for real owner
		owner := "N/A"
		size := info.Size()
		sizeParsed := tbl.ByteCountIEC(size)

		row := []string{
			info.Name(),
			// Size in bytes
			fmt.Sprintf("%d", size),
			// Human-readable size
			sizeParsed,
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
	// Parse and apply filter before sorting/printing
	var filterExpr *tbl.FilterExpr
	if filterString != "" {
		var err error
		filterExpr, err = tbl.ParseFilter(filterString)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid filter: %v\n", err)
			// Optionally return err or continue without filtering
		}
	}
	results = tbl.ApplyFilter(results, filterExpr)

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
