package about_channels

import "./koans"

func TestEveryChannelReceiveValue(t *koans.T) {
	count := 0
	ch := make(chan int, 100)
	wh := make(chan int)
	go func () {
		loop: for {
			select {
			case <-ch:
				count++
			default:
				break loop
			}
		}
		wh <- 1
	}()
	go func () {
		loop: for {
			select {
			case <-ch:
				count++
			default:
				break loop
			}
		}
		wh <- 2
	}()
	for i := 0; i < 100; i++ {
		ch <- i
	}
	<-wh
	<-wh
	t.AssertEquals(koans.Int__, count)

}
