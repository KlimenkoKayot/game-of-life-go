package life

import (
	"fmt"
	"math/rand"
)

type World struct {
	Height int
	Width  int
	Cells  [][]int
}

/*
 * Определяет количество живых соседей у клетки
 * (на торе)
 */
func (w *World) Neighbours(x, y int) (int, error) {
	cnt := 0
	if err := w.CheckPosition(x, y); err != nil {
		return 0, err
	}

	var (
		xl, xr, yl, yr int
	)
	if x != 0 {
		xl = x - 1
	} else {
		xl = w.Height - 1
	}
	if x == w.Height-1 {
		xr = 0
	} else {
		xr = x + 1
	}
	if y != 0 {
		yl = y - 1
	} else {
		yl = w.Width - 1
	}
	if y == w.Width-1 {
		yr = 0
	} else {
		yr = y + 1
	}
	cnt += w.Cells[xl][y]
	cnt += w.Cells[xl][yl]
	cnt += w.Cells[xl][yr]
	cnt += w.Cells[xr][y]
	cnt += w.Cells[xr][yl]
	cnt += w.Cells[xr][yr]
	cnt += w.Cells[x][yl]
	cnt += w.Cells[x][yr]

	return cnt, nil
}

/*
 * Определяет состояние клетки в следующем состоянии
 */
func (w *World) Next(x, y int) (int, error) {
	n, err := w.Neighbours(x, y)
	if err != nil {
		return 0, err
	}

	alive := w.Cells[x][y]
	if alive == 1 && (n > 4 || n < 2) {
		alive = 0
	}
	if alive == 0 && n == 3 {
		alive = 1
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

// Инвертирует состояние клетки
func (w *World) InvertCell(x, y int) error {
	if err := w.CheckPosition(x, y); err != nil {
		return err
	}
	if w.Cells[x][y] == 1 {
		w.Cells[x][y] = 0
	} else {
		w.Cells[x][y] = 1
	}
	return nil
}

/*
 * Создает новый рандомный мир с заполнением от 0 до 100
 * noexcept
 */
func (w *World) Seed(fill int) error {
	if fill < 0 || fill > 100 {
		return fmt.Errorf("fill must be: %d <= fill <= %d", 0, 100)
	}
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			w.Cells[i][j] = 0
			if rand.Intn(100+1) <= fill {
				w.Cells[i][j] = 1
			}
		}
	}
	return nil
}

/*
 * Функция создания мира с размерами высота и ширина
 * Индексация поля с 1
 */
func NewWorld(height, width int) *World {
	cells := make([][]int, height)
	for i := range cells {
		cells[i] = make([]int, width)
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
			if w.Cells[i][j] == 1 {
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
	if x < 0 || x >= w.Height {
		return fmt.Errorf("x must be: %d <= y < %d", 0, w.Width)
	}
	if y < 0 || y >= w.Width {
		return fmt.Errorf("y must be: %d <= y < %d", 0, w.Height)
	}
	return nil
}
