package model

/**
request
*/

type JFQueryRequest struct {
	GoodReq JFGoodsReq `json:"goodsReq"`
	Result  JingfenQueryResult
}

func (jFQueryRequest JFQueryRequest) GetMethod() string {
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
	EliteId   int `json:"eliteId"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
	/**
	排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30DaysSku：sku维度30天引单量，
	comments：评论数，goodComments：好评数)
	*/
	SortName     string `json:"sortName"`
	Sort         string `json:"sort"`
	Pid          string `json:"pid"`
	Fields       string `json:"fields"`
	ForbidTypes  string `json:"forbidTypes"`
	GroupId      int    `json:"groupId"`
	OwnerUnionId int    `json:"ownerUnionId"`
}

/**
response
*/

// todo 待优化 jd_union_open_goods_jingfen_query_responce这个需要抽出来 map不知道行不行
type JFQueryResponse struct {
	Body struct {
		Code string `json:"code"`
		// 这里想直接拿JingfenQueryResult接，接不到，必须拿string先接然后再转
		QueryResult string `json:"queryResult"`
	} `json:"jd_union_open_goods_jingfen_query_responce"`
}

type JingfenQueryResult struct {
	Code       int           `json:"code"`
	Data       []JFGoodsResp `json:"data"`
	Message    string        `json:"message"`
	RequestId  string        `json:"requestId"`
	TotalCount int           `json:"totalCount"`
}

type JFGoodsResp struct {
	CategoryInfo           struct{}   `json:"categoryInfo"`
	Comments               int        `json:"comments"`
	CommissionInfo         struct{}   `json:"commission_info"`
	CouponInfo             struct{}   `json:"couponInfo"`
	GoodCommentsShare      float32    `json:"goodCommentsShare"`
	ImageInfo              struct{}   `json:"imageInfo"`
	InOrderCount30Days     int        `json:"inOrderCount30Days"`
	MaterialUrl            string     `json:"materialUrl"`
	PriceInfo              struct{}   `json:"priceInfo"`
	ShopInfo               struct{}   `json:"shopInfo"`
	SkuId                  int        `json:"skuId"`
	SkuName                string     `json:"skuName"`
	Spuid                  int        `json:"spuid"`
	BrandCode              string     `json:"brandCode"`
	BrandName              string     `json:"brandName"`
	Owner                  string     `json:"owner"`
	PinGouInfo             struct{}   `json:"pinGouInfo"`
	ResourceInfo           struct{}   `json:"resourceInfo"`
	InOrderCount30DaysSku  int        `json:"inOrderCount30DaysSku"`
	SeckillInfo            struct{}   `json:"seckillInfo"`
	JxFlags                []int      `json:"jxFlags"`
	VideoInfo              struct{}   `json:"videoInfo"`
	DocumentInfo           struct{}   `json:"documentInfo"`
	BookInfo               struct{}   `json:"bookInfo"`
	ForbidTypes            []int      `json:"forbidTypes"`
	DeliveryType           int        `json:"deliveryType"`
	SkuLabelInfo           struct{}   `json:"skuLabelInfo"`
	PromotionLabelInfoList []struct{} `json:"promotionLabelInfoList"`
	SecondPriceInfoList    []struct{} `json:"secondPriceInfoList"`
	PreSaleInfo            struct{}   `json:"preSaleInfo"`
	ReserveInfo            struct{}   `json:"reserveInfo"`
}
