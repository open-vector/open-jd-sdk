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
	request, _ := json.Marshal(model.JFQueryRequest{GoodReq: goodReq})

	jdClient := client.New()
	res := jdClient.Execute("jd.union.open.goods.jingfen.query", request)
	fmt.Println(res)
}
