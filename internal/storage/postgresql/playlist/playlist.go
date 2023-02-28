package playlist

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	playListDomain "homework/internal/domain/playlist"
)

type PlayListStorage interface {
	GetPlayListByID(ctx context.Context, playListId int) (playListDomain.Play, error)
	CreatePlayList(ctx context.Context, playlist playListDomain.Play) (playListDomain.Play, error)
	StartPlayList(ctx context.Context, playListId int) (playListDomain.Play, error)
	NextSong(ctx context.Context)
	PrevSong(ctx context.Context)
	AddSong(ctx context.Context, songId int, playListId int) (playListDomain.AddSong, error)
}

type playListStorage struct {
	DB *pgxpool.Pool
}

func (p playListStorage) GetPlayListByID(ctx context.Context, playListId int) (playListDomain.Play, error) {
	//TODO implement me
	panic("implement me")
}

func (p playListStorage) CreatePlayList(ctx context.Context, playlist playListDomain.Play) (playListDomain.Play, error) {
	//TODO implement me
	panic("implement me")
}

func (p playListStorage) StartPlayList(ctx context.Context, playListId int) (playListDomain.Play, error) {
	//TODO implement me
	panic("implement me")
}

func (p playListStorage) NextSong(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (p playListStorage) PrevSong(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (p playListStorage) AddSong(ctx context.Context, songId int, playListId int) (playListDomain.AddSong, error) {
	//TODO implement me
	panic("implement me")
}

func NewStorage(DB *pgxpool.Pool) PlayListStorage {
	return &playListStorage{DB: DB}
}
