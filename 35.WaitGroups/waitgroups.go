// To wait for multiple goroutines to finish, we can
// use a *wait group*.
//要等待多个 Go 协程完成，我们可以使用 wait group。

package main

import (
	"fmt"
	"sync"
	"time"
)

// This is the function we'll run in every goroutine.
// Note that a WaitGroup must be passed to functions by
// pointer.
//这是我们将在每个 Go 协程中运行的函数。请注意，WaitGroup 必须通过指针传递给函数
func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)

	// Sleep to simulate an expensive task.
	//等待 1s 来模拟一个耗时的任务。
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)

	// Notify the WaitGroup that this worker is done.
	//通知 WaitGroup 此执行已完成。
	wg.Done()
}

func main() {

	// This WaitGroup is used to wait for all the
	// goroutines launched here to finish.
	//此 WaitGroup 用于等待此处启动的所有 Go 协程完成。
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup
	// counter for each.
	//启动几个 Go 协程并为每个 Go 协程增加 WaitGroup 计数器。
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Block until the WaitGroup counter goes back to 0;
	// all the workers notified they're done.
	//在 WaitGroup 计数器返回 0 之前阻塞;所有工作人员都通知他们已完成。
	wg.Wait()
}
