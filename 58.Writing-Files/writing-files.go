// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.
//Go 写文件和我们前面看过的读操作有着相似的方式。
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// To start, here's how to dump a string (or just
	// bytes) into a file.
	//开始，这里是展示如写入一个字符串（或者只是一些字节）到一个文件。
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// For more granular writes, open a file for writing.
	//对于更细粒度的写入，先打开一个文件。
	f, err := os.Create("/tmp/dat2")
	check(err)

	// It's idiomatic to defer a `Close` immediately
	// after opening a file.
	//打开文件后，习惯立即使用 defer 调用文件的 Close操作
	defer f.Close()

	// You can `Write` byte slices as you'd expect.
	//你可以写入你想写入的字节切片
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// A `WriteString` is also available.
	//WriteString 也是可用的。
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a `Sync` to flush writes to stable storage.
	//调用 Sync 来将缓冲区的信息写入磁盘。
	f.Sync()

	// `bufio` provides buffered writers in addition
	// to the buffered readers we saw earlier.
	//bufio 提供了和我们前面看到的带缓冲的读取器一样的带缓冲的写入器。
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// Use `Flush` to ensure all buffered operations have
	// been applied to the underlying writer.
	//使用 Flush 来确保所有缓存的操作已写入底层写入器
	w.Flush()

}
