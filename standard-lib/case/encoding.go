package _case

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func Encoding() {
	type user struct {
		ID   int64
		Name string
		Age  uint8
	}
	u := user{ID: 1, Name: "wang", Age: 27}
	//json序列化
	bytes, err := json.Marshal(u)
	fmt.Println(bytes, err)
	//json反序列化
	u1 := user{}
	err = json.Unmarshal(bytes, &u1)
	fmt.Println(u1, err)
	//base64编解码
	str := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println(str)
	//base64反解码
	bytes1, err := base64.StdEncoding.DecodeString(str)
	fmt.Println(bytes1, err)
	//16进制编解码
	str1 := hex.EncodeToString(bytes1)
	fmt.Println(str1)
	//反解16进制解码
	bytes2, err := hex.DecodeString(str1)
	fmt.Println(bytes2, err)
}
