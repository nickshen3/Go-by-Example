// [Timers](timers) are for when you want to do
// something once in the future - _tickers_ are for when
// you want to do something repeatedly at regular
// intervals. Here's an example of a ticker that ticks
// periodically until we stop it.
/*
定时器 是当你想要在未来某一刻执行一次时 使用的 - 打点器 则是当你想要在固定的时间间隔重复执行 准备的。
这里是一个打点器的例子，它将定时的执行，直到我 们将它停止。
*/
package main

import "time"
import "fmt"

func main() {

	// Tickers use a similar mechanism to timers: a
	// channel that is sent values. Here we'll use the
	// `range` builtin on the channel to iterate over
	// the values as they arrive every 500ms.
	/*
	打点器和定时器的机制有点相似：一个通道用来发送数据。 
	这里我们在这个通道上使用内置的 range 来迭代值每隔 500ms 发送一次的值。
	*/
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// Tickers can be stopped like timers. Once a ticker
	// is stopped it won't receive any more values on its
	// channel. We'll stop ours after 1600ms.
	/*
	打点器可以和定时器一样被停止。一旦一个打点停止了， 
	将不能再从它的通道中接收到值。我们将在运行后 1600ms 停止这个打点器。
	*/
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
///当我们运行这个程序时，这个打点器会在我们停止它前打点3次。