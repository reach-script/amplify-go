// ---------------------------------------------------------------
//
//  decode_jwt.go
//
//                Feb/07/2021
// ---------------------------------------------------------------
package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ---------------------------------------------------------------
func decode_jwt_proc(str_token string) string {
	aaa := strings.Split(str_token, ".")
	str_bbb := strings.Replace(aaa[1], "-", "+", -1)
	str_ccc := strings.Replace(str_bbb, "_", "/", -1)

	llx := len(str_ccc)
	nnx := ((4 - llx%4) % 4)
	ssx := strings.Repeat("=", nnx)
	str_ddd := strings.Join([]string{str_ccc, ssx}, "")
	ppp, err := b64.StdEncoding.DecodeString(str_ddd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "*** error *** StdEncoding.DecodeString ***\n")
		fmt.Println("error:", err)
		return "error"
	}

	uEnc := b64.URLEncoding.EncodeToString([]byte(ppp))
	decode, _ := b64.URLEncoding.DecodeString(uEnc)

	return string(decode)
}

// ---------------------------------------------------------------
func main() {

	fmt.Fprintf(os.Stderr, "*** 開始 ***\n")
	file_token := os.Args[1]
	fmt.Fprintf(os.Stderr, "file_token = "+file_token+"\n")

	buff, _ := ioutil.ReadFile(file_token)
	fmt.Fprintf(os.Stderr, "len(buff) = %d\n", len(buff))

	json_str := decode_jwt_proc(string(buff))
	fmt.Println(json_str)

	fmt.Fprintf(os.Stderr, "*** 終了 ***\n")
}

// ---------------------------------------------------------------
