package server

import (
	"net/http"
	"path/filepath"
	"sync"

	"github.com/klimenkokayot/game-of-life-go/http/server/handler"
	"github.com/klimenkokayot/game-of-life-go/internal/service"
)

func Run(height, width int) http.Handler {
	mux := http.NewServeMux()

	tmp := service.New(height, width)
	ls := handler.LifeState{
		LifeService: tmp,
		Mutex:       &sync.Mutex{},
	}

	mux.HandleFunc("/game", ls.Index)
	mux.HandleFunc("/api/v1/state", ls.GetState)
	mux.HandleFunc("/api/v1/next", ls.NextState)
	mux.HandleFunc("/api/v1/toggle", ls.ToggleCell)
	mux.HandleFunc("/api/v1/seed", ls.Seed)
	mux.HandleFunc("/api/v1/neighbours", ls.GetNumNeighbours)
	mux.HandleFunc("/api/v1/near_neighbours", ls.GetNearNumNeighbours)

	/*
	 * TODO сделать нормально
	 */
	mux.HandleFunc("/api/v1/size", ls.Size)

	staticDir := filepath.Join(".", "web", "static")
	fs := http.FileServer(http.Dir(staticDir))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
