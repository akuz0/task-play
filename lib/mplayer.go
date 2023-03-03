package main1

import (
	"bufio"
	"context"
	"fmt"
	lib2 "homework/lib/manager"
	"homework/lib/model"
	"homework/lib/mp"
	"os"
	"strings"
)

var musicMananger *lib2.MusicManagerList
var id = 0
var ctrl, signal chan int

var number = 0

func handleLibCommands(commands []string) {
	switch commands[1] {
	case "list":
		for i := 0; i < musicMananger.Len(); i++ {
			musicEntry, _ := musicMananger.Get(i)
			fmt.Println(i+1, ":", musicEntry.Id, musicEntry.Name, musicEntry.Duration)
		}
	case "add":
		if len(commands) == 4 {

			musicMananger.Add(&model.MusicEntry{
				Id:       id,
				Name:     commands[2],
				Duration: 2,
			})
			id++
		} else {
			fmt.Println("Используйте: lib add <name><duration>")
		}
	case "remove":
		if len(commands) == 3 {
		} else {
			fmt.Println("Временно не работает. Используйте: lib remove <name>")
		}
	default:
		fmt.Println("Неизвестная команда:", commands[1])
	}
}

func handlePlay(commands []string, ctx context.Context, status chan int) {
	switch commands[0] {
	case "play":
		if len(commands) != 2 {
			fmt.Println("Используйте комманду: play <name>")
			return
		}
		e := musicMananger.Find(commands[1])
		if e == nil {
			fmt.Println("Трек", commands[1], "не найден")
			return
		}
		numberRes := e.Id
		number = numberRes
		mp.Play(e, ctx, status)

	case "pause":
		mp.Pause(status)

	case "next":
		e, err := musicMananger.Get(number)
		if err != nil {
			fmt.Println("Трек", commands[1], "не найден")
			return
		}
		mp.Next(e, ctx, status)

	case "prev":
		e, err := musicMananger.Perv(number)
		if err != nil {
			fmt.Println("Трек", commands[1], "не найден")
			return
		}
		mp.Prev(e, ctx, status)
	}
}

func main() {

	ctx := context.Background()
	// создаем дочерний контекст, который можно распостранить в нижележащине функции
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	status := make(chan int, 1)

	fmt.Println(`
Доступные команды:
lib list -- Список треков
lib add <name><duration> -- Добавить трек
play <name> -- Воспроизвести трек
`)
	musicMananger = lib2.NewMusuicManager()
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Ввод команды-> ")
		rawLine, _, _ := r.ReadLine()

		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}

		commands := strings.Split(line, " ")
		if commands[0] == "lib" {
			handleLibCommands(commands)
		} else if commands[0] == "play" {
			// lib add a1 a a MP3
			status <- 2
			handlePlay(commands, ctx, status)
		} else if commands[0] == "pause" {
			// lib add a1 a a MP3
			status <- 1
			handlePlay(commands, ctx, status)
		} else if commands[0] == "next" {
			// lib add a1 a a MP3
			status <- 2
			handlePlay(commands, ctx, status)
		} else if commands[0] == "prev" {
			// lib add a1 a a MP3
			status <- 2
			handlePlay(commands, ctx, status)
		} else {
			fmt.Println("Unrecognized command:", commands[0])
		}
	}
}
