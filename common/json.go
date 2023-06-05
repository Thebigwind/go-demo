package common

import "encoding/json"

// ToJSONStringPretty encodingJson MarshalIndent Pretty
func ToJSONStringPretty(v interface{}) string {
	jsonStr, _ := json.MarshalIndent(v, "", "    ")
	return string(jsonStr)
}

// FromJSONString encodingJson Unmarshal string
func FromJSONString(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}
