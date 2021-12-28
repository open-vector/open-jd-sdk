package main

import (
	"encoding/json"
	"fmt"
	"github.com/open-vector/mallPromotie/backend/client"
	"github.com/open-vector/mallPromotie/backend/model"
)

func main() {
	// 定义请求体
	var goodReq model.JFGoodsReq
	goodReq.EliteId = 2
	jfQueryRequest := model.JFQueryRequest{GoodReq: goodReq}

	// 实例化一个client
	jdClient := client.New()

	// 定义响应体
	var jingfenQueryResult model.JingfenQueryResult
	// 调用请求，并将结果反序列化到响应体里
	// 这里发送请求已经抽象好，但是不知道是不是最好的方法，每个requet都要实现接口RequestInterface的GetByte方法，响应也没抽象
	// todo 待优化
	json.Unmarshal(jdClient.Execute("jd.union.open.goods.jingfen.query", jfQueryRequest), &jingfenQueryResult)
	fmt.Println(jfQueryRequest)
}
