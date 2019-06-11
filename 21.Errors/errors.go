// In Go it's idiomatic to communicate errors via an
// explicit, separate return value. This contrasts with
// the exceptions used in languages like Java and Ruby and
// the overloaded single result / error value sometimes
// used in C. Go's approach makes it easy to see which
// functions return errors and to handle them using the
// same language constructs employed for any other,
// non-error tasks.
/*
符合 Go 语言习惯的做法是使用一个独立、明确的返回值来传递错误信息。 
这与使用异常(exception)的 Java 和 Ruby 以及在 C 语言中
有时用到的重载(overloaded)的单返回/错误值有着明显的不同。 
Go 语言的处理方式能清楚的知道哪个函数 返回了错误，并能像调用那些没有出错的函数一样调用。
*/

package main

import "errors"
import "fmt"

// By convention, errors are the last return value and
// have type `error`, a built-in interface.
//按照惯例，错误通常是最后一个返回值并且是 error 类型，一个内建的接口。
func f1(arg int) (int, error) {
	if arg == 42 {

		// `errors.New` constructs a basic `error` value
		// with the given error message.
		//errors.New 构造一个使用给定的错误信息的基本 error 值。
		return -1, errors.New("can't work with 42")

	}

	// A `nil` value in the error position indicates that
	// there was no error.
	//返回错误值为 nil 代表没有错误。
	return arg + 3, nil
}

// It's possible to use custom types as `error`s by
// implementing the `Error()` method on them. Here's a
// variant on the example above that uses a custom type
// to explicitly represent an argument error.
/*
通过实现 Error 方法来自定义 error 类型是可以的。 
这里使用自定义错误类型来表示上面例子中的参数错误。
*/
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// In this case we use `&argError` syntax to build
		// a new struct, supplying values for the two
		// fields `arg` and `prob`.
		/*
		在这个例子中，我们使用 &argError 语法来建立一个 新的结构体，
		并提供了 arg 和 prob 这两个字段 的值。
		*/
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// The two loops below test out each of our
	// error-returning functions. Note that the use of an
	// inline error check on the `if` line is a common
	// idiom in Go code.
	/*
	下面的两个循环测试了各个返回错误的函数。注意在 if 行内的错误检查代码，在 Go 中是一个普遍的用法。
	*/
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// If you want to programmatically use the data in
	// a custom error, you'll need to get the error as an
	// instance of the custom error type via type
	// assertion.
	//你如果想在程序中使用一个自定义错误类型中的数据，你需要通过类型断言来得到这个错误类型的实例。
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
