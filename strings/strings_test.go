package strings

import "testing"

func TestStrCut(t *testing.T) {
	s := "123456789"

	cut := StrCut(s, 3)
	if cut != "123" {
		t.Error("err")
	}
}
