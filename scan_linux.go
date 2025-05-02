//go:build linux

package main

import (
	"fmt"
	"os"
	"os/user"
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
		ctime, owner := getLinuxMeta(info)

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

func getLinuxMeta(info os.FileInfo) (ctime string, owner string) {
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return "N/A", "N/A"
	}
	// Use Ctim for inode change time (Linux does not provide creation time in Stat_t)
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
