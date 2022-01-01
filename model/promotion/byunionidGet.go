package promotion

// 社交媒体获取推广链接接口

/**
request
*/
type BysubunionidGetRequest struct {
	PromotionCodeReq PromotionCodeReq `json:"promotionCodeReq"`
}

func (BysubunionidGetRequest) GetMethod() string {
	return "jd.union.open.promotion.bysubunionid.get"
}

/**
response
*/
