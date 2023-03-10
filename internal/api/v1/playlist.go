package v1

import (
	"encoding/json"
	"homework/internal/domain/playlist"
	handlerc "homework/lib"
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
	res, err := a.playListService.AddSong(r.Context(), p.Id, p.SongName, 0)
	if err != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
	//todo move to service
	handlerc.Add(p.Id, p.SongName)

	errJson := json.NewEncoder(w).Encode(res)
	if errJson != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
}

func (a apiServer) NextSong(w http.ResponseWriter, r *http.Request) {
	errorText := "Not have next song"
	res, err := a.playListService.NextSong(r.Context())
	if err != nil {
		if err.Error() == errorText {
			json.NewEncoder(w).Encode(&playlist.Response{400, errorText})
			return
		}
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
	decoder := json.NewDecoder(r.Body)
	var p playlist.Play
	err := decoder.Decode(&p)
	if err != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
	handlerc.Pause()
	res, err := a.playListService.PauseSong(r.Context(), p.Id)
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

	//todo move to service
	handlerc.Start(p.Name)
	res.Name = p.Name

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
		if err.Error() == "Not have prev song" {
			json.NewEncoder(w).Encode(&playlist.Response{400, "Not have prev song"})
			return
		}
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
	errJson := json.NewEncoder(w).Encode(res)
	if errJson != nil {
		json.NewEncoder(w).Encode(&ResponseError{"Error", 500})
		return
	}
}
