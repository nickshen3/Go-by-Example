// In the previous example we looked at
// [spawning external processes](spawning-processes). We
// do this when we need an external process accessible to
// a running Go process. Sometimes we just want to
// completely replace the current Go process with another
// (perhaps non-Go) one. To do this we'll use Go's
// implementation of the classic
// <a href="http://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>
// function.
//在前面的例子中，我们了解了生成外部进程 的知识，当我们需要访问外部进程时需要这样做，
//但是有时候，我们只想 用其他的（也许是非 Go 程序）来完全替代当前的 Go 进程。
//这时候，我们 可以使用经典的 exec 方法的 Go 实现。
package main

import "syscall"
import "os"
import "os/exec"

func main() {

	// For our example we'll exec `ls`. Go requires an
	// absolute path to the binary we want to execute, so
	// we'll use `exec.LookPath` to find it (probably
	// `/bin/ls`).
	//在我们的例子中，我们将执行 ls 命令。Go 需要提供我 们需要执行的可执行文件的绝对路径，
	//所以我们将使用 exec.LookPath 来得到它（大概是 /bin/ls）
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec` requires arguments in slice form (as
	// apposed to one big string). We'll give `ls` a few
	// common arguments. Note that the first argument should
	// be the program name.
	//Exec 需要的参数是切片的形式的（不是放在一起的一个大字 符串）。
	//我们给 ls 一些基本的参数。注意，第一个参数需要 是程序名。
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` also needs a set of [environment variables](environment-variables)
	// to use. Here we just provide our current
	// environment.
	//Exec 同样需要使用环境变量。 这里我们仅提供当前的环境变量。
	env := os.Environ()

	// Here's the actual `syscall.Exec` call. If this call is
	// successful, the execution of our process will end
	// here and be replaced by the `/bin/ls -a -l -h`
	// process. If there is an error we'll get a return
	// value.
	//这里是 syscall.Exec 调用。如果这个调用成功，那么我们的 进程将在这里被替换成 /bin/ls -a -l -h 进程。
	//如果存 在错误，那么我们将会得到一个返回值。
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
