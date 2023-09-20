package main

import (
	"fmt"
	"log"
	"time"

	"github.com/eiannone/keyboard"
)

func ping(finish chan bool, wait chan bool) {

	ch := make(chan bool)
	time.Sleep(time.Second)
	fmt.Println("ping")
	close(wait)

	go pong(finish, ch)
	<-ch
}

func pong(finish chan bool, wait chan bool) {

	ch := make(chan bool)
	time.Sleep(time.Second)
	fmt.Println("pong")
	close(wait)

	go ping(finish, ch)
	<-ch
}

func WaitTillInput(finish chan bool) {
	if err := keyboard.Open(); err != nil {
		return
	}
	defer keyboard.Close()
	for {
		ch, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Println(err)
			panic("Vse ochen ploxo!!!!")
		} else {
			if ch == 'q' {
				finish <- true
				return
			}
		}

	}

}

func main() {
	finish := make(chan bool)
	wait := make(chan bool)
	go ping(finish, wait)
	go WaitTillInput(finish)
	<-finish
	<-wait
}
