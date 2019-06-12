// The primary mechanism for managing state in Go is
// communication over channels. We saw this for example
// with [worker pools](worker-pools). There are a few other
// options for managing state though. Here we'll
// look at using the `sync/atomic` package for _atomic
// counters_ accessed by multiple goroutines.
/*
Go 中最主要的状态管理方式是通过通道间的沟通来完成的，我们 在工作池的例子中碰到过，
但是还是有一些其他的方法来管理状态的。
这里我们将看看如何使用 sync/atomic 包在多个 Go 协程中进行 原子计数 。
*/
package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

	// We'll use an unsigned integer to represent our
	// (always-positive) counter.
	//我们将使用一个无符号整型数来表示（永远是正整数）这个计数器。
	var ops uint64

	// To simulate concurrent updates, we'll start 50
	// goroutines that each increment the counter about
	// once a millisecond.
	/*
	为了模拟并发更新，我们启动 50 个 Go 协程，对计数 器每隔 1ms （译者注：应为非准确时间）进行一次加一操作。
	*/
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// To atomically increment the counter we
				// use `AddUint64`, giving it the memory
				// address of our `ops` counter with the
				// `&` syntax.
				//使用 AddUint64 来让计数器自动增加，使用 & 语法来给出 ops 的内存地址。
				atomic.AddUint64(&ops, 1)

				// Wait a bit between increments.
				//在增加的时候等一下
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Wait a second to allow some ops to accumulate.
	//等待一秒，让 ops 的自加操作执行一会
	time.Sleep(time.Second)

	// In order to safely use the counter while it's still
	// being updated by other goroutines, we extract a
	// copy of the current value into `opsFinal` via
	// `LoadUint64`. As above we need to give this
	// function the memory address `&ops` from which to
	// fetch the value.
	/*
	为了在计数器还在被其它 Go 协程更新时，安全的使用它， 
	我们通过 LoadUint64 将当前值的拷贝提取到 opsFinal 中。
	和上面一样，我们需要给这个函数所取值的内存地址 &ops
	*/
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
