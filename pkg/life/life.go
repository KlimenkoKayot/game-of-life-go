package life

import (
	"fmt"
	"math/rand"

	"github.com/klimenkokayot/game-of-life-go/pkg/support"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

// Определяет количество живых соседей у клетки
func (w *World) Neighbours(x, y int) (int, error) {
	cnt := 0
	if err := w.CheckPosition(x, y); err != nil {
		return 0, err
	}

	if x != 0 {
		cnt += support.B2I(w.Cells[x-1][y])
		if y != 0 {
			cnt += support.B2I(w.Cells[x-1][y-1])
		}
		if y != w.Width-1 {
			cnt += support.B2I(w.Cells[x-1][y+1])
		}
	}
	if x != w.Height-1 {
		cnt += support.B2I(w.Cells[x+1][y])
		if y != 0 {
			cnt += support.B2I(w.Cells[x+1][y-1])
		}
		if y != w.Width-1 {
			cnt += support.B2I(w.Cells[x+1][y+1])
		}
	}
	if y != 0 {
		cnt += support.B2I(w.Cells[x][y-1])
	}
	if y != w.Width-1 {
		cnt += support.B2I(w.Cells[x][y+1])
	}

	return cnt, nil
}

// Определяет состояние клетки в следующем состоянии
func (w *World) Next(x, y int) (bool, error) {
	n, err := w.Neighbours(x, y)
	if err != nil {
		return false, err
	}

	alive := w.Cells[x][y]
	if alive && (n > 4 || n < 2) {
		alive = false
	}
	if !alive && n == 3 {
		alive = true
	}
	return alive, nil
}

/*
 * Обновляет состояние мира на следующий этап
 * noexcept
 */
func (w *World) NextState() {
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			w.Cells[i][j], _ = w.Next(i, j)
		}
	}
}

// Меняет состояние клетки на True
func (w *World) SetTrue(x, y int) error {
	if err := w.CheckPosition(x, y); err != nil {
		return nil
	}
	w.Cells[x][y] = true
	return nil
}

// Меняет состояние клетки на False
func (w *World) SetFalse(x, y int) error {
	if err := w.CheckPosition(x, y); err != nil {
		return nil
	}
	w.Cells[x][y] = false
	return nil
}

/*
 * Создает новый рандомный мир
 * noexcept
 */
func (w *World) Seed() {
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			w.Cells[i][j] = false
			if rand.Intn(10) == 0 {
				w.Cells[i][j] = true
			}
		}
	}
}

/*
 * Функция создания мира с размерами высота и ширина
 * Индексация поля с 1
 */
func NewWorld(height, width int) *World {
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}
}

/*
 * Функции для вывода в консоль, дебага и всякой хрени
 *
 */

var (
	trueSquare  = "\xF0\x9F\x9F\xAB"
	falseSquare = "\xF0\x9F\x9F\xA9"
)

// Конвертирует состояние игры в строковый формат
func (w *World) String() string {
	result := ""
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			if w.Cells[i][j] {
				result += trueSquare
			} else {
				result += falseSquare
			}
		}
		if i+1 != w.Height {
			result += "\n"
		}
	}
	return result
}

// Валидация координатов точки (норм или выходит за поле)
func (w *World) CheckPosition(x, y int) error {
	if x < 0 || x >= w.Width {
		return fmt.Errorf("x must be: %d <= x < %d", 0, w.Width)
	}
	if y < 0 || y >= w.Height {
		return fmt.Errorf("y must be: %d <= y < %d", 0, w.Height)
	}
	return nil
}
