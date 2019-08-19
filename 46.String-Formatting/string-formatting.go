// Go offers excellent support for string formatting in
// the `printf` tradition. Here are some examples of
// common string formatting tasks.
//Go 在传统的printf 中对字符串格式化提供了优异的支持。这里是一些基本的字符串格式化的人物的例子。
package main

import "fmt"
import "os"

type point struct {
	x, y int
}

func main() {

	// Go offers several printing "verbs" designed to
	// format general Go values. For example, this prints
	// an instance of our `point` struct.
	//Go 为常规 Go 值的格式化设计提供了多种打印方式。例如，这里打印了 point 结构体的一个实例。
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// If the value is a struct, the `%+v` variant will
	// include the struct's field names.
	//如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
	fmt.Printf("%+v\n", p)

	// The `%#v` variant prints a Go syntax representation
	// of the value, i.e. the source code snippet that
	// would produce that value.
	//%#v 形式则输出这个值的 Go 语法表示。例如，值的运行源代码片段。
	fmt.Printf("%#v\n", p)

	// To print the type of a value, use `%T`.
	//需要打印值的类型，使用 %T。
	fmt.Printf("%T\n", p)

	// Formatting booleans is straight-forward.
	//格式化布尔值是简单的。
	fmt.Printf("%t\n", true)

	// There are many options for formatting integers.
	// Use `%d` for standard, base-10 formatting.
	//格式化整形数有多种方式，使用 %d进行标准的十进制格式化。
	fmt.Printf("%d\n", 123)

	// This prints a binary representation.
	//这个输出二进制表示形式。
	fmt.Printf("%b\n", 14)

	// This prints the character corresponding to the
	// given integer.
	//这个输出给定整数的对应字符。
	fmt.Printf("%c\n", 33)

	// `%x` provides hex encoding.
	//%x 提供十六进制编码。
	fmt.Printf("%x\n", 456)

	// There are also several formatting options for
	// floats. For basic decimal formatting use `%f`.
	//对于浮点型同样有很多的格式化选项。使用 %f 进行最基本的十进制格式化。
	fmt.Printf("%f\n", 78.9)

	// `%e` and `%E` format the float in (slightly
	// different versions of) scientific notation.
	//%e 和 %E 将浮点型格式化为（稍微有一点不同的）科学技科学记数法表示形式。
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// For basic string printing use `%s`.
	//使用 %s 进行基本的字符串输出。
	fmt.Printf("%s\n", "\"string\"")

	// To double-quote strings as in Go source, use `%q`.
	//像 Go 源代码中那样带有双引号的输出，使用 %q。
	fmt.Printf("%q\n", "\"string\"")

	// As with integers seen earlier, `%x` renders
	// the string in base-16, with two output characters
	// per byte of input.
	//和上面的整形数一样，%x 输出使用 base-16 编码的字符串，每个字节使用 2 个字符表示。
	fmt.Printf("%x\n", "hex this")

	// To print a representation of a pointer, use `%p`.
	//要输出一个指针的值，使用 %p。
	fmt.Printf("%p\n", &p)

	// When formatting numbers you will often want to
	// control the width and precision of the resulting
	// figure. To specify the width of an integer, use a
	// number after the `%` in the verb. By default the
	// result will be right-justified and padded with
	// spaces.
	/*
	当输出数字的时候，你将经常想要控制输出结果的宽度和精度，
	可以使用在 % 后面使用数字来控制输出宽度。
	默认结果使用右对齐并且通过空格来填充空白部分。
	*/
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// You can also specify the width of printed floats,
	// though usually you'll also want to restrict the
	// decimal precision at the same time with the
	// width.precision syntax.
	/*
	你也可以指定浮点型的输出宽度，同时也可以通过 宽度.精度 的语法来指定输出的精度。
	*/
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// To left-justify, use the `-` flag.
	//要左对齐，使用 - 标志。
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// You may also want to control width when formatting
	// strings, especially to ensure that they align in
	// table-like output. For basic right-justified width.
	/*
	你也许也想控制字符串输出时的宽度，特别是要确保他们在类表格输出时的对齐。这是基本的右对齐宽度表示。
	*/
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// To left-justify use the `-` flag as with numbers.
	//要左对齐，和数字一样，使用 - 标志。
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// So far we've seen `Printf`, which prints the
	// formatted string to `os.Stdout`. `Sprintf` formats
	// and returns a string without printing it anywhere.
	/*
	到目前为止，我们已经看过 Printf了，它通过 os.Stdout输出格式化的字符串。
	Sprintf 则格式化并返回一个字符串而不带任何输出。
	*/
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// You can format+print to `io.Writers` other than
	// `os.Stdout` using `Fprintf`.
	//你可以使用 Fprintf 来格式化并输出到 io.Writers而不是 os.Stdout。
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
