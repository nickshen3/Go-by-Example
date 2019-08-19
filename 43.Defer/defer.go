// _Defer_ is used to ensure that a function call is
// performed later in a program's execution, usually for
// purposes of cleanup. `defer` is often used where e.g.
// `ensure` and `finally` would be used in other languages.
/*
Defer 被用来确保一个函数调用在程序执行结束前执行。同样用来执行一些清理工作。 
defer 用在像其他语言中的ensure 和 finally用到的地方。
*/

package main

import "fmt"
import "os"

// Suppose we wanted to create a file, write to it,
// and then close when we're done. Here's how we could
// do that with `defer`.
/*
假设我们想要创建一个文件，向它进行写操作，然后在结束时关闭它。这里展示了如何通过 defer 来做到这一切。
*/
func main() {

	// Immediately after getting a file object with
	// `createFile`, we defer the closing of that file
	// with `closeFile`. This will be executed at the end
	// of the enclosing function (`main`), after
	// `writeFile` has finished.
	/*
	在 closeFile 后得到一个文件对象，我们使用 defer通过 closeFile 来关闭这个文件。
	这会在封闭函数（main）结束时执行，就是 writeFile 结束后。
	*/
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
/*
执行程序，确认这个文件在写入后是已关闭的。
*/