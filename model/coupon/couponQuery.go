package coupon

import "github.com/open-vector/open-jd-sdk/model"

// 优惠券领取情况查询接口(申请)
// 通过领券链接查询优惠券的平台、面额、期限、可用状态、剩余数量等详细信息，通常用于和商品信息一起展示优惠券券信息。需向cps-qxsq@jd.com申请权限。

/**
request
*/

type CouponQueryRequest struct {
	CouponUrls []string `json:"couponUrls,omitempty"`
}

func (CouponQueryRequest) GetMethod() string {
	return "jd.union.open.coupon.query"
}

/**
response
*/

type CouponQueryResult struct {
	model.BaseResult
	Data []CouponResp `json:"data"`
}

type CouponResp struct {
	TakeEndTime   int    `json:"takeEndTime"`
	TakeBeginTime int    `json:"takeBeginTime"`
	RemainNum     int    `json:"remainNum"`
	Yn            string `json:"yn"`
	Num           int    `json:"num"`
	Quota         int    `json:"quota"`
	Link          string `json:"link"`
	Discount      int    `json:"discount"`
	BeginTime     int    `json:"beginTime"`
	EndTime       int    `json:"endTime"`
	Platform      string `json:"platform"`
}
