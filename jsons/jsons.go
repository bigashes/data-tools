package jsons

import "encoding/json"

// JsonMarshalString json转为string
func JsonMarshalString(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	if b == nil {
		return "[]", nil
	}
	return string(b), nil
}

// JsonUnmarshalString string转为json
func JsonUnmarshalString(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}

// JsonMarshalMustString json转为string
func JsonMarshalMustString(v interface{}) string {
	b, _ := JsonMarshalString(v)
	return b
}
