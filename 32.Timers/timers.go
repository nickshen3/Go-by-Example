// We often want to execute Go code at some point in the
// future, or repeatedly at some interval. Go's built-in
// _timer_ and _ticker_ features make both of these tasks
// easy. We'll look first at timers and then
// at [tickers](tickers).
/*
我们常常需要在后面一个时刻运行 Go 代码，或者在某段时间 间隔内重复运行。
Go 的内置 定时器 和 打点器 特性让这 些很容易实现。我们将先学习定时器，然后再学习打点器。
*/
package main

import "time"
import "fmt"

func main() {

	// Timers represent a single event in the future. You
	// tell the timer how long you want to wait, and it
	// provides a channel that will be notified at that
	// time. This timer will wait 2 seconds.
	/*
	定时器表示在未来某一时刻的独立事件。你告诉定时器 需要等待的时间，然后它将提供一个用于通知的通道。 
	这里的定时器将等待 2 秒。
	*/
	timer1 := time.NewTimer(2 * time.Second)

	// The `<-timer1.C` blocks on the timer's channel `C`
	// until it sends a value indicating that the timer
	// expired.
	//<-timer1.C 直到这个定时器的通道 C 明确的发送了 定时器失效的值之前，将一直阻塞。
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// If you just wanted to wait, you could have used
	// `time.Sleep`. One reason a timer may be useful is
	// that you can cancel the timer before it expires.
	// Here's an example of that.
	/*
	如果你需要的仅仅是单纯的等待，你需要使用 time.Sleep。 
	定时器是有用原因之一就是你可以在定时器失效之前，取消这个 定时器。这是一个例子
	*/
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
///第一个定时器将在程序开始后 ~2s 失效，但是第二个在它 没失效之前就停止了。