// Go's `math/rand` package provides
// [pseudorandom number](http://en.wikipedia.org/wiki/Pseudorandom_number_generator)
// generation.
// Go 的 math/rand 包提供了伪随机数生成器（英）。
package main

import "time"
import "fmt"
import "math/rand"

func main() {

	// For example, `rand.Intn` returns a random `int` n,
	// `0 <= n < 100`.
	// 例如，rand.Intn 返回一个随机的整数 n，0 <= n <= 100。
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// `rand.Float64` returns a `float64` `f`,
	// `0.0 <= f < 1.0`.
	//rand.Float64 返回一个64位浮点数 f，0.0 <= f <= 1.0。
	fmt.Println(rand.Float64())

	// This can be used to generate random floats in
	// other ranges, for example `5.0 <= f' < 10.0`.
	//这个技巧可以用来生成其他范围的随机浮点数，例如5.0 <= f <= 10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// The default number generator is deterministic, so it'll
	// produce the same sequence of numbers each time by default.
	// To produce varying sequences, give it a seed that changes.
	// Note that this is not safe to use for random numbers you
	// intend to be secret, use `crypto/rand` for those.
	/*
	默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。要产生变化的序列，
	需要给定一个变化的种子。需要注意的是，如果你出于加密目的，需要使用随机数的话，
	请使用 crypto/rand 包，此方法不够安全
	*/
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Call the resulting `rand.Rand` just like the
	// functions on the `rand` package.
	//调用上面返回的 rand.Source 的函数和调用 rand 包中函数是相同的。
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// If you seed a source with the same number, it
	// produces the same sequence of random numbers.
	//如果使用相同的种子生成的随机数生成器，将会产生相同的随机数序列。
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
/*
参阅 math/rand 包文档，提供了 Go 可以提供的其他随量的参考信息。
*/