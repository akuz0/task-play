package mp

import (
	"context"
	"homework/lib/model"
)

type Player interface {
	Play(id string, name string, duration string, ctx context.Context, status chan int)
	Pause(status chan<- int)
	Next(status chan<- int) string
}

var p Player

func Play(e *model.MusicEntry, ctx context.Context, status chan int) {
	if p == nil {
		p = &MPPlayer{}
	}
	go p.Play(e.Id, e.Name, e.Duration, ctx, status)
}

func Next(e *model.MusicEntry, ctx context.Context, status chan int) {
	//id := p.Next(status)

	if p == nil {
		p = &MPPlayer{}
	}

	go p.Play(e.Id, e.Name, e.Duration, ctx, status)
}

func Prev(e *model.MusicEntry, ctx context.Context, status chan int) {
	//id := p.Prev(status)

	if p == nil {
		p = &MPPlayer{}
	}

	go p.Play(e.Id, e.Name, e.Duration, ctx, status)
}

func Pause(status chan int) {
	p.Pause(status)
}
