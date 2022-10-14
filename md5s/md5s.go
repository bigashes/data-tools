package md5s

import (
	"crypto/md5"
	"fmt"
)

func Md5(text string) string {
	m := md5.Sum([]byte(text))
	return fmt.Sprintf("%x", m)
}
