package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	KB = int64(1000)
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
)

func GetPathSize(path string, includeHidden bool, recursive bool) (int64, error) {
	if !includeHidden && hasHiddenSegment(path) {
		return 0, nil
	}

	info, err := os.Lstat(path)

	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	var total int64

	fileList, err := os.ReadDir(path)

	if err != nil {
		return 0, err
	}

	for _, file := range fileList {
		entry, e := file.Info()

		if e != nil {
			continue
		}

		if !includeHidden && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		if entry.IsDir() {
			if !recursive {
				continue
			}

			p := filepath.Join(path, entry.Name())

			s, er := GetPathSize(p, includeHidden, recursive)

			if er != nil {
				return 0, er
			}

			total += s
			continue
		}

		total += entry.Size()
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

func DefaultFormat(size int64) string {
	return fmt.Sprintf("%dB", size)
}

func FormatSize(size int64) string {
	var result float64
	unit := "B"

	switch {
	case size >= EB:
		result = float64(size) / float64(EB)
		unit = "EB"
	case size >= PB:
		result = float64(size) / float64(PB)
		unit = "PB"
	case size >= TB:
		result = float64(size) / float64(TB)
		unit = "TB"
	case size >= GB:
		result = float64(size) / float64(GB)
		unit = "GB"
	case size >= MB:
		result = float64(size) / float64(MB)
		unit = "MB"
	case size >= KB:
		result = float64(size) / float64(KB)
		unit = "KB"
	default:
		result = float64(size)
	}

	return fmt.Sprintf("%.1f%s", result, unit)
}

func BuildOutput(size int64, path string, human bool) string {
	if human {
		return fmt.Sprintf("%s\t%s\n", FormatSize(size), path)
	}

	return fmt.Sprintf("%s\t%s\n", DefaultFormat(size), path)
}

func PrintSize(size int64, path string, human bool) {
	fmt.Print(BuildOutput(size, path, human))
}
