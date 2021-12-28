package main

import (
	"encoding/json"
	"fmt"
	"github.com/open-vector/mallPromotie/backend/client"
	"github.com/open-vector/mallPromotie/backend/model"
)

func main() {
	fmt.Println("open-jd-sdk")
	var goodReq model.JFGoodsReq
	goodReq.EliteId = 2
	jfQueryRequest := model.JFQueryRequest{GoodReq: goodReq}

	jdClient := client.New()
	var jingfenQueryResult model.JingfenQueryResult
	json.Unmarshal(jdClient.Execute("jd.union.open.goods.jingfen.query", jfQueryRequest), &jingfenQueryResult)
	fmt.Println(jfQueryRequest)
}
