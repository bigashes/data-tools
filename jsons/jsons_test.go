package jsons

import (
	"testing"
)

func TestJson(t *testing.T) {
	var a []int64
	marshal, err := JsonMarshalString(a)
	if err != nil {
		return
	}

	t.Log(marshal)
}
