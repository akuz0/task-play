package playlist

import (
	"context"
	playListDomain "homework/internal/domain/playlist"
)

type PlayListService interface {
	GetPlayListByID(ctx context.Context, playListId int) (playListDomain.Play, error)
	CreatePlayList(ctx context.Context, playlist playListDomain.Play) (playListDomain.Play, error)
	StartPlayList(ctx context.Context, playListId int) (playListDomain.Play, error)
	NextSong(ctx context.Context)
	PrevSong(ctx context.Context)
	AddSong(ctx context.Context, songId int, playListId int) (playListDomain.AddSong, error)
}
