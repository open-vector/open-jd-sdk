package model

import (
	"encoding/json"
	"strings"
)

// RequestInterface 定义一个通用请求interface 所有的request都必须实现这两个方法
type RequestInterface interface {
	GetMethod() string
}

// 返回结果统一处理
// 这里只考虑了不报错的情况，报错情况出了问题再说
func ResponseHandle(body []byte, response interface{}, method string) {
	if strings.HasSuffix(method, "query") {
		var resultMap map[string]ResponseQuery
		json.Unmarshal(body, &resultMap)
		for _, v := range resultMap {
			json.Unmarshal([]byte(v.QueryResult), response)
		}
	} else {
		var resultMap map[string]ResponseGet
		json.Unmarshal(body, &resultMap)
		for _, v := range resultMap {
			json.Unmarshal([]byte(v.GetResult), response)
		}
	}

}

type ResponseQuery struct {
	Code string `json:"code"`
	// 这里想直接拿JingfenQueryResult接，接不到，必须拿string先接然后再转
	QueryResult string `json:"queryResult"`
}

type ResponseGet struct {
	Code string `json:"code"`
	// 这里想直接拿JingfenQueryResult接，接不到，必须拿string先接然后再转
	GetResult string `json:"getResult"`
}

type BaseResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
