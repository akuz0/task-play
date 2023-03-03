package mp

import (
	"fmt"
	"homework/lib/model"
	"sync"
	"time"
)

type MP3Player struct {
	musicId   int
	musicName string
	progress  int
	mux       sync.Mutex
	//isPlaying bool
	chWork       <-chan struct{}
	chWorkBackup <-chan struct{}
	chControl    chan struct{}
	wg           sync.WaitGroup
}

func (p *MP3Player) Start(entity *model.MusicEntry) {
	if entity.Name != p.musicName {
		i := entity.Duration
		p.progress = i
	}
	p.musicName = entity.Name
	ch := make(chan struct{})
	close(ch)
	p.chWork = ch
	p.chWorkBackup = ch

	// chControl
	p.chControl = make(chan struct{})

	// wg
	p.wg = sync.WaitGroup{}
	p.wg.Add(1)

	go p.routine()
}

func (p *MP3Player) routine() {
	defer p.wg.Done()

	for {
		select {
		case <-p.chWork:
			time.Sleep(10 * time.Second)
			fmt.Println(". ", p.progress)
			inc(&p.mux, &p.progress)
			if p.progress == 0 {
				fmt.Printf("Music was played\n")
				return
			}
		case _, ok := <-p.chControl:
			if ok {
				continue
			}
			return
		}
	}
}

func (p *MP3Player) Pause() {
	p.chWork = nil
	p.chControl <- struct{}{}
	fmt.Println("music pause")
}

func (p *MP3Player) Prepare() (start, pause, play, quit, wait func()) {
	play = func() {
		fmt.Println("play")
		p.chWork = p.chWorkBackup
		p.chControl <- struct{}{}
	}

	quit = func() {
		p.chWork = nil
		close(p.chControl)
		fmt.Println("quit")
	}

	wait = func() {
		p.wg.Wait()
	}

	return
}

func sleep() {
	time.Sleep(1 * time.Second)
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
