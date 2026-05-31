# 🧩 Тетрис на 4 языках программирования

Реализация классической игры Тетрис на **Python**, **JavaScript**, **C++** и **Go**.  
Проект демонстрирует, как одна и та же логика может быть воплощена в разных экосистемах.

---

## 📁 Содержание

| Язык | Папка | Запуск |
|------|-------|--------|
| Python | `python/` | `python3 tetris.py` или `python3 tetris_pygame.py` |
| JavaScript | `javascript/` | открыть `index.html` в браузере |
| C++ | `cpp/` | `g++ tetris.cpp -lncurses -o tetris && ./tetris` |
| Go | `go/` | `go run tetris.go` |

---

## 🎮 Управление

- Стрелка влево / вправо — движение
- Стрелка вверх — поворот фигуры
- Стрелка вниз — ускорение
- Пробел — мгновенный сброс
- `q` — выход

---

## 🧪 Пример кода (логика поворота на Go)

```go
func rotatePiece(piece [][]int) [][]int {
    rows := len(piece)
    cols := len(piece[0])
    rotated := make([][]int, cols)
    for i := range rotated {
        rotated[i] = make([]int, rows)
    }
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            rotated[j][rows-1-i] = piece[i][j]
        }
    }
    return rotated
}
