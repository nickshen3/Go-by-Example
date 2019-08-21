// Writing a basic HTTP server is easy using the
// `net/http` package.
//使用 net / http 包可以轻松编写基本的 HTTP 服务器。
package main

import (
	"fmt"
	"net/http"
)

// A fundamental concept in `net/http` servers is
// *handlers*. A handler is an object implementing the
// `http.Handler` interface. A common way to write
// a handler is by using the `http.HandlerFunc` adapter
// on functions with the appropriate signature.
//net / http服务器中的一个基本概念是处理程序。处理程序是实现 http.Handler接口的对象。
//写 入处理程序的常用方法是使用带有适当签名的 http.HandlerFunc adapteron 函数。
func hello(w http.ResponseWriter, req *http.Request) {

	// Functions serving as handlers take a
	// `http.ResponseWriter` and a `http.Request` as
	// arguments. The response writer is used to fill in the
	// HTTP response. Here out simple response is just
	// "hello\n".
	//作为处理程序的函数采用 http:ResponseWriter 和 http.Request asarguments。
	//响应编写器用于 填写 HTTP 响应。这里简单的回应就是 “hellon’’
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// This handler does something a little more
	// sophisticated by reading all the HTTP request
	// headers and echoing them into the response body.
	//这个处理程序通过读取所有 HTTP 请求标头并将它们回显到响应主体来做一些更复杂的事情。
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// We register our handlers on server routes using the
	// `http.HandleFunc` convenience function. It sets up
	// the *default router* in the `net/http` package and
	// takes a function as an argument.
	//我们使用 http.HandleFunc便捷函数在服务器路由上注册我们的处理程序。
	//它在 net / http包中 设置默认路由器并将一个函数作为参数。
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Finally, we call the `ListenAndServe` with the port
	// and a handler. `nil` tells it to use the default
	// router we've just set up.
	//最后，我们将 “ListenAndServe’’ 称为 portland 处理程序。nil 告诉它使用我们刚设置的默认路由器。
	http.ListenAndServe(":8090", nil)
}
