package activity

import "github.com/open-vector/open-jd-sdk/model"

// 活动查询接口

/**
request
*/

type ActivityQuery struct {
	activityReq ActivityReq `json:"activityReq"`
}

func (ActivityQuery) GetMethod() string {
	return "jd.union.open.activity.query"
}

type ActivityReq struct {
	PageIndex  int    `json:"pageIndex"`
	PageSize   int    `json:"pageSize"`
	PoolId     int    `json:"PoolId"`
	ActiveDate string `json:"activeDate"`
}

/**
response
*/

type QueryResult struct {
	model.BaseResult
	data       []ActivityResp `json:"data"`
	TotalCount int            `json:"totalCount"`
}

type ActivityResp struct {
	UrlPC              string      `json:"urlPC"`
	UrlM               string      `json:"urlM"`
	Title              string      `json:"title"`
	Advantage          string      `json:"advantage"`
	StartTime          int         `json:"startTime"`
	EndTime            int         `json:"endTime"`
	EliteId            int         `json:"eliteId"`
	Recommend          int         `json:"recommend"`
	DownloadUrl        string      `json:"downloadUrl"`
	DownloadCode       string      `json:"downloadCode"`
	UpdateTime         int         `json:"updateTime"`
	Content            string      `json:"content"`
	Tag                string      `json:"tag"`
	ActStatus          int         `json:"actStatus"`
	PromotionStartTime int         `json:"promotionStartTime"`
	PromotionEndTime   int         `json:"promotionEndTime"`
	PlatformType       int         `json:"platformType"`
	ImgList            interface{} `json:"imgList"`
	Id                 int         `json:"id"`
	CategoryList       interface{} `json:"categoryList"`
}
