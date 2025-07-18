//go:build linux

package scan

import (
	"fmt"
	"os"
	"os/user"
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
		ctime, owner := getLinuxMeta(info)
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

func getLinuxMeta(info os.FileInfo) (ctime string, owner string) {
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return "N/A", "N/A"
	}
	t := time.Unix(stat.Ctim.Sec, stat.Ctim.Nsec)
	ctime = t.Format("2006-01-02 15:04:05")

	uid := fmt.Sprint(stat.Uid)
	gid := fmt.Sprint(stat.Gid)
	u, err := user.LookupId(uid)
	if err == nil {
		uid = u.Username
	}
	g, err := user.LookupGroupId(gid)
	if err == nil {
		gid = g.Name
	}
	owner = fmt.Sprintf("%s:%s", uid, gid)
	return
}
