package model

import "encoding/json"

// RequestInterface 定义一个通用请求interface 所有的request都必须实现这两个方法
type RequestInterface interface {
	GetMethod() string
}

// 返回结果统一处理
// 这里只考虑了不报错的情况，报错情况出了问题再说
func ResponseHandle(body []byte, response interface{}) {
	var resultMap map[string]Response
	json.Unmarshal(body, &resultMap)
	for _, v := range resultMap {
		json.Unmarshal([]byte(v.QueryResult), response)
	}
}

type Response struct {
	Code string `json:"code"`
	// 这里想直接拿JingfenQueryResult接，接不到，必须拿string先接然后再转
	QueryResult string `json:"queryResult"`
}
