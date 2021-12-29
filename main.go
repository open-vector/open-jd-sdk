package main

import (
	"fmt"
	"github.com/open-vector/open-jd-sdk/client"
	"github.com/open-vector/open-jd-sdk/model/goods"
	"github.com/open-vector/open-jd-sdk/util/jsonUtil"
)

func main() {
	// 定义请求体
	goodReq := goods.JFGoodsReq{EliteId: 2}
	jfQueryRequest := goods.JingfenQueryRequest{goodReq}

	// 实例化一个client
	// todo 让使用sdk的用户指定ak sk等等 目前写死
	jdClient := client.New()

	// 定义响应并发送请求
	var jingfenQueryResult goods.JingfenQueryResult
	jdClient.Execute(jfQueryRequest, &jingfenQueryResult)
	fmt.Println(jingfenQueryResult)
	fmt.Println(jsonUtil.ObjToString(jingfenQueryResult))
}
