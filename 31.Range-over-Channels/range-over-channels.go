// In a [previous](range) example we saw how `for` and
// `range` provide iteration over basic data structures.
// We can also use this syntax to iterate over
// values received from a channel.
/*
在前面的例子中，我们讲过 for 和 range 为基本的数据结构提供了迭代的功能。
我们也可以使用这个语法 来遍历从通道中取得的值。
*/
package main

import "fmt"

func main() {

	// We'll iterate over 2 values in the `queue` channel.
	//我们将遍历在 queue 通道中的两个值。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This `range` iterates over each element as it's
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 2 elements.
	/*
	这个 range 迭代从 queue 中得到的每个值。
	因为我们 在前面 close 了这个通道，这个迭代会在接收完 2 个值 之后结束。
	如果我们没有 close 它，我们将在这个循环中 继续阻塞执行，等待接收第三个值
	*/
	for elem := range queue {
		fmt.Println(elem)
	}
}
///这个例子也让我们看到，一个非空的通道也是可以关闭的， 但是通道中剩下的值仍然可以被接收到。