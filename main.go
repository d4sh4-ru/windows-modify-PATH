package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/windows/registry"
)

// Функция получения текущего PATH
func getCurrentPath() (string, registry.Key, error) {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.ALL_ACCESS)
	if err != nil {
		return "", 0, fmt.Errorf("ошибка при открытии реестра: %v", err)
	}

	currentPath, _, err := key.GetStringValue("Path")
	if err != nil {
		key.Close()
		return "", 0, fmt.Errorf("ошибка при чтении PATH: %v", err)
	}

	return currentPath, key, nil
}

// Функция обновления PATH
func updatePath(newPath string, key registry.Key) error {
	err := key.SetStringValue("Path", newPath)
	if err != nil {
		return fmt.Errorf("ошибка при записи нового значения PATH: %v", err)
	}
	return nil
}

// Функция добавления пути
func addPath(path string) {
	currentPath, key, err := getCurrentPath()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer key.Close()

	paths := strings.Split(currentPath, ";")
	for _, p := range paths {
		if strings.EqualFold(p, path) {
			fmt.Printf("Путь '%s' уже существует в PATH.\n", path)
			return
		}
	}

	paths = append(paths, path)
	err = updatePath(strings.Join(paths, ";"), key)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Путь '%s' успешно добавлен в PATH.\n", path)
}

// Функция удаления пути
func removePath(path string) {
	currentPath, key, err := getCurrentPath()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer key.Close()

	paths := strings.Split(currentPath, ";")
	newPaths := make([]string, 0, len(paths))
	found := false

	for _, p := range paths {
		if !strings.EqualFold(p, path) {
			newPaths = append(newPaths, p)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Printf("Путь '%s' не найден в PATH.\n", path)
		return
	}

	err = updatePath(strings.Join(newPaths, ";"), key)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Путь '%s' успешно удален из PATH.\n", path)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Использование: go run main.go [add|remove] <путь>")
		return
	}

	operation, path := os.Args[1], os.Args[2]

	switch operation {
	case "add":
		addPath(path)
	case "remove":
		removePath(path)
	default:
		fmt.Println("Неверная операция. Используйте 'add' или 'remove'.")
	}
}
