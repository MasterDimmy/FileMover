package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Сканирует указанный каталог, на основе даты изменения файлов создает YYYY/MM там же и переносит туда файлы.")	

	// Флаг для указания каталога
	dirPtr := flag.String("dir", ".", "directory to scan")
	flag.Parse()

	if *dirPtr == "." {
		flag.Usage()
		return
	}


	// Проверка и сканирование каталога
	err := filepath.Walk(*dirPtr, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Пропуск, если это каталог
		if info.IsDir() {
			return nil
		}

		// Получение времени последнего изменения файла
		modTime := info.ModTime()

		// Формирование пути назначения
		newDir := filepath.Join(*dirPtr, fmt.Sprintf("%d", modTime.Year()), fmt.Sprintf("%02d", modTime.Month()))
		err = os.MkdirAll(newDir, os.ModePerm)
		if err != nil {
			return err
		}

		// Перемещение файла
		newPath := filepath.Join(newDir, info.Name())
		moveFile(path, newPath)
		return nil
	})

	if err != nil {
		fmt.Println("Ошибка при сканировании каталога:", err)
	}
}

func moveFile(src, dst string) error {
    fmt.Printf("%s => %s\n",src,dst)

    return os.Rename(src, dst)
}



