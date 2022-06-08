package main

import "C"
import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	cfg := net.ListenConfig{
		KeepAlive: time.Minute,
	}
	l, err := cfg.Listen(ctx, "tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	wg := &sync.WaitGroup{}
	log.Println("im started!")

	messagesFromServer := make(chan string, 100)
	go inputScanner(messagesFromServer)

	go func() {

		for {
			conn, err := l.Accept()
			if err != nil {
				log.Println(err)
			} else {
				wg.Add(1)
				go handleConn(ctx, conn, wg, messagesFromServer)

			}
		}
	}()
	<-ctx.Done()
	log.Println("done")
	l.Close()
	wg.Wait()
	log.Println("exit")

}

func handleConn(ctx context.Context, conn net.Conn, wg *sync.WaitGroup, messagesFromServer chan string) {
	defer wg.Done()
	defer conn.Close()

	tck := time.NewTicker(time.Second)

	for {

		select {

		case <-ctx.Done():
			return
		case t := <-tck.C:
			_, _ = fmt.Fprintf(conn, "now: %s\n", t)
		case mess := <-messagesFromServer:
			_, _ = fmt.Fprintf(conn, "now: %s\n", mess)
		}
	}
}

func inputScanner(messagesFromServer chan string) {

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		messagesFromServer <- input.Text()
	}
}
