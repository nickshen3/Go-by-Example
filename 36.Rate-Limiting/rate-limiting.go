// <em>[Rate limiting](http://en.wikipedia.org/wiki/Rate_limiting)</em>
// is an important mechanism for controlling resource
// utilization and maintaining quality of service. Go
// elegantly supports rate limiting with goroutines,
// channels, and [tickers](tickers).
/*
速率限制(英) 是 一个重要的控制服务资源利用和质量的途径。
Go 通过 Go 协程、通道和打点器优美的支持了速率限制。
*/
package main

import "time"
import "fmt"

func main() {

	// First we'll look at basic rate limiting. Suppose
	// we want to limit our handling of incoming requests.
	// We'll serve these requests off a channel of the
	// same name.
	/*
	首先我们将看一下基本的速率限制。假设我们想限制我们 接收请求的处理，
	我们将这些请求发送给一个相同的通道。
	*/
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// This `limiter` channel will receive a value
	// every 200 milliseconds. This is the regulator in
	// our rate limiting scheme.
	//这个 limiter 通道将每 200ms 接收一个值。这个是 速率限制任务中的管理器。
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on a receive from the `limiter` channel
	// before serving each request, we limit ourselves to
	// 1 request every 200 milliseconds.
	//通过在每次请求前阻塞 limiter 通道的一个接收，我们限制 自己每 200ms 执行一次请求。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// We may want to allow short bursts of requests in
	// our rate limiting scheme while preserving the
	// overall rate limit. We can accomplish this by
	// buffering our limiter channel. This `burstyLimiter`
	// channel will allow bursts of up to 3 events.
	/*
	有时候我们想临时进行速率限制，并且不影响整体的速率控制， 我们可以通过通道缓冲来实现。 
	这个 burstyLimiter 通道用来进行 3 次临时的脉冲型速率限制。
	*/
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting.
	//想将通道填充需要临时改变3次的值，做好准备。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200 milliseconds we'll try to add a new
	// value to `burstyLimiter`, up to its limit of 3.
	//每 200 ms 我们将添加一个新的值到 burstyLimiter中， 直到达到 3 个的限制。
	go func() {
		//for t := range time.Tick(2000 * time.Millisecond) { 改为2000，效果更明显
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Now simulate 5 more incoming requests. The first
	// 3 of these will benefit from the burst capability
	// of `burstyLimiter`.
	//现在模拟超过 5 个的接入请求。它们中刚开始的 3 个将 由于受 burstyLimiter 的“脉冲”影响。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
