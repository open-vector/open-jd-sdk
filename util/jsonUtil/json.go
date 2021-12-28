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

func StringToObj(str interface{}) string {
	return ""
}
