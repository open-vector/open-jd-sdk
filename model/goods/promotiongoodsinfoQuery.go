package goods

import "github.com/open-vector/open-jd-sdk/model"

// 根据skuId查询商品信息接口
// 通过SKUID查询推广商品的名称、主图、类目、价格、物流、是否自营、30天引单数量等详细信息，支持批量获取。通常用于在媒体侧展示商品详情。

/**
request
*/

type PromotiongoodsinfoQueryRequest struct {
	// 多个sku逗号分隔
	SkuIds string `json:"skuIds,omitempty"`
}

func (PromotiongoodsinfoQueryRequest) GetMethod() string {
	return "jd.union.open.goods.promotiongoodsinfo.query"
}

/**
response
*/

type PromotiongoodsinfoQueryResult struct {
	model.BaseResult
	Data []PromotionGoodsResp `json:"data"`
}

type PromotionGoodsResp struct {
	SkuId             int    `json:"skuId"`
	UnitPrice         int    `json:"unitPrice"`
	MaterialUrl       string `json:"materialUrl"`
	EndDate           int    `json:"endDate"`
	IsFreeFreightRisk int    `json:"isFreeFreightRisk"`
	IsFreeShipping    int    `json:"isFreeShipping"`
	CommisionRatioWl  int    `json:"commisionRatioWl"`
	CommisionRatioPc  int    `json:"commisionRatioPc"`
	ImgUrl            string `json:"imgUrl"`
	Vid               int    `json:"vid"`
	CidName           string `json:"cidName"`
	WlUnitPrice       int    `json:"wlUnitPrice"`
	Cid2Name          string `json:"cid2Name"`
	IsSeckill         int    `json:"isSeckill"`
	Cid2              int    `json:"cid2"`
	Cid3Name          string `json:"cid3Name"`
	InOrderCount      int    `json:"inOrderCount"`
	Cid3              int    `json:"cid3"`
	ShopId            int    `json:"shopId"`
	IsJdSale          int    `json:"isJdSale"`
	GoodsName         string `json:"goodsName"`
	StartDate         int    `json:"startDate"`
	Cid               int    `json:"cid"`
}
