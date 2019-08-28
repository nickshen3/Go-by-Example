// In the previous example we saw how to manage simple
// counter state using [atomic operations](atomic-counters).
// For more complex state we can use a <em>[mutex](http://en.wikipedia.org/wiki/Mutual_exclusion)</em>
// to safely access data across multiple goroutines.
/*
在前面的例子中，我们看到了如何使用原子操作来管理简单的计数器。
对于更加复杂的情况，我们可以使用一个互斥锁 来在 Go 协程间安全的访问数据。
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// For our example the `state` will be a map.
	//在我们的例子中，state 是一个 map。
	var state = make(map[int]int)

	// This `mutex` will synchronize access to `state`.
	//这里的 mutex 将同步对 state 的访问。
	var mutex = &sync.Mutex{}

	// We'll keep track of how many read and write
	// operations we do.
	//我们将跟踪我们执行的读写操作数量。
	var readOps uint64
	var writeOps uint64

	// Here we start 100 goroutines to execute repeated
	// reads against the state, once per millisecond in
	// each goroutine.
	//在这里，我们启动 100 个 goroutines 来执行针对状态的重复读取，每毫秒一次。
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				// For each read we pick a key to access,
				// `Lock()` the `mutex` to ensure
				// exclusive access to the `state`, read
				// the value at the chosen key,
				// `Unlock()` the mutex, and increment
				// the `readOps` count.
				/*
					每次循环读取，我们使用一个键来进行访问， Lock() 这个 mutex 来确保对 state 的 独占访问，
					读取选定的键的值，Unlock() 这个 mutex，并且 ops 值加 1。
				*/
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				// Wait a bit between reads.
				//在读取之间稍等一下。
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// We'll also start 10 goroutines to simulate writes,
	// using the same pattern we did for reads.
	//同样的，我们运行 10 个 Go 协程来模拟写入操作，使用 和读取相同的模式。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let the 10 goroutines work on the `state` and
	// `mutex` for a second.
	//让这 10 个 Go 协程对 state 和 mutex 的操作 运行 1 s。
	time.Sleep(time.Second)

	// Take and report final operation counts.
	//获取并输出最终的操作计数。
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// With a final lock of `state`, show how it ended up.
	//对 state 使用一个最终的锁，显示它是如何结束的。
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
