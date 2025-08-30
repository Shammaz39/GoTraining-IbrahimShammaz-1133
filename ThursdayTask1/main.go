package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {

	wg := new(sync.WaitGroup)
	quit := make(chan struct{})

	tickers := []string{"AAPL", "GOOG", "INFY"}

	go func() {

		wg.Add(1)
		go func() {
			for {
				select {
				case <-quit:
					fmt.Println("10 secs are completed")
					return
				default:
					time.Sleep(1 * time.Second)
					for _, s := range tickers {
						wg.Add(1)
						go getFeed(s, wg, quit)
					}
				}
			}

		}()
		wg.Done()
	}()

	time.Sleep(10 * time.Second)
	close(quit)
	wg.Wait()

}

func getFeed(stock string, wg *sync.WaitGroup, quit chan struct{}) {
	defer wg.Done()
	price := rand.Float64() * 100
	for {
		select {
		case <-quit:
			return
		case <-time.After(time.Second):
			timeStr := time.Now().Format("15:04:05")
			fmt.Printf("[%s] %s: %.2f\n", timeStr, stock, price)
		}
	}
}
