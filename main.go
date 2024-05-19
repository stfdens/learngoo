package main

import (
	"fmt"
	"sync"
)

func kirimdata(ch chan int, message int) {
	ch <- message
}

func terimadata(ch chan int) {
	message := <-ch
	fmt.Println(message)
}

func main() {
	ch := make(chan int) // Saluran berbuffer ukuran 10
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			kirimdata(ch, i)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		go terimadata(ch)
	}

}
