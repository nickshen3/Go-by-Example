// Sometimes our Go programs need to spawn other, non-Go
// processes. For example, the syntax highlighting on this
// site is [implemented](https://github.com/mmcgrana/gobyexample/blob/master/tools/generate.go)
// by spawning a [`pygmentize`](http://pygments.org/)
// process from a Go program. Let's look at a few examples
// of spawning processes from Go.
//有时，我们的 Go 程序需要生成其他的，非 Go 进程。
//例如，这个 网站的语法高亮是通过在 Go 程序中生成一个 pygmentize 来实现的。 
//让我们看一些关于 Go 生成进程的例子。
package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {

	// We'll start with a simple command that takes no
	// arguments or input and just prints something to
	// stdout. The `exec.Command` helper creates an object
	// to represent this external process.
	//我们将从一个简单的命令开始，没有参数或者输入，仅打印 一些信息到标准输出流。
	//exec.Command 函数帮助我们创 建一个表示这个外部进程的对象。
	dateCmd := exec.Command("date")

	// `.Output` is another helper that handles the common
	// case of running a command, waiting for it to finish,
	// and collecting its output. If there were no errors,
	// `dateOut` will hold bytes with the date info.
	//.Output 是另一个帮助我们处理运行一个命令的常见情况 的函数，它等待命令运行完成，
	//并收集命令的输出。如果没 有出错，dateOut 将获取到日期信息的字节。
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Next we'll look at a slightly more involved case
	// where we pipe data to the external process on its
	// `stdin` and collect the results from its `stdout`.
	//下面我们将看看一个稍复杂的例子，我们将从外部进程的 stdin 输入数据并从 stdout 收集结果。
	grepCmd := exec.Command("grep", "hello")

	// Here we explicitly grab input/output pipes, start
	// the process, write some input to it, read the
	// resulting output, and finally wait for the process
	// to exit.
	//这里我们明确的获取输入/输出管道，运行这个进程，写入 一些输入信息，读取输出的结果，最后等待程序运行结束。
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	// We ommited error checks in the above example, but
	// you could use the usual `if err != nil` pattern for
	// all of them. We also only collect the `StdoutPipe`
	// results, but you could collect the `StderrPipe` in
	// exactly the same way.
	//上面的例子中，我们忽略了错误检测，但是你可以使用 if err != nil 的方式来进行错误检查，
	//我们也只收集 StdoutPipe 的结果，但是你可以使用相同的方法收集 StderrPipe 的结果。
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Note that when spawning commands we need to
	// provide an explicitly delineated command and
	// argument array, vs. being able to just pass in one
	// command-line string. If you want to spawn a full
	// command with a string, you can use `bash`'s `-c`
	// option:
	//注意，在生成命令时，我们需要提供显式描述的命令和参数 数组，而不能只传递一个命令行字符串。
	//如果你想使用一个 字符串生成一个完整的命令，那么你可以使用 bash 命令 的 -c 选项：
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
