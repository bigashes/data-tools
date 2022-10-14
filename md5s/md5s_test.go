package md5s

import (
	"log"
	"testing"
)

func TestMd5s(t *testing.T) {
	a := "123456"
	b := Md5(a)

	log.Printf("%s", b)
}
