package mp

import (
	"context"
	"sync"
	"testing"
)

type mockContext struct{}

func (ctx mockContext) Done() <-chan struct{} {
	ch := make(chan struct{})
	close(ch)
	return ch
}

func TestMPPlayer_Play(t *testing.T) {
	type fields struct {
		stat      int
		progress  int
		musicId   string
		isPlaying bool
		mux       sync.Mutex
	}
	type args struct {
		id       string
		name     string
		duration string
		ctx      context.Context
		status   chan int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		//{
		//	name: "test1",
		//	fields: fields{
		//		stat: 0,
		//		isPlaying: false,
		//		musicId: "0",
		//		progress: 0,
		//		mux:
		//	},
		//	args: args{id: "0", name: "test1"},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &MPPlayer{
				stat:      tt.fields.stat,
				progress:  tt.fields.progress,
				musicId:   tt.fields.musicId,
				isPlaying: tt.fields.isPlaying,
				mux:       tt.fields.mux,
			}
			p.Play(tt.args.id, tt.args.name, tt.args.duration, tt.args.ctx, tt.args.status)
		})
	}
}
