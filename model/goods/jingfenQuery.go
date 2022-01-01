package goods

import "github.com/open-vector/open-jd-sdk/model"

// 精粉精选商品查询接口
// 京东联盟精选优质商品，每日更新，可通过频道ID查询各个频道下的精选商品。用获取的优惠券链接调用转链接口时，需传入搜索接口link字段返回的原始优惠券链接，
// 切勿对链接进行任何encode、decode操作，否则将导致转链二合一推广链接时校验失败。

/**
request
*/

type JingfenQueryRequest struct {
	GoodsReq JFGoodsReq `json:"goodsReq"`
}

func (JingfenQueryRequest) GetMethod() string {
	return "jd.union.open.goods.jingfen.query"
}

type JFGoodsReq struct {
	/**
	频道ID:1-好券商品,2-精选卖场,10-9.9包邮,15-京东配送,22-实时热销榜,23-为你推荐,24-数码家电,25-超市,26-母婴玩具,27-家具日用,
	28-美妆穿搭,30-图书文具,31-今日必推,32-京东好物,33-京东秒杀,34-拼购商品,40-高收益榜,41-自营热卖榜,108-秒杀进行中,
	109-新品首发,110-自营,112-京东爆品,125-首购商品,129-高佣榜单,130-视频商品,153-历史最低价商品榜,210-极速版商品,238-新人价商品,
	247-京喜9.9,249-京喜秒杀,315-秒杀未开始,340-时尚趋势品,341-3C新品,342-智能新品,343-3C长尾商品,345-时尚新品,346-时尚爆品,
	426-京喜自营,1001-选品库
	*/
	EliteId   int `json:"eliteId,omitempty"`
	PageIndex int `json:"pageIndex,omitempty"`
	PageSize  int `json:"pageSize,omitempty"`
	/**
	排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30DaysSku：sku维度30天引单量，
	comments：评论数，goodComments：好评数)
	*/
	SortName     string `json:"sortName,omitempty"`
	Sort         string `json:"sort,omitempty"`
	Pid          string `json:"pid,omitempty"`
	Fields       string `json:"fields,omitempty"`
	ForbidTypes  string `json:"forbidTypes,omitempty"`
	GroupId      int    `json:"groupId,omitempty"`
	OwnerUnionId int    `json:"ownerUnionId,omitempty"`
}

/**
response
*/

type JingfenQueryResult struct {
	model.BaseResult
	Data       []JFGoodsResp `json:"data"`
	RequestId  string        `json:"requestId"`
	TotalCount int           `json:"totalCount"`
}

type JFGoodsResp struct {
	CategoryInfo           interface{}   `json:"categoryInfo"`
	Comments               int           `json:"comments"`
	CommissionInfo         interface{}   `json:"commissionInfo"`
	CouponInfo             interface{}   `json:"couponInfo"`
	GoodCommentsShare      float32       `json:"goodCommentsShare"`
	ImageInfo              interface{}   `json:"imageInfo"`
	InOrderCount30Days     int           `json:"inOrderCount30Days"`
	MaterialUrl            string        `json:"materialUrl"`
	PriceInfo              interface{}   `json:"priceInfo"`
	ShopInfo               interface{}   `json:"shopInfo"`
	SkuId                  int           `json:"skuId"`
	SkuName                string        `json:"skuName"`
	Spuid                  int           `json:"spuid"`
	BrandCode              string        `json:"brandCode"`
	BrandName              string        `json:"brandName"`
	Owner                  string        `json:"owner"`
	PinGouInfo             interface{}   `json:"pinGouInfo"`
	ResourceInfo           interface{}   `json:"resourceInfo"`
	InOrderCount30DaysSku  int           `json:"inOrderCount30DaysSku"`
	SeckillInfo            interface{}   `json:"seckillInfo"`
	JxFlags                []int         `json:"jxFlags"`
	VideoInfo              interface{}   `json:"videoInfo"`
	DocumentInfo           interface{}   `json:"documentInfo"`
	BookInfo               interface{}   `json:"bookInfo"`
	ForbidTypes            []int         `json:"forbidTypes"`
	DeliveryType           int           `json:"deliveryType"`
	SkuLabelInfo           interface{}   `json:"skuLabelInfo"`
	PromotionLabelInfoList []interface{} `json:"promotionLabelInfoList"`
	SecondPriceInfoList    []interface{} `json:"secondPriceInfoList"`
	PreSaleInfo            interface{}   `json:"preSaleInfo"`
	ReserveInfo            interface{}   `json:"reserveInfo"`
}
