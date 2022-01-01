package goods

import "github.com/open-vector/open-jd-sdk/model"

// 猜你喜欢商品推荐

/**
request
*/
type MaterialQueryRequest struct {
	GoodsReq MaterialGoodsReq `json:"goodsReq"`
}

func (MaterialQueryRequest) GetMethod() string {
	return "jd.union.open.goods.material.query"
}

type MaterialGoodsReq struct {
	EliteId      int    `json:"eliteId"`
	PageIndex    int    `json:"pageIndex"`
	PageSize     int    `json:"pageSize"`
	SortName     string `json:"sortName"`
	Sort         string `json:"sort"`
	Pid          string `json:"pid"`
	SubUnionId   string `json:"subUnionId"`
	SiteId       string `json:"siteId"`
	PositionId   string `json:"positionId"`
	Ext1         string `json:"ext1"`
	SkuId        int    `json:"skuId"`
	HasCoupon    int    `json:"hasCoupon"`
	UserIdType   int    `json:"userIdType"`
	UserId       string `json:"userId"`
	Fields       string `json:"fields"`
	ForbidTypes  string `json:"forbidTypes"`
	OrderId      string `json:"orderId"`
	GroupId      int    `json:"groupId"`
	OwnerUnionId int    `json:"ownerUnionId"`
	BenefitType  int    `json:"benefitType"`
}

/**
response
*/

type MaterialQueryResult struct {
	model.BaseResult
	TotalCount int                 `json:"totalCount"`
	Data       []MaterialGoodsResp `json:"data"`
}

type MaterialGoodsResp struct {
	CategoryInfo           interface{} `json:"categoryInfo"`
	Comments               int         `json:"comments"`
	CommissionInfo         interface{} `json:"commissionInfo"`
	CouponInfo             interface{} `json:"couponInfo"`
	GoodCommentsShare      int         `json:"goodCommentsShare"`
	ImageInfo              interface{} `json:"imageInfo"`
	InOrderCount30Days     int         `json:"inOrderCount30Days"`
	PriceInfo              interface{} `json:"priceInfo"`
	ShopInfo               interface{} `json:"shopInfo"`
	SkuId                  int         `json:"skuId"`
	SkuName                string      `json:"skuName"`
	IsHot                  int         `json:"isHot"`
	Spuid                  int         `json:"spuid"`
	BrandCode              string      `json:"brandCode"`
	Owner                  string      `json:"owner"`
	PinGouInfo             interface{} `json:"pinGouInfo"`
	ResourceInfo           interface{} `json:"resourceInfo"`
	InOrderCount30DaysSku  int         `json:"inOrderCount30DaysSku"`
	SeckillInfo            interface{} `json:"seckillInfo"`
	JxFlags                []int       `json:"jxFlags"`
	VideoInfo              interface{} `json:"videoInfo"`
	PromotionInfo          interface{} `json:"promotionInfo"`
	BookInfo               interface{} `json:"bookInfo"`
	ForbidTypes            []int       `json:"forbidTypes"`
	DeliveryType           interface{} `json:"deliveryType"`
	SkuLabelInfo           interface{} `json:"skuLabelInfo"`
	PromotionLabelInfoList string      `json:"promotionLabelInfoList"`
	MaterialUrl            interface{} `json:"materialUrl"`
	PreSaleInfo            interface{} `json:"preSaleInfo"`
	ReserveInfo            int         `json:"reserveInfo"`
}
