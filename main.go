package main

import (
	"fmt"
	"github.com/open-vector/mallPromotie/backend/client"
	"github.com/open-vector/mallPromotie/backend/model"
	"github.com/open-vector/mallPromotie/backend/util/jsonUtil"
)

func main() {
	// 定义请求体
	goodReq := model.JFGoodsReq{EliteId: 2}
	jfQueryRequest := model.JFQueryRequest{GoodReq: goodReq}

	// 实例化一个client
	// todo 让使用sdk的用户指定ak sk等等 目前写死
	jdClient := client.New()

	// 定义响应体
	//var jingfenQueryResult model.JingfenQueryResult
	// 调用请求，并将结果反序列化到响应体里
	// 这里发送请求已经抽象好，但是不知道是不是最好的方法，每个requet都要实现接口RequestInterface的GetByte方法，响应也没抽象
	// todo 待优化 这个method应该也是可以抽一下
	var jingfenQueryResult model.JingfenQueryResult
	jdClient.Execute(jfQueryRequest, &jingfenQueryResult)
	fmt.Println(jingfenQueryResult)
	fmt.Println(jsonUtil.ObjToString(jingfenQueryResult))
}
