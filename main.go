package main

import (
	"net/http"

	"github.com/leonlarsson/the-finals-api-go/handlers"
	"github.com/leonlarsson/the-finals-api-go/middleware"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /v1/leaderboard/{leaderboardId}", handlers.HandleLeaderboard)
	router.HandleFunc("GET /v1/leaderboard/{leaderboardId}/{platformId}", handlers.HandleLeaderboard)

	println("Starting server...")
	http.ListenAndServe(":80", middleware.Logging(router))
}
