package v1

import (
	"homework/internal/service/playlist"
	"homework/specs"
)

// Быстрая проверка актуальности текущего интерфейса сервера.
var _ specs.ServerInterface = &apiServer{}

type apiServer struct {
	playListService playlist.PlayListService
}

type ResponseError struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
}

func NewAPIServer(playListService playlist.PlayListService) specs.ServerInterface {
	return &apiServer{
		playListService: playListService,
	}
}
