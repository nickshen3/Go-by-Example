// Use `os.Exit` to immediately exit with a given
// status.
//使用 os.Exit 来立即进行带给定状态的退出。
package main

import "fmt"
import "os"

func main() {

	// `defer`s will _not_ be run when using `os.Exit`, so
	// this `fmt.Println` will never be called.
	//当使用 os.Exit 时 defer 将不会 执行，所以这里的 fmt.Println 将永远不会被调用。
	defer fmt.Println("!")

	// Exit with status 3.
	//退出并且退出状态为 3。
	os.Exit(3)
}

// Note that unlike e.g. C, Go does not use an integer
// return value from `main` to indicate exit status. If
// you'd like to exit with a non-zero status you should
// use `os.Exit`.
/*
注意，不像例如 C 语言，Go 不使用在 main 中返回一个整 数来指明退出状态。
如果你想以非零状态退出，那么你就要 使用 os.Exit。
*/
