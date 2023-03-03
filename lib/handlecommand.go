package main1

import (
	"fmt"
	lib2 "homework/lib/manager"
	"homework/lib/model"
	"homework/lib/mp"
	"math/rand"
)

var p = mp.MP3Player{}

var start, pause, play, quit, wait = p.Prepare()
var musicMananger1 = lib2.NewMusuicManager()
var id1 = 0

func Start(name string) {
	entity := musicMananger1.Find(name)
	if entity == nil {
		fmt.Println("Трек не найден")
		return
	}
	p.Start(entity)
}

func Add(id int, name string) {
	model1 := &model.MusicEntry{
		Id:       id,
		Name:     name,
		Duration: rand.Intn(70-10) + 10, //todo - this is random DURATION
	}
	musicMananger1.Add(model1)
}

func Pause() {
	p.Pause()
}
