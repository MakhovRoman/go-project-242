package pathsize

import (
	"fmt"
	"os"
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

		if e != nil || info.IsDir() {
			continue
		}

		total += info.Size()
	}

	return total, nil
}

func PrintSize(size int64, path string) {
	//TODO: можно добавить уведомление, что папка не содержит вложенных файлов
	fmt.Printf("%dB	%s\n", size, path)
}
