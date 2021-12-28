package jsonUtil

import "encoding/json"

func ObjToString(obj interface{}) string {
	bytes, _ := json.Marshal(obj)
	return string(bytes)
}

func GetByte(obj interface{}) []byte {
	bytes, _ := json.Marshal(obj)
	return bytes
}

func StringToObj(str string, t interface{}) error {
	data := []byte(str)
	return json.Unmarshal(data, t)
}
