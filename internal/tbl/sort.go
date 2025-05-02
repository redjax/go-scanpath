package tbl

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

// Map column names to indices
var columnIndices = map[string]int{
	"name":        0,
	"size":        1,
	"created":     2,
	"modified":    3,
	"owner":       4,
	"permissions": 5,
}

// Helper: parse time in your format
func parseTime(s string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", s)
}

func SortResults(results [][]string, column, order string) {
	idx, ok := columnIndices[strings.ToLower(column)]
	if !ok {
		idx = 0 // Default to "name"
	}
	desc := strings.ToLower(order) == "desc"

	sort.Slice(results, func(i, j int) bool {
		a, b := results[i][idx], results[j][idx]
		// Numeric sort for size
		if column == "size" {
			ai, _ := strconv.ParseInt(a, 10, 64)
			bi, _ := strconv.ParseInt(b, 10, 64)
			if desc {
				return ai > bi
			}
			return ai < bi
		}
		// Time sort for created/modified
		if column == "created" || column == "modified" {
			at, _ := parseTime(a)
			bt, _ := parseTime(b)
			if desc {
				return at.After(bt)
			}
			return at.Before(bt)
		}
		// Alphanumeric sort for others (case-insensitive)
		la, lb := strings.ToLower(a), strings.ToLower(b)
		if desc {
			return la > lb
		}
		return la < lb
	})
}
