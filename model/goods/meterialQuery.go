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
	EliteId    int    `json:"eliteId,omitempty"`
	PageIndex  int    `json:"pageIndex,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
	SortName   string `json:"sortName,omitempty"`
	Sort       string `json:"sort,omitempty"`
	Pid        string `json:"pid,omitempty"`
	SubUnionId string `json:"subUnionId,omitempty"`
	SiteId     string `json:"siteId,omitempty"`
	// 数字string
	PositionId  string `json:"positionId,omitempty"`
	Ext1        string `json:"ext1,omitempty"`
	SkuId       int    `json:"skuId,omitempty"`
	HasCoupon   int    `json:"hasCoupon,omitempty"`
	UserIdType  int    `json:"userIdType,omitempty"`
	UserId      string `json:"userId,omitempty"`
	Fields      string `json:"fields,omitempty"`
	ForbidTypes string `json:"forbidTypes,omitempty"`
	// 数字string
	OrderId      string `json:"orderId,omitempty"`
	GroupId      int    `json:"groupId,omitempty"`
	OwnerUnionId int    `json:"ownerUnionId,omitempty"`
	BenefitType  int    `json:"benefitType,omitempty"`
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
