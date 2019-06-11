// _Maps_ are Go's built-in [associative data type](http://en.wikipedia.org/wiki/Associative_array)
// (sometimes called _hashes_ or _dicts_ in other languages).
//map 是 Go 内置关联数据类型（ 在一些其他的语言中称为哈希(hash) 或者字典(dict) ）。

package main

import "fmt"

func main() {

	// To create an empty map, use the builtin `make`:
	// `make(map[key-type]val-type)`.
	//要创建一个空 map，需要使用内建的 make: make(map[key-type]val-type).
	m := make(map[string]int)

	// Set key/value pairs using typical `name[key] = val`
	// syntax.
	//使用典型的 make[key] = val 语法来设置键值对。
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map with e.g. `fmt.Println` will show all of
	// its key/value pairs.
	//使用例如 fmt.Println 来打印一个 map 将会输出所有的 键值对。
	fmt.Println("map:", m)

	// Get a value for a key with `name[key]`.
	//使用 name[key] 来获取一个键的值
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// The builtin `len` returns the number of key/value
	// pairs when called on a map.
	//当对一个 map 调用内建的 len 时，返回的是键值对 数目
	fmt.Println("len:", len(m))

	// The builtin `delete` removes key/value pairs from
	// a map.
	//内建的 delete 可以从一个 map 中移除键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	// The optional second return value when getting a
	// value from a map indicates if the key was present
	// in the map. This can be used to disambiguate
	// between missing keys and keys with zero values
	// like `0` or `""`. Here we didn't need the value
	// itself, so we ignored it with the _blank identifier_
	// `_`.
	/*
	当从一个 map 中取值时，可选的第二返回值指示这个键 是否在这个 map 中。
	这可以用来消除键不存在和键有零值， 像 0 或者 "" 而产生的歧义。
	这里我们不需要值，所以 用_空白标识符(blank identifier)_忽略。
	*/
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// You can also declare and initialize a new map in
	// the same line with this syntax.
	//你也可以通过这个语法在同一行声明和初始化一个新的 map。
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
