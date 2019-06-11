// A _goroutine_ is a lightweight thread of execution.
//Go 协程(goroutine) 在执行上来说是轻量级的线程。
package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Suppose we have a function call `f(s)`. Here's how
	// we'd call that in the usual way, running it
	// synchronously.
	//假设我们有一个函数叫做 f(s)。一般会这样同步(synchronously)调用
	f("direct")

	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.
	//使用 go f(s) 在一个 Go 协程中调用这个函数。 这个新的 Go 协程将会并发(concurrently)执行这个函数。
	go f("goroutine")

	// You can also start a goroutine for an anonymous
	// function call.
	//你也可以为匿名函数启动一个 Go 协程。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in
	// separate goroutines now, so execution falls through
	// to here. This `Scanln` requires we press a key
	// before the program exits.
	/*
	现在这两个 Go 协程在独立的 Go 协程中异步(asynchronously)运行，
	所以 程序直接运行到这一行。这里的 Scanln 代码需要我们 在程序退出前按下任意键结束。
	*/
	fmt.Scanln()
	fmt.Println("done")
}
