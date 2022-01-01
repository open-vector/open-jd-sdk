package main

import (
	"fmt"
	"github.com/open-vector/open-jd-sdk/client"
	"github.com/open-vector/open-jd-sdk/model/goods"
	"github.com/open-vector/open-jd-sdk/util/jsonUtil"
)

func materialQuery() {
	// 定义请求体
	goodReq := goods.MaterialGoodsReq{EliteId: 2}
	request := goods.MaterialQueryRequest{GoodsReq: goodReq}
	fmt.Println(jsonUtil.ObjToString(request))

	// 实例化一个client
	// todo 让使用sdk的用户指定ak sk等等 目前写死
	jdClient := client.New("2ac1c5db7e31d4a85d145ac19fa3d4e8", "629a0b8f8a41479f82d1521d12854f50")

	// 定义响应并发送请求
	var response goods.MaterialQueryResult
	jdClient.Execute(request, &response)
	fmt.Println(response)
	fmt.Println(jsonUtil.ObjToString(response))
}

func jingfenQuery() {
	// 定义请求体
	goodReq := goods.JFGoodsReq{EliteId: 2}
	request := goods.JingfenQueryRequest{GoodsReq: goodReq}
	fmt.Println(jsonUtil.ObjToString(request))

	// 实例化一个client
	jdClient := client.New("2ac1c5db7e31d4a85d145ac19fa3d4e8", "629a0b8f8a41479f82d1521d12854f50")

	// 定义响应并发送请求
	var response goods.JingfenQueryResult
	jdClient.Execute(request, &response)
	fmt.Println(response)
	fmt.Println(jsonUtil.ObjToString(response))
}

func main() {
	materialQuery()
}
