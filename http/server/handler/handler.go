package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/klimenkokayot/game-of-life-go/internal/service"
)

type LifeState struct {
	*service.LifeService
}

func (ls *LifeState) View(w http.ResponseWriter, r *http.Request) {
	str := ls.World.String()
	w.Write([]byte(str))
	ls.World.NextState()
}

func (ls *LifeState) Seed(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("fill")
	fill, _ := strconv.Atoi(data)
	ls.World.Seed(fill)
}

func (ls *LifeState) SetTrue(w http.ResponseWriter, r *http.Request) {
	strx := r.FormValue("x")
	stry := r.FormValue("y")
	x, _ := strconv.Atoi(strx)
	y, _ := strconv.Atoi(stry)
	err := ls.World.SetTrue(x, y)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		io.WriteString(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (ls *LifeState) SetFalse(w http.ResponseWriter, r *http.Request) {
	strx := r.FormValue("x")
	stry := r.FormValue("y")
	x, _ := strconv.Atoi(strx)
	y, _ := strconv.Atoi(stry)
	err := ls.World.SetFalse(x, y)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		io.WriteString(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Меняет состояние на следующее
func (ls *LifeState) NextState(w http.ResponseWriter, r *http.Request) {
	ls.World.NextState()
	w.WriteHeader(http.StatusOK)
}

// JSON
func (ls *LifeState) GetState(w http.ResponseWriter, r *http.Request) {
	str := ls.World.String()
	tojson := make(map[string]interface{})
	tojson["result"] = str

	data, err := json.Marshal(tojson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, data)
}
