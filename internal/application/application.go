package application

import (
	"fmt"
	"net/http"

	"github.com/klimenkokayot/game-of-life-go/http/server"
)

func Run(height, width int) {
	handler := server.Run(height, width)
	fmt.Printf("Server started at port :8080\n")
	http.ListenAndServe(":8080", handler)
}
