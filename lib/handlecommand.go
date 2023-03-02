package main1

import (
	"context"
	"fmt"
	lib2 "homework/lib/manager"
	"homework/lib/model"
	"homework/lib/mp"
	"strconv"
)

var p mp.MP3Player

func Start1(ctx context.Context, entity *model.MusicEntry) {
	//status = make(chan int, 1)
	//mp.Play(entity, ctx, status)
}

var start, pause, play, quit, wait = p.Prepare()
var musicMananger1 = lib2.NewMusuicManager()
var id1 = 0

func Start() {
	e := musicMananger1.Find("test1")
	if e == nil {
		fmt.Println("Трек не найден")
		return
	}
	p.Start()
}

func Add() {
	musicMananger1.Add(&model.MusicEntry{
		Id:       strconv.Itoa(id1),
		Name:     "test1",
		Duration: "60",
	})
	id1++
}

func Pause() {
	pause()
	//status = make(chan int, 1)
	//mp.Pause(status)
}
