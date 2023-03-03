package playlist

// AddSong defines model for AddSong.
type AddSong struct {
	// Id Идентификатор трека
	Id int `json:"id"`

	// SongName Имя трека
	SongName string `json:"songName"`

	// SongName Имя трека
	Status string `json:"status"`
}

// Pause defines model for Pause.
type Pause struct {
	// Id Идентификатор плейлиста
	Id int `json:"id"`
}

// Play defines model for Play.
type Play struct {
	// Id Идентификатор плейлиста
	Id int `json:"id"`

	// Name Имя трека
	Name string `json:"name"`
}

type Response struct {
	StatusCode       int    `json:"statusCode"`
	ErrorDescription string `json:"ErrorDescription"`
}
