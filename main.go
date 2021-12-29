package main

import (
	"fmt"
	"github.com/open-vector/open-jd-sdk/client"
	"github.com/open-vector/open-jd-sdk/model"
	"github.com/open-vector/open-jd-sdk/util/jsonUtil"
)

func main() {
	// 定义请求体
	goodReq := model.JFGoodsReq{EliteId: 2}
	jfQueryRequest := model.JingfenQueryRequest{goodReq}

	// 实例化一个client
	// todo 让使用sdk的用户指定ak sk等等 目前写死
	jdClient := client.New()

	// 定义响应并发送请求
	var jingfenQueryResult model.JingfenQueryResult
	jdClient.Execute(jfQueryRequest, &jingfenQueryResult)
	fmt.Println(jingfenQueryResult)
	fmt.Println(jsonUtil.ObjToString(jingfenQueryResult))
}
