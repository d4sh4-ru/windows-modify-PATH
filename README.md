# windows-modify-PATH

Этот проект предоставляет утилиту на Go для управления переменной окружения `PATH` в Windows. Он расширяет возможности NSIS (Nullsoft Scriptable Install System) по добавлению и удалению путей из `PATH` без необходимости городить костыли для редактирования реестра.

⚠ **Ограничение NSIS**: В стандартном NSIS переменная `$PATH` может обрезаться, так как NSIS работает с ограниченной длиной строк (обычно до 1024 или 4096 символов в зависимости от версии). Это может привести к некорректному обновлению `PATH`, если он слишком длинный. Использование этой утилиты позволяет обходить это ограничение.

## Возможности

- Добавление пути в `PATH`
- Удаление пути из `PATH`
- Проверка наличия пути перед добавлением
- Работа с реестром Windows через `golang.org/x/sys/windows/registry`

## Установка

1. Убедитесь, что у вас установлен Go.
2. Скачайте репозиторий и соберите бинарный файл:
   ```sh
   go build -o modify_path.exe main.go
   ```

## Использование

```sh
modify_path.exe [add|remove] <путь>
```

Примеры:
- Добавить путь:
  ```sh
  modify_path.exe add "C:\MyProgram\bin"
  ```
- Удалить путь:
  ```sh
  modify_path.exe remove "C:\MyProgram\bin"
  ```

## Лицензия
MIT
```
MIT License

Copyright (c) 2025 Скулкин Денис

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```