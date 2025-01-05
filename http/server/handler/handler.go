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
		json.NewEncoder(w).Encode(fmt.Errorf("bad param row"))
		return
	}

	col, err := strconv.Atoi(r.URL.Query().Get("col"))
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fmt.Errorf("bad param col"))
		return
	}

	ls.Mutex.Lock()
	err = ls.LifeService.World.InvertCell(row, col)
	ls.Mutex.Unlock()

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fmt.Errorf("bad param col"))
		return
	}

	// После переключения возвращаем JSON state
	ls.GetState(w, r)
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
