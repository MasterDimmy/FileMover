package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Сканирует указанный каталог, на основе даты изменения файлов создает каталог YYYY/MM и переносит в него файлы.")	

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

// moveFile перемещает файл из одного места в другое
func moveFile(src, dst string) error {
	fmt.Printf("%s => %s\n",src,dst)
	err := func() error {
		input, err := os.Open(src)
		if err != nil {
			return err
		}
		defer input.Close()

		output, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer output.Close()

		_, err = io.Copy(output, input)
		if err != nil {
			return err
		}
		return nil
	}()
	if err!=nil {
		return err
	}

	return os.Remove(src)
}

