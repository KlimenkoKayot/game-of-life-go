package server

import (
	"net/http"

	"github.com/klimenkokayot/game-of-life-go/http/server/handler"
	"github.com/klimenkokayot/game-of-life-go/internal/service"
)

func Run(height, width int) http.Handler {
	mux := http.NewServeMux()

	tmp := service.New(height, width)
	ls := handler.LifeState{
		LifeService: tmp,
	}

	mux.HandleFunc("/api/v1/view", ls.View)
	mux.HandleFunc("/api/v1/seed", ls.Seed)
	mux.HandleFunc("/api/v1/settrue", ls.SetTrue)
	mux.HandleFunc("/api/v1/setfalse", ls.SetFalse)

	return mux
}
