package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

var (
	ErrCh chan error
	comp bool
)

func init() {
	ErrCh = make(chan error, 1)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	in, errCh := fetch(ctx)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutineおわた")
				return
			case e := <- errCh:
				ErrCh <- e
				return
			}
		}
	}()

	go func() {
		for v := range in {
			fmt.Printf("%#v\n", v)
			time.Sleep(time.Second)
		}
		cancel()
		comp = true
	}()

	wait()
}

func wait() {
	for {
		select {
		case err := <-ErrCh:
			fmt.Printf("err: %#v\n", err)
			return
		default:
			if comp {
				return
			}
		}
	}
}

func fetch(ctx context.Context) (<- chan int, <- chan error) {
	in := make(chan int, 10)
	errCh := make(chan error, 1)

	go func() {
		defer func() {
			close(in)
			close(errCh)
		}()

		ints := []int{1,2,3,4,5}
		for _, v := range ints {
			in <- v
		}
		time.Sleep(3 * time.Second)

		errCh <- errors.New("hoge")
	}()

	return in, errCh
}
