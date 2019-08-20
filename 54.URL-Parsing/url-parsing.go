// URLs provide a [uniform way to locate resources](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// Here's how to parse URLs in Go.
//URL 提供了一个统一资源定位方式。这里了解一下 Go 中是如何解析 URL 的。
package main

import "fmt"
import "net"
import "net/url"

func main() {

	// We'll parse this example URL, which includes a
	// scheme, authentication info, host, port, path,
	// query params, and query fragment.
	//我们将解析这个 URL 示例，它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parse the URL and ensure there are no errors.
	//解析这个 URL 并确保解析没有出错。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Accessing the scheme is straightforward.
	//直接访问 scheme。
	fmt.Println(u.Scheme)

	// `User` contains all authentication info; call
	// `Username` and `Password` on this for individual
	// values.
	//User 包含了所有的认证信息，这里调用 Username和 Password 来获取独立值。
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// The `Host` contains both the hostname and the port,
	// if present. Use `SplitHostPort` to extract them.
	//Host 同时包括主机名和端口信息，如过端口存在的话，使用 strings.Split() 从 Host 中手动提取端口。
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Here we extract the `path` and the fragment after
	// the `#`.
	//这里我们提出路径和查询片段信息。
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// To get query params in a string of `k=v` format,
	// use `RawQuery`. You can also parse query params
	// into a map. The parsed query param maps are from
	// strings to slices of strings, so index into `[0]`
	// if you only want the first value.
	/*
	要得到字符串中的 k=v 这种格式的查询参数，可以使用 RawQuery 函数。
	你也可以将查询参数解析为一个map。已解析的查询参数 map 以查询字符串为键，
	对应值字符串切片为值，所以如何只想得到一个键对应的第一个值，将索引位置设置为 [0] 就行了。
	*/
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
/*
运行我们的 URL 解析程序，显示全部我们提取的 URL 的不同数据块。
*/