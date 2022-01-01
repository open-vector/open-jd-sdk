package promotion

import "github.com/open-vector/open-jd-sdk/model"

// 网站/APP获取推广链接接口

/**
request
*/

type CommonGetRequest struct {
	PromotionCodeReq PromotionCodeReq `json:"promotionCodeReq"`
}

func (CommonGetRequest) GetMethod() string {
	return "jd.union.open.promotion.common.get"
}

type PromotionCodeReq struct {
	MaterialId    string `json:"materialId,omitempty"`
	SiteId        string `json:"siteId,omitempty"`
	PositionId    int    `json:"positionId,omitempty"`
	SubUnionId    string `json:"subUnionId,omitempty"`
	Ext1          string `json:"ext1,omitempty"`
	Protocol      int    `json:"protocol,omitempty"`
	Pid           string `json:"pid,omitempty"`
	CouponUrl     string `json:"couponUrl,omitempty"`
	GiftCouponKey string `json:"giftCouponKey,omitempty"`
	ChannelId     int    `json:"channelId,omitempty"`
}

/**
response
*/

type GetResult struct {
	model.BaseResult
	Data PromotionCodeResp `json:"data"`
}

type PromotionCodeResp struct {
	ClickURL      string `json:"clickURL"`
	JCommand      string `json:"JCommand"`
	ShortURL      string `json:"shortURL"`
	JShortCommand string `json:"jShortCommand"`
}
