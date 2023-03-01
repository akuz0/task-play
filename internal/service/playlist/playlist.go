package playlist

import (
	"context"
	playListDomain "homework/internal/domain/playlist"
)

type PlayListService interface {
	GetPlayListByID(ctx context.Context, playListId int) (playListDomain.Play, error)
	CreatePlayList(ctx context.Context, playListId int) (playListDomain.Play, error)
	StartPlayList(ctx context.Context, playListId int) (playListDomain.Play, error)
	NextSong(ctx context.Context) (playListDomain.Play, error)
	PrevSong(ctx context.Context) (playListDomain.Play, error)
	AddSong(ctx context.Context, songId int, playListId int) (playListDomain.AddSong, error)
}

type PlayListStorage interface {
	CreatePlayList(ctx context.Context, playListId int) (playListDomain.Play, error)
	StartPlayList(ctx context.Context, playListId int) (playListDomain.Play, error)
	NextSong(ctx context.Context) (playListDomain.Play, error)
	PrevSong(ctx context.Context) (playListDomain.Play, error)
	AddSong(ctx context.Context, songId int, playListId int) (playListDomain.AddSong, error)
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
	playlist, err := s.playListStorage.StartPlayList(ctx, playListId)
	if err != nil {
		return playListDomain.Play{}, err
	}
	return playlist, err
}

func (s service) NextSong(ctx context.Context) (playListDomain.Play, error) {
	playlist, err := s.playListStorage.NextSong(ctx)
	if err != nil {
		return playListDomain.Play{}, err
	}
	return playlist, err
}

func (s service) PrevSong(ctx context.Context) (playListDomain.Play, error) {
	playlist, err := s.playListStorage.PrevSong(ctx)
	if err != nil {
		return playListDomain.Play{}, err
	}
	return playlist, err
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
