package promotion

// 工具商线报推广(申请)

/**
request
*/

type ByunionidGetRequest struct {
	PromotionCodeReq PromotionCodeReq `json:"promotionCodeReq"`
}

func (ByunionidGetRequest) GetMethod() string {
	return "jd.union.open.promotion.byunionid.get"
}

/**
response
*/
