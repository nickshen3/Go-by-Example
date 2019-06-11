// _Slices_ are a key data type in Go, giving a more
// powerful interface to sequences than arrays.
//Slice 是 Go 中一个关键的数据类型，是一个比数组更 加强大的序列接口
package main

import "fmt"

func main() {

	// Unlike arrays, slices are typed only by the
	// elements they contain (not the number of elements).
	// To create an empty slice with non-zero length, use
	// the builtin `make`. Here we make a slice of
	// `string`s of length `3` (initially zero-valued).
	/*
	与数组不同，slice 的类型仅由它所包含的元素决定（不需要 元素的个数）。
	要创建一个长度非零的空 slice，需要使用内建的方法 make。
	这里我们创建了一 个长度为3的 string 类型 slice（初始化为零值）。
	*/
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// We can set and get just like with arrays.
	//我们可以和数组一样设置和得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` returns the length of the slice as expected.
	//len 返回 slice 的长度
	fmt.Println("len:", len(s))

	// In addition to these basic operations, slices
	// support several more that make them richer than
	// arrays. One is the builtin `append`, which
	// returns a slice containing one or more new values.
	// Note that we need to accept a return value from
	// `append` as we may get a new slice value.
	/*
	除了基本操作外，slice 支持比数组更丰富的操作。 其中一个是内建的 append，
	它返回一个包含了一个 或者多个新值的 slice。注意由于 append 可能返回 新的 slice，
	我们需要接受其返回值。
	*/
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be `copy`'d. Here we create an
	// empty slice `c` of the same length as `s` and copy
	// into `c` from `s`.
	//Slice 也可以被 copy。这里我们创建一个空的和 s 有 相同长度的 slice c，并且将 s 复制给 c。
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a "slice" operator with the syntax
	// `slice[low:high]`. For example, this gets a slice
	// of the elements `s[2]`, `s[3]`, and `s[4]`.
	/*
	Slice 支持通过 slice[low:high] 语法进行“切片”操 作。
	例如，这里得到一个包含元素 s[2], s[3], s[4] 的 slice。
	*/
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) `s[5]`.
	//这个 slice 从 s[0] 切片到 s[5]（不包含）。
	l = s[:5]
	fmt.Println("sl2:", l)

	// And this slices up from (and including) `s[2]`.
	//这个 slice 从 s[2] （包含）开始切片。
	l = s[2:]
	fmt.Println("sl3:", l)

	// We can declare and initialize a variable for slice
	// in a single line as well.
	//我们可以在一行代码中声明并初始化一个 slice 变量。
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slices can be composed into multi-dimensional data
	// structures. The length of the inner slices can
	// vary, unlike with multi-dimensional arrays.
	//Slice 可以组成多维数据结构。内部的 slice 长度可以不 一致，这和多维数组不同。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
