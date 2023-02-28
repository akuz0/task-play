package playlist

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	playListDomain "homework/internal/domain/playlist"
	"log"
)

type PlayListStorage interface {
	GetPlayListByID(ctx context.Context, playListId int) (playListDomain.Play, error)
	CreatePlayList(ctx context.Context, playListId int) (playListDomain.Play, error)
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

func (p playListStorage) CreatePlayList(ctx context.Context, playListId int) (playListDomain.Play, error) {
	//todo: add const
	row, err := p.DB.Query(ctx, "INSERT INTO playlist VALUES ($1, null, 'stop') RETURNING id", playListId)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var playList playListDomain.Play
	//err = row.Scan(&playListId)
	//if err != nil {
	//	log.Println(err)
	//}
	playList.Id = playListId
	return playList, err
}

func (p playListStorage) StartPlayList(ctx context.Context, playListId int) (playListDomain.Play, error) {

	row, err := p.DB.Query(ctx, "select id from playlist WHERE status = 'play'")

	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var id int
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		row, err = p.DB.Query(ctx, "UPDATE playlist SET status = 'stop' where id = $1", id)
		if err != nil {
			log.Fatal(err)
		}
		defer row.Close()
	}

	row, err = p.DB.Query(ctx, "UPDATE playlist SET status = 'play' where id = $1", playListId)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var playList playListDomain.Play
	playList.Id = playListId
	return playList, err
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
