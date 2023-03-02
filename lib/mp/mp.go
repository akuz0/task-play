package mp

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type MPPlayer struct {
	stat      int
	progress  int
	musicId   string
	isPlaying bool
	mux       sync.Mutex
}

// Possible worker states.
const (
	Stopped = 0
	Paused  = 1
	Running = 2
)

func (p *MPPlayer) Pause(status chan<- int) {
	if p.isPlaying {
		p.stat = 1
		status <- 1
	}
}

func (p *MPPlayer) Next(status chan<- int) string {
	if p.isPlaying {
		p.stat = 0
		status <- 0
	}
	return p.musicId
}

func (p *MPPlayer) Play(id string, name string, duration string, ctx context.Context, status chan int) {
	if id != p.musicId {
		i, err := strconv.Atoi(duration)
		if err != nil {
			return
		}
		p.progress = i

	}
	for {
		select {
		case state := <-status:
			switch state {
			case Stopped:
				fmt.Printf("Worker %d: Stopped\n", 1)
				return
			case Running:
				p.musicId = id
				fmt.Printf("Worker %d: Running\n", 1)
			case Paused:
				fmt.Printf("Worker %d: Paused\n", 1)
				return
			}

		default:
			time.Sleep(10 * time.Second)
			fmt.Println(". ", p.progress)
			inc(&p.mux, &p.progress)
			if p.progress == 0 {
				fmt.Printf("Music was played\n")
				return
			}
		}
	}
}

// Увеличение переменной на единицу.
// Функция принимает указатель, чтобы
// изменять оригинальную переменную.
func inc(mux *sync.Mutex, n *int) {
	// Блокировка мьютекса.
	mux.Lock()
	// Снятие блокировки мьютекса.
	defer mux.Unlock()

	// После того, как мьютекс заблокирован,
	// выполняем изменение.
	*n = *n - 1
}
