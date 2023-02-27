package v1

import (
	"homework/specs"
	"net/http"
)

// Быстрая проверка актуальности текущего интерфейса сервера.
var _ specs.ServerInterface = &apiServer{}

type apiServer struct {
}

func (a apiServer) AddSong(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (a apiServer) NextSong(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (a apiServer) PauseList(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (a apiServer) PlayList(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (a apiServer) PrevSong(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewAPIServer() specs.ServerInterface {
	return &apiServer{}
}
