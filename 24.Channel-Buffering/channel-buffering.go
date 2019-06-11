// By default channels are _unbuffered_, meaning that they
// will only accept sends (`chan <-`) if there is a
// corresponding receive (`<- chan`) ready to receive the
// sent value. _Buffered channels_ accept a limited
// number of  values without a corresponding receiver for
// those values.
/*
默认通道是 无缓冲 的，这意味着只有在对应的接收（<- chan） 通道准备好接收时，
才允许进行发送（chan <-）。可缓存通道 允许在没有对应接收方的情况下，缓存限定数量的值。
*/

package main

import "fmt"

func main() {

	// Here we `make` a channel of strings buffering up to
	// 2 values.
	//这里我们 make 了一个通道，最多允许缓存 2 个值。
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these
	// values into the channel without a corresponding
	// concurrent receive.
	//因为这个通道是有缓冲区的，即使没有一个对应的并发接收 方，我们仍然可以发送这些值。
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive these two values as usual.
	//然后我们可以像前面一样接收这两个值。
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
