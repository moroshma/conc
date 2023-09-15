package main

import (
	"fmt"
	"time"
)

func ping(finish chan bool, wait chan bool, n int) {

	ch := make(chan bool)
	time.Sleep(time.Second)
	fmt.Println("ping")
	close(wait)
	if n == 10 {
		finish <- true
	}

	n++
	go pong(finish, ch, n)
	<-ch

}

func pong(finish chan bool, wait chan bool, n int) {

	ch := make(chan bool)
	time.Sleep(time.Second)
	fmt.Println("pong")
	close(wait)
	if n == 10 {
		finish <- true
	}
	n++
	go ping(finish, ch, n)
	<-ch

}

/*
func getKeyTimeout(tm time.Duration) (ch rune, err error) {
	if err = keyboard.Open(); err != nil {
		return
	}
	defer keyboard.Close()

	var (
		chChan  = make(chan rune, 1)
		errChan = make(chan error, 1)

		timer = time.NewTimer(tm)
	)
	defer timer.Stop()

	go func(chChan chan<- rune, errChan chan<- error) {
		ch, _, err := keyboard.GetSingleKey()
		if err != nil {
			errChan <- err
			return
		}
		chChan <- ch
	}(chChan, errChan)

	select {
	case <-timer.C:
		return 0, nil
	case ch = <-chChan:
	case err = <-errChan:
	}

	return
}*/

func main() {
	finish := make(chan bool)
	wait := make(chan bool)
	go ping(finish, wait, 0)
	fmt.Println(<-finish, <-wait)
}
