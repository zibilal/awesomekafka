package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
	"context"
)

func main() {
	fmt.Println("Starts...")

	done := make(chan struct{})

	count := func(done chan <- struct{}, ctx context.Context) {

		n := 0
		for {
			select {
			case <- ctx.Done():
				fmt.Println("Program is cancelled by outside process")
			default:
				fmt.Printf("-->%d\n", n)
				time.Sleep(1 * time.Second)
				if n == 20 {
					done <- struct{}{}
				}
				n++
			}
		}
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go count(done, ctx)

	timer := time.NewTimer(30 * time.Second)

	select {
	case <- signals:
		fmt.Println("Program is interrupted")
	case <- done:
		fmt.Println("Done")
	case <- timer.C:
		fmt.Println("Too long, cancelling the program")
		cancel()
	}

	fmt.Println("Ends")
}
