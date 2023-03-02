package mp

import (
	"fmt"
	"sync"
	"time"
)

type MP3Player struct {
	chWork       <-chan struct{}
	chWorkBackup <-chan struct{}
	chControl    chan struct{}
	wg           sync.WaitGroup
}

func (p *MP3Player) Start() {
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

	i := 0
	for {
		select {
		case <-p.chWork:
			fmt.Println(i)
			i++
			time.Sleep(250 * time.Millisecond)
		case _, ok := <-p.chControl:
			if ok {
				continue
			}
			return
		}
	}
}

func (p *MP3Player) Prepare() (start, pause, play, quit, wait func()) {
	//start = func() {
	//	// chWork, chWorkBackup
	//	ch := make(chan struct{})
	//	close(ch)
	//	p.chWork = ch
	//	p.chWorkBackup = ch
	//
	//	// chControl
	//	p.chControl = make(chan struct{})
	//
	//	// wg
	//	p.wg = sync.WaitGroup{}
	//	p.wg.Add(1)
	//
	//	go routine()
	//}

	pause = func() {
		p.chWork = nil
		p.chControl <- struct{}{}
		fmt.Println("pause")
	}

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
