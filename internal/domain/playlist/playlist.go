package playlist

// AddSong defines model for AddSong.
type AddSong struct {
	// Id Идентификатор трека
	Id int `json:"id"`

	// PlayListId Идентификатор плейлиста
	PlayListId int `json:"playListId"`
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
}
