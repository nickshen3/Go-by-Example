// Basic sends and receives on channels are blocking.
// However, we can use `select` with a `default` clause to
// implement _non-blocking_ sends, receives, and even
// non-blocking multi-way `select`s.
/*
常规的通过通道发送和接收数据是阻塞的。
然而，我们可以 使用带一个 default 子句的 select 来实现非阻塞 的 发送、接收，
甚至是非阻塞的多路 select。
*/

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Here's a non-blocking receive. If a value is
	// available on `messages` then `select` will take
	// the `<-messages` `case` with that value. If not
	// it will immediately take the `default` case.
	/*
	这里是一个非阻塞接收的例子。如果在 messages 中 存在，
	然后 select 将这个值带入 <-messages case 中。
	如果不是，就直接到 default 分支中。
	*/
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// A non-blocking send works similarly. Here `msg`
	// cannot be sent to the `messages` channel, because
	// the channel has no buffer and there is no receiver.
	// Therefore the `default` case is selected.
	//一个非阻塞发送的实现方法和上面一样。
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// We can use multiple `case`s above the `default`
	// clause to implement a multi-way non-blocking
	// select. Here we attempt non-blocking receives
	// on both `messages` and `signals`.
	/*
	我们可以在 default 前使用多个 case 子句来实现 一个多路的非阻塞的选择器。
	这里我们试图在 messages 和 signals 上同时使用非阻塞的接收操作。
	*/
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
