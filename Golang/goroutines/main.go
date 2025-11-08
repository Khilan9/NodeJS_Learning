package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}

func getStatusCode(endpoint string, wg *sync.WaitGroup, mut *sync.Mutex) {

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
	defer wg.Done()
}

func goexamplewithlock() {
	wg2 := &sync.WaitGroup{}
	mut2 := &sync.RWMutex{}

	var score = []int{0}

	wg2.Add(4) // Correct count for 4 goroutines

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg2, mut2)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg2, mut2)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg2, mut2)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Read scores")
		m.RLock()
		fmt.Println(score)
		m.RUnlock()
		wg.Done()
	}(wg2, mut2)

	wg2.Wait()
	fmt.Println("Score at end", score)
}

func simplegowaitgroup() {
	// First part - website status check
	var wg sync.WaitGroup
	var mut sync.Mutex

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		wg.Add(1) // Add before launching goroutine
		go getStatusCode(web, &wg, &mut)
	}

	wg.Wait()
	fmt.Println(signals)
}

func goexamplechannel() {
	messages := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Channel blocked for 5 seconds as no one is listening")
		messages <- "ping"
		wg.Done()
	}()
	go func() {
		fmt.Println("Channel will start listening after 5 seconds")
		time.Sleep(5 * time.Second)
		msg, ischannelhavejob := <-messages
		if ischannelhavejob {
			fmt.Println("channel have some jobs")
			fmt.Println(msg)
		}
		wg.Done()
	}()
	wg.Wait()
}

func goexamplebuffer() {
	// here 2 indicates buffer size
	// sending operations won't block for 2 sender until the buffer is full
	messages := make(chan string, 2)

	// fmt.Println(<-messages) this will block as no one is sending data into channel
	messages <- "buffered"
	messages <- "channel"
	// messages <- "channel2" this will block as buffer size is 2

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func gochanneldir() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// here pings is used to send data
	go func(pings chan<- string, msg string) {
		pings <- msg
	}(pings, "passed message")

	// here pings is used to recieve data while pongs is used to send data to this channel
	go func(pings <-chan string, pongs chan<- string) {
		msg := <-pings
		pongs <- msg
	}(pings, pongs)

	fmt.Println(<-pongs)
}

func closechannelexample() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		// time.Sleep(3 * time.Second)
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)

}

func queuechannel() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

}

func timerexample() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

func tickerexample() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	// done <- true
	fmt.Println("Ticker stopped")
}

func atomiccnter() {
	var ops uint64 = 0
	// var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				// mutex.Lock()
				ops++
				// mutex.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops)
}
func main() {
	// simplegowaitgroup()
	// fmt.Println("Race condition - LearnCodeonline.in")
	// goexamplewithlock()
	// goexamplebuffer()
	// goexamplechannel()
	// gochanneldir()
	// closechannelexample()
	// queuechannel()
	// timerexample()
	// tickerexample()
	atomiccnter()
}
