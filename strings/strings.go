package strings

//StrCut
func StrCut(str string, strLen int) string {
	rs := []rune(str)
	length := len(rs)

	if length <= strLen {
		return str
	}
	return string(rs[0:strLen])
}
