// The `filepath` package provides functions to parse
// and construct *file paths* in a way that is portable
// between operating systems; `dir/file` on Linux vs.
// `dir\file` on Windows, for example.
//filepath 包提供了以可在操作系统之间移植的方式解析和构造文件路径的函数; 
//例如，Windows 上 的 dir / file 和 Windows 上的.dirfile‘。
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// `Join` should be used to construct paths in a
	// portable way. It takes any number of arguments
	// and constructs a hierarchical path from them.
	//Join 应该用于以便携方式构造路径。它接受任意数量的参数并从中构造分层路径。
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// You should always use `Join` instead of
	// concatenating `/`s or `\`s manually. In addition
	// to providing portability, `Join` will also
	// normalize paths by removing superfluous separators
	// and directory changes.
	//你应该总是使用 Join 而不是手动连接/s 或 s。除了提供可移植性之外，
	//Join‘还会通过删除多余的 分隔符和目录更改来对路径进行 alsonormal 化。
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` and `Base` can be used to split a path to the
	// directory and the file. Alternatively, `Split` will
	// return both in the same call.
	//Dir 和 Base 可用于分割目录和文件的路径。或者，Split 将在同一个调用中返回。
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// We can check whether a path is absolute.
	//我们可以检查路径是否是绝对的。
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Some file names have extensions following a dot. We
	// can split the extension out of such names with `Ext`.
	//某些文件名在点后面有扩展名。我们可以使用 Ext 将扩展名从这些名称中分离出来。
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// To find the file's name with the extension removed,
	// use `strings.TrimSuffix`.
	//要找到删除了扩展名的文件名，请使用 strings.TrimSuffix。
	fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` finds a relative path between a *base* and a
	// *target*. It returns an error if the target cannot
	// be made relative to base.
	//Rel 找到 * base 和 * * 之间的相对路径。如果目标不能相对于 base，则返回错误。
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
