package category

import (
	"github.com/open-vector/open-jd-sdk/model"
)

// 商品类目查询接口
// 根据商品的父类目id查询子类目id信息，通常用获取各级类目对应关系，以便将推广商品归类。业务参数parentId、grade都输入0可查询所有一级类目ID，
// 之后再用其作为parentId查询其子类目。

/**
request
*/

type CateGoryQuery struct {
	Req CategoryReq `json:"req"`
}

func (CateGoryQuery) GetMethod() string {
	return "jd.union.open.category.goods.get"
}

type CategoryReq struct {
	ParentId int `json:"ParentId"`
	Grade    int `json:"grade"`
}

/**
response
*/

type GetResult struct {
	model.BaseResult
	Data []CategoryResp `json:"data"`
}

type CategoryResp struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Grade    int    `json:"grade"`
	ParentId int    `json:"parentId"`
}
