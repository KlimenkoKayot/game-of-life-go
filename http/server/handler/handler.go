package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"sync"
	"text/template"

	"github.com/klimenkokayot/game-of-life-go/internal/service"
)

type LifeState struct {
	LifeService *service.LifeService
	Mutex       *sync.Mutex
}

func (ls *LifeState) Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join(".", "web", "templates", "index.html")))
	tmpl.Execute(w, nil)
}

/////////////////////////////

func (ls *LifeState) Seed(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("fill")
	fill, err := strconv.Atoi(data)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fmt.Errorf("bad param fill"))
		return
	}

	ls.Mutex.Lock()
	ls.LifeService.World.Seed(fill)
	ls.Mutex.Unlock()

	ls.GetState(w, r)
}

func (ls *LifeState) GetNumNeighbours(w http.ResponseWriter, r *http.Request) {
	ls.Mutex.Lock()
	data := ls.LifeService.World.NumNeighbours
	ls.Mutex.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (ls *LifeState) GetNearNumNeighbours(w http.ResponseWriter, r *http.Request) {
	row, _ := strconv.Atoi(r.URL.Query().Get("row"))
	col, _ := strconv.Atoi(r.URL.Query().Get("col"))

	ls.Mutex.Lock()
	data, err := ls.LifeService.World.GetNearNumNeighbours(row, col)
	ls.Mutex.Unlock()

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Меняет состояние на следующее, возвращает новое поле
func (ls *LifeState) NextState(w http.ResponseWriter, r *http.Request) {
	ls.Mutex.Lock()
	ls.LifeService.World.NextState()
	ls.Mutex.Unlock()

	ls.GetState(w, r)
}

// Инвертирует переменную, возвращает JSON state
func (ls *LifeState) ToggleCell(w http.ResponseWriter, r *http.Request) {
	row, err := strconv.Atoi(r.URL.Query().Get("row"))
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	col, err := strconv.Atoi(r.URL.Query().Get("col"))
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	ls.Mutex.Lock()
	err = ls.LifeService.World.InvertCell(row, col)
	ls.Mutex.Unlock()

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// После переключения возвращаем JSON state текущей клетки
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fmt.Sprintf("%d", ls.LifeService.World.Cells[row][col]))
}

// JSON
func (ls *LifeState) GetState(w http.ResponseWriter, r *http.Request) {
	ls.Mutex.Lock()
	data := ls.LifeService.World.Cells
	ls.Mutex.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

/*
 * TODO сделать нормально
 */
func (ls *LifeState) Size(w http.ResponseWriter, r *http.Request) {
	ls.Mutex.Lock()
	data := []int{ls.LifeService.World.Height, ls.LifeService.World.Width}
	ls.Mutex.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
