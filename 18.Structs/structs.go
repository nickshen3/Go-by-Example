// Go's _structs_ are typed collections of fields.
// They're useful for grouping data together to form
// records.
//Go 的结构体(struct) 是带类型的字段(fields)集合。 这在组织数据时非常有用。
package main

import "fmt"

// This `person` struct type has `name` and `age` fields.
//这里的 person 结构体包含了 name 和 age 两个字段。
type person struct {
	name string
	age  int
}

func main() {

	// This syntax creates a new struct.
	//使用这个语法创建新的结构体元素。
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initializing a struct.
	//你可以在初始化一个结构体元素时指定字段名字。
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	//省略的字段将被初始化为零值。
	fmt.Println(person{name: "Fred"})

	// An `&` prefix yields a pointer to the struct.
	//& 前缀生成一个结构体指针。
	fmt.Println(&person{name: "Ann", age: 40})

	// Access struct fields with a dot.
	//使用.来访问结构体字段。
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the
	// pointers are automatically dereferenced.
	//也可以对结构体指针使用. - 指针会被自动解引用。
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable.
	//结构体是可变(mutable)的。
	sp.age = 51
	fmt.Println(sp.age)
}
