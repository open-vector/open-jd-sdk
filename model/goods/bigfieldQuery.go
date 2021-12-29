package goods

import "github.com/open-vector/open-jd-sdk/model"

// 商品详情查询接口

/**
request
*/

type BigFieldQueryRequest struct {
	GoodsReq BigFieldGoodsReq `json:"goodsReq"`
}

func (BigFieldQueryRequest) GetMethod() string {
	return "jd.union.open.goods.bigfield.query"
}

type BigFieldGoodsReq struct {
	SkuIds []int    `json:"skuIds"`
	Fields []string `json:"fields"`
}

/**
response
*/

type BigfieldQueryResult struct {
	model.BaseResult
	Data []BigFieldGoodsResp `json:"data"`
}

type BigFieldGoodsResp struct {
	SkuId             int         `json:"skuId"`
	SkuName           string      `json:"skuName"`
	CategoryInfo      interface{} `json:"categoryInfo"`
	ImageInfo         interface{} `json:"imageInfo"`
	BaseBigFieldInfo  interface{} `json:"baseBigFieldInfo"`
	BookBigFieldInfo  interface{} `json:"bookBigFieldInfo"`
	VideoBigFieldInfo interface{} `json:"videoBigFieldInfo"`
	MainSkuId         int         `json:"mainSkuId"`
	ProductId         int         `json:"productId"`
	SkuStatus         int         `json:"skuStatus"`
	Owner             string      `json:"owner"`
	DetailImages      string      `json:"detailImages"`
}
