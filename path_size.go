package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	byteCoef = 1024
	KB       = int64(byteCoef)
	MB       = KB * byteCoef
	GB       = MB * byteCoef
	TB       = GB * byteCoef
	PB       = TB * byteCoef
	EB       = PB * byteCoef
)

func GetPathSize(path string, includeHidden bool, recursive bool, human bool) (string, error) {
	s, err := getSize(path, includeHidden, recursive)

	if err != nil {
		return "", err
	}

	return BuildOutput(s, path, human), nil
}

func getSize(path string, includeHidden bool, recursive bool) (int64, error) {
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

	entries, err := os.ReadDir(path)

	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		if !includeHidden && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		if entry.IsDir() {
			if !recursive {
				continue
			}

			p := filepath.Join(path, entry.Name())

			s, er := getSize(p, includeHidden, recursive)

			if er != nil {
				return 0, er
			}

			total += s
			continue
		}

		info, e := entry.Info()

		if e != nil {
			continue
		}

		total += info.Size()
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
		return FormatSize(size)
	}

	return DefaultFormat(size)
}

func PrintSize(size string, path string) {
	fmt.Printf("%s\t%s", size, path)
}
