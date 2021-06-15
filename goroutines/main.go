package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg := &errgroup.Group{}

	eg.Go(func() error {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
		return nil
	})

	eg.Go(func() error {
		time.Sleep(1 * time.Second)
		return errors.New("error cause of something wrong")
	})

	fmt.Println(eg.Wait().Error())
}
