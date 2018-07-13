package encryptions

import (
	"crypto/md5"
	"io"
	"fmt"
	"bytes"
	"encoding/hex"
)

func Salt(password string) (string){
	// 计算密码MD5
	c := md5.New()
	io.WriteString(c, password)
	spw := fmt.Sprintf("%x\n", c.Sum(nil))

	// 指定两个(salt)
	salt1 := "@#$%"
	salt2 := "^&*()"

	// 拼接密码MD5
	buf := bytes.NewBufferString("")

	// 拼接密码
	io.WriteString(buf, salt1)
	io.WriteString(buf, spw)
	io.WriteString(buf, salt2)

	// 拼接密码计算MD5
	t := md5.New()
	io.WriteString(t, buf.String())

	// 输出
	fmt.Printf("%x\n", t.Sum(nil))
	return  hex.EncodeToString(t.Sum(nil))
}