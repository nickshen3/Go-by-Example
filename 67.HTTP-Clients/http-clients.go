// The Go standard library comes with excellent support
// for HTTP clients and servers in the `net/http`
// package. In this example we'll use it to issue simple
// HTTP requests.
//Go标准库为 net / httppackage中的HTTP客户端和服务器提供了出色的支持。
//在这个例子中，我 们将使用它来发出简单的 HTTP 请求。
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Issue an HTTP GET request to a server. `http.Get` is a
	// convenient shortcut around creating an `http.Client`
	// object and calling its `Get` method; it uses the
	// `http.DefaultClient` object which has useful default
	// settings.
	//向服务器发出 HTTP GET 请求。http.Get 是创建 http.Clientobject 并调用其 Get 方法的简便快 捷方式; 
	//它使用 http.DefaultClient 对象，该对象具有有用的默认设置。
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	//打印 HTTP 响应状态。
	fmt.Println("Response status:", resp.Status)

	// Print the first 5 lines of the response body.
	//打印响应正文的前 5 行
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
