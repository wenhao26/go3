package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Service struct {
	ch chan bool
	wg *sync.WaitGroup
}

func New() *Service {
	return &Service{
		ch: make(chan bool),
		wg: &sync.WaitGroup{},
	}
}

func (s *Service) Stop() {
	close(s.ch)
	s.wg.Wait()
}

func (s *Service) AddTask() {
	s.wg.Add(1)
	defer s.wg.Done()

	for {
		select {
		case <-s.ch:
			fmt.Println("Stop!!!")
			return
		default:

		}

		time.Sleep(time.Second * 2)
		fmt.Println("Add..." + time.Now().Format("2006.01.02 15:04:05"))
	}
}

func main() {
	s := New()
	go s.AddTask()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-ch)

	s.Stop()
}
