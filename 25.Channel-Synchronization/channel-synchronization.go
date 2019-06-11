// We can use channels to synchronize execution
// across goroutines. Here's an example of using a
// blocking receive to wait for a goroutine to finish.
// When waiting for multiple goroutines to finish,
// you may prefer to use a [WaitGroup](waitgroups).
/*
我们可以使用通道来同步 Go 协程间的执行状态。这里是一个 使用阻塞的接受方式来等待一个 Go 协程的运行结束。
*/

package main

import "fmt"
import "time"

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.
/*
这是一个我们将要在 Go 协程中运行的函数。done 通道 将被用于通知其他 Go 协程这个函数已经工作完毕。
*/
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we're done.
	//发送一个值来通知我们已经完工啦。
	done <- true
}

func main() {

	// Start a worker goroutine, giving it the channel to
	// notify on.
	//运行一个 worker Go协程，并给予用于通知的通道。
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the
	// worker on the channel.
	//程序将在接收到通道中 worker 发出的通知前一直阻塞。
	<-done
}
