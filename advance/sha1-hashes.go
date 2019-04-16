/*
	SHA1 hash加密算法
*/
package main

//Go implements several hash functions in various crypto/* packages.
import "crypto/sha1"
import "fmt"

func main() {
	s := "sha1 this string"

	//The pattern for generating a hash is sha1.New(),
	//sha1.Write(bytes), then sha1.Sum([]byte{}). Here we start with a new hash.
	h := sha1.New()

	//Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes.
	h.Write([]byte(s))

	//This gets the finalized hash result as a byte slice.
	//The argument to Sum can be used to append to an
	//existing byte slice: it usually isn’t needed.
	bs := h.Sum(nil)

	//SHA1 values are often printed in hex, for example in git commits. Use the %x format verb to convert a hash results to a hex string.
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}

/*

运行： go run sha1-hashes.go
结果：
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94
*/
