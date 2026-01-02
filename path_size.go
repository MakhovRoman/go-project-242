package code

import (
	"fmt"
	"os"
)

const (
	KB = int64(1000)
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
)

func GetSize(path string) (int64, error) {
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
		info, e := file.Info()
		//macOS добавляет .DS_Store
		if e != nil || info.IsDir() || info.Name() == ".DS_Store" {
			continue
		}

		total += info.Size()
	}

	return total, nil
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
