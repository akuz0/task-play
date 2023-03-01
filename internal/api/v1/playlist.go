package v1

import (
	"encoding/json"
	"homework/internal/domain/playlist"
	"net/http"
)

func (a apiServer) AddSong(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var p playlist.AddSong
	err := decoder.Decode(&p)
	if err != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
	res, err := a.playListService.AddSong(r.Context(), p.Id, p.PlayListId)
	if err != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
	errJson := json.NewEncoder(w).Encode(res)
	if errJson != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
}

func (a apiServer) NextSong(w http.ResponseWriter, r *http.Request) {
	res, err := a.playListService.NextSong(r.Context())
	if err != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
	errJson := json.NewEncoder(w).Encode(res)
	if errJson != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
}

func (a apiServer) PauseList(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (a apiServer) PlayList(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var p playlist.Play
	err := decoder.Decode(&p)
	if err != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
	res, err := a.playListService.StartPlayList(r.Context(), p.Id)
	if err != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
	errJson := json.NewEncoder(w).Encode(res)
	if errJson != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
}

func (a apiServer) CreatePlayList(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var p playlist.Play
	err := decoder.Decode(&p)
	if err != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
	res, err := a.playListService.CreatePlayList(r.Context(), p.Id)
	if err != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
	errJson := json.NewEncoder(w).Encode(res)
	if errJson != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
}

func (a apiServer) PrevSong(w http.ResponseWriter, r *http.Request) {
	res, err := a.playListService.PrevSong(r.Context())
	if err != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
	errJson := json.NewEncoder(w).Encode(res)
	if errJson != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
}
