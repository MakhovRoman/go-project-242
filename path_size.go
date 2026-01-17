// Package code provides utilities for calculating filesystem path sizes
package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetPathSize calculates the total size of a file or directory at the given path.
// It supports optional recursive traversal, inclusion of hidden files and human-readable output formatting.
func GetPathSize(path string, recursive, humanize, includeHidden bool) (string, error) {
	s, err := getSize(path, includeHidden, recursive)

	if err != nil {
		return "", err
	}

	return buildOutput(s, humanize), nil
}

func getSize(path string, includeHidden, recursive bool) (int64, error) {
	if !includeHidden && hasHiddenSegment(path) {
		return 0, nil
	}

	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if info.Mode()&os.ModeSymlink != 0 {
		target, err := filepath.EvalSymlinks(path)
		if err != nil {
			return 0, err
		}

		return getSize(target, includeHidden, recursive)
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	var total int64

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		if !includeHidden && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		p := filepath.Join(path, entry.Name())
		entryInfo, e := entry.Info()

		if e != nil {
			continue
		}

		if entryInfo.Mode()&os.ModeSymlink != 0 {
			s, er := getSize(p, includeHidden, recursive)
			if er != nil {
				return 0, er
			}
			total += s
			continue
		}

		if entry.IsDir() {
			if !recursive {
				continue
			}

			s, er := getSize(p, includeHidden, recursive)
			if er != nil {
				return 0, er
			}

			total += s
			continue
		}

		total += entryInfo.Size()
	}

	return total, nil
}

func hasHiddenSegment(path string) bool {
	clean := filepath.Clean(path)
	parts := strings.Split(clean, string(os.PathSeparator))

	for _, p := range parts {
		if strings.HasPrefix(p, ".") && p != "." && p != ".." {
			return true
		}
	}

	return false
}

func defaultFormat(size int64) string {
	return fmt.Sprintf("%dB", size)
}

func formatSize(size int64) string {
	val := float64(size)
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	i := 0

	for val >= 1024 && i < len(units)-1 {
		val /= 1024
		i++
	}

	if i == 0 {
		return fmt.Sprintf("%d%s", size, units[i])
	}

	return fmt.Sprintf("%.1f%s", val, units[i])
}

func buildOutput(size int64, humanize bool) string {
	if humanize {
		return formatSize(size)
	}

	return defaultFormat(size)
}
