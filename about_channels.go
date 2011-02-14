package about_channels

import "./koans"
import "fmt"

func TestEveryChannelReceiveValue(t *koans.T) {
	ch := make(chan int, 100)
	wh := make(chan int)
	go func () {
		for {
			select {
			case value := <-ch:
				fmt.Printf("Thread 1 : %v\n", value)
			default:
				break
			}
		}
		fmt.Printf("Thread 1 : exit\n")
		wh <- 1
	}()
	go func () {
		for {
			select {
			case value := <-ch:
				fmt.Printf("Thread 2 : %v\n", value)
			default:
				break
			}
		}
		fmt.Printf("Thread 2 : exit\n")
		wh <- 2
	}()
	for i := 0; i < 100; i++ {
		ch <- i
	}
	<-wh
	<-wh

}
