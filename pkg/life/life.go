package life

import (
	"fmt"
	"math/rand"
)

type World struct {
	Height        int
	Width         int
	Cells         [][]int
	NumNeighbours [][]int
}

/*
 * Определяет количество живых соседей у клетки
 * (на торе)
 */
func (w *World) Neighbours(x, y int) (int, error) {
	cnt := 0
	x, y = w.CheckPosition(x, y)

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
 *	Если клетка вне поля, то возвращаем ошибку и -1,
 *  иначе - состояние клетки
 */
func (w *World) GetCellState(x, y int) (int, error) {
	x, y = w.CheckPosition(x, y)
	return w.Cells[x][y], nil
}

func (w *World) GetCellNumNeighbours(x, y int) (int, error) {
	x, y = w.CheckPosition(x, y)
	return w.NumNeighbours[x][y], nil
}

func (w *World) GetNearNumNeighbours(x, y int) ([][]int, error) {
	data := make([][]int, 3)
	for i := range data {
		data[i] = make([]int, 3)
	}
	// смещение от центральной клетки
	move := []int{-1, 0, 1}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			data[i][j], _ = w.GetCellNumNeighbours(x+move[i], y+move[j])
		}
	}
	return data, nil
}

func (w *World) ResetNearNeighbours(x, y int) error {
	x, y = w.CheckPosition(x, y)

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

	add := 0
	if w.Cells[x][y] == 1 {
		add = 1
	} else {
		add = -1
	}
	w.NumNeighbours[xl][y] += add
	w.NumNeighbours[xl][yl] += add
	w.NumNeighbours[xl][yr] += add
	w.NumNeighbours[xr][y] += add
	w.NumNeighbours[xr][yl] += add
	w.NumNeighbours[xr][yr] += add
	w.NumNeighbours[x][yl] += add
	w.NumNeighbours[x][yr] += add

	return nil
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
	if alive == 1 && (n > 3 || n < 2) {
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
	newState := make([][]int, w.Height)
	cntState := make([][]int, w.Height)
	for i := range w.Height {
		newState[i] = make([]int, w.Width)
		cntState[i] = make([]int, w.Width)
	}
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			newState[i][j], _ = w.Next(i, j)
		}
	}
	w.Cells = newState
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			cntState[i][j], _ = w.Neighbours(i, j)
		}
	}
	w.NumNeighbours = cntState
}

// Инвертирует состояние клетки
func (w *World) InvertCell(x, y int) error {
	x, y = w.CheckPosition(x, y)
	if w.Cells[x][y] == 1 {
		w.Cells[x][y] = 0
		w.ResetNearNeighbours(x, y)
	} else {
		w.Cells[x][y] = 1
		w.ResetNearNeighbours(x, y)
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
			if rand.Intn(100+1) < fill {
				w.Cells[i][j] = 1
			}
		}
	}
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			w.NumNeighbours[i][j], _ = w.Neighbours(i, j)
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
	cells2 := make([][]int, height)
	for i := range cells {
		cells[i] = make([]int, width)
		cells2[i] = make([]int, width)
	}
	return &World{
		Height:        height,
		Width:         width,
		Cells:         cells,
		NumNeighbours: cells2,
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

// Валидация координатов точки на торе, возвращает верные координаты в случае попытки выхода за поле
func (w *World) CheckPosition(x, y int) (int, int) {
	if x == -1 {
		x = w.Height - 1
	} else if x == w.Height {
		x = 0
	}
	if y == -1 {
		y = w.Width - 1
	} else if y == w.Width {
		y = 0
	}
	return x, y
}
