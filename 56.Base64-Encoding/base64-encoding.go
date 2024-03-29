// Go provides built-in support for [base64
// encoding/decoding](http://en.wikipedia.org/wiki/Base64).
//Go 提供内建的 base64 编解码支持。
package main

// This syntax imports the `encoding/base64` package with
// the `b64` name instead of the default `base64`. It'll
// save us some space below.
//这个语法引入了 encoding/base64 包并使用名称 b64代替默认的 base64。这样可以节省点空间。
import b64 "encoding/base64"
import "fmt"

func main() {

	// Here's the `string` we'll encode/decode.
	//这是将要编解码的字符串。
	data := "abc123!?$*&()'-=@~"

	// Go supports both standard and URL-compatible
	// base64. Here's how to encode using the standard
	// encoder. The encoder requires a `[]byte` so we
	// convert our `string` to that type.
	//Go 同时支持标准的和 URL 兼容的 base64 格式。编码需要使用 []byte 类型的参数，所以要将字符串转成此类型。
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Decoding may return an error, which you can check
	// if you don't already know the input to be
	// well-formed.
	//解码可能会返回错误，如果不确定输入信息格式是否正确，那么，你就需要进行错误检查了。
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// This encodes/decodes using a URL-compatible base64
	// format.
	//使用 URL 兼容的 base64 格式进行编解码。
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
/*
标准 base64 编码和 URL 兼容 base64 编码的编码字符串存在稍许不同（后缀为 + 和 -），
但是两者都可以正确解码为原始字符串。
*/