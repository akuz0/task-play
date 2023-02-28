package playlist

import (
	"context"
	playListDomain "homework/internal/domain/playlist"
)

type PlayListService interface {
	GetPlayListByID(ctx context.Context, playListId int) (playListDomain.Play, error)
	CreatePlayList(ctx context.Context, playListId int) (playListDomain.Play, error)
	StartPlayList(ctx context.Context, playListId int) (playListDomain.Play, error)
	NextSong(ctx context.Context)
	PrevSong(ctx context.Context)
	AddSong(ctx context.Context, songId int, playListId int) (playListDomain.AddSong, error)
}

type PlayListStorage interface {
	CreatePlayList(ctx context.Context, playListId int) (playListDomain.Play, error)
}

type service struct {
	playListStorage PlayListStorage
}

func (s service) GetPlayListByID(ctx context.Context, playListId int) (playListDomain.Play, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) CreatePlayList(ctx context.Context, playListId int) (playListDomain.Play, error) {
	playlist, err := s.playListStorage.CreatePlayList(ctx, playListId)
	if err != nil {
		return playListDomain.Play{}, err
	}
	return playlist, err
}

func (s service) StartPlayList(ctx context.Context, playListId int) (playListDomain.Play, error) {
	panic("implement me")
}

func (s service) NextSong(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (s service) PrevSong(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (s service) AddSong(ctx context.Context, songId int, playListId int) (playListDomain.AddSong, error) {
	//TODO implement me
	panic("implement me")
}

func NewPlayListService(playListStorage PlayListStorage) PlayListService {
	return &service{
		playListStorage: playListStorage,
	}
}
