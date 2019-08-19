// Go offers built-in support for JSON encoding and
// decoding, including to and from built-in and custom
// data types.
//Go 提供内置的 JSON 编解码支持，包括内置或者自定义类型与 JSON 数据之间的转化。

package main

import "encoding/json"
import "fmt"
import "os"

// We'll use these two structs to demonstrate encoding and
// decoding of custom types below.
//下面我们将使用这两个结构体来演示自定义类型的编码和解码。
type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// First we'll look at encoding basic data types to
	// JSON strings. Here are some examples for atomic
	// values.
	//首先我们来看一下基本数据类型到 JSON 字符串的编码过程。这里是一些原子值的例子。
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// And here are some for slices and maps, which encode
	// to JSON arrays and objects as you'd expect.
	//这里是一些切片和 map 编码成 JSON 数组和对象的例子。
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// The JSON package can automatically encode your
	// custom data types. It will only include exported
	// fields in the encoded output and will by default
	// use those names as the JSON keys.
	//JSON 包可以自动的编码你的自定义类型。编码仅输出可导出的字段，并且默认使用他们的名字作为 JSON 数据的键。
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// You can use tags on struct field declarations
	// to customize the encoded JSON key names. Check the
	// definition of `response2` above to see an example
	// of such tags.
	//你可以给结构字段声明标签来自定义编码的 JSON 数据键名称。在上面 Response2 的定义可以作为这个标签这个的一个例子。
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Now let's look at decoding JSON data into Go
	// values. Here's an example for a generic data
	// structure.
	//现在来看看解码 JSON 数据为 Go 值的过程。这里是一个普通数据结构的解码例子。
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// We need to provide a variable where the JSON
	// package can put the decoded data. This
	// `map[string]interface{}` will hold a map of strings
	// to arbitrary data types.
	/*
	我们需要提供一个 JSON 包可以存放解码数据的变量。这里的 map[string]interface{} 将保存一个 string 为键，值为任意值的map。
	*/
	var dat map[string]interface{}

	// Here's the actual decoding, and a check for
	// associated errors.
	//这里就是实际的解码和相关的错误检查。
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// In order to use the values in the decoded map,
	// we'll need to convert them to their appropriate type.
	// For example here we convert the value in `num` to
	// the expected `float64` type.
	/*
	为了使用解码 map 中的值，我们需要将他们进行适当的类型转换。例如这里我们将 num 的值转换成 float64类型。
	*/
	num := dat["num"].(float64)
	fmt.Println(num)

	// Accessing nested data requires a series of
	// conversions.
	//访问嵌套的值需要一系列的转化。
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// We can also decode JSON into custom data types.
	// This has the advantages of adding additional
	// type-safety to our programs and eliminating the
	// need for type assertions when accessing the decoded
	// data.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// In the examples above we always used bytes and
	// strings as intermediates between the data and
	// JSON representation on standard out. We can also
	// stream JSON encodings directly to `os.Writer`s like
	// `os.Stdout` or even HTTP response bodies.
	/*
	我们也可以解码 JSON 值到自定义类型。这个功能的好处就是可以为我们的程序带来额外的类型安全加强，
	并且消除在访问数据时的类型断言。
	*/
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
/*
在上面的例子中，我们经常使用 byte 和 string 作为使用标准输出时数据和 JSON 表示之间的中间值。
我们也可以和os.Stdout 一样，直接将 JSON 编码直接输出至 os.Writer流中，或者作为 HTTP 响应体。
*/