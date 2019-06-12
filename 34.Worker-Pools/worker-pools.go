// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.
//在这个例子中，我们将看到如何使用 Go 协程和通道 实现一个工作池 。
package main

import "fmt"
import "time"

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to
// simulate an expensive task.
/*
这是我们将要在多个并发实例中支持的任务了。这些执行者 将从 jobs 通道接收任务，
并且通过 results 发送对应 的结果。我们将让每个任务间隔 1s 来模仿一个耗时的任务。
*/
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	//为了使用 worker 工作池并且收集他们的结果，我们需要 2 个通道。
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	//这里启动了 3 个 worker，初始是阻塞的，因为 还没有传递任务。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	//这里我们发送 5 个 jobs，然后 close 这些通道 来表示这些就是所有的任务了。
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work.
	// This also ensures that the worker goroutines have
	// finished. An alternative way to wait for multiple
	// goroutines is to use a [WaitGroup](waitgroups).
	//最后，我们收集所有这些任务的返回值。
	for a := 1; a <= 5; a++ {
		<-results
	}
}
/*
执行这个程序，显示 5 个任务被多个 worker 执行。
整个程序 处理所有的任务仅执行了 3s 而不是 9s，是因为 3 个 worker 是并行的。
*/