package life

import (
	"fmt"
	"math/rand"

	"github.com/klimenkokayot/game-of-life-go/pkg/support"
)

type World struct {
	Height uint
	Width  uint
	Cells  [][]bool
}

// Определяет количество живых соседей у клетки
func (w *World) Neighbours(x, y int) int {
	cnt := 0
	cnt += support.B2I(w.Cells[x-1][y])
	cnt += support.B2I(w.Cells[x-1][y+1])
	cnt += support.B2I(w.Cells[x][y+1])
	cnt += support.B2I(w.Cells[x+1][y+1])
	cnt += support.B2I(w.Cells[x+1][y])
	cnt += support.B2I(w.Cells[x+1][y-1])
	cnt += support.B2I(w.Cells[x][y-1])
	cnt += support.B2I(w.Cells[x-1][y-1])
	return cnt
}

// Определяет состояние клетки в следующем состоянии
func (w *World) Next(x, y int) bool {
	n := w.Neighbours(x, y)
	alive := w.Cells[x][y]
	if alive && (n > 4 || n < 2) {
		alive = false
	}
	if !alive && n == 3 {
		alive = true
	}
	return alive
}

// Обновляет состояние мира на следующий этап
func (w *World) NextState() {
	for i := 1; i < int(w.Height)-1; i++ {
		for j := 1; j < int(w.Width)-1; j++ {
			w.Cells[i][j] = w.Next(i, j)
		}
	}
}

// Создает новый рандомный мир
func (w *World) Seed() {
	for i := 1; i < int(w.Height)-1; i++ {
		for j := 1; j < int(w.Width)-1; j++ {
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
func NewWorld(height, width uint) *World {
	cells := make([][]bool, height+2)
	for i := range cells {
		cells[i] = make([]bool, width+2)
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}
}

////////////////////////////////////////////////

func (w *World) Print() {
	brownSquare := "\xF0\x9F\x9F\xAB"
	greenSquare := "\xF0\x9F\x9F\xA9"
	for i := 1; i < int(w.Height)-1; i++ {
		for j := 1; j < int(w.Width)-1; j++ {
			if w.Cells[i][j] {
				fmt.Print(greenSquare)
			} else {
				fmt.Print(brownSquare)
			}
		}
		if i+1 != int(w.Height) {
			fmt.Printf("\n")
		}
	}
}
