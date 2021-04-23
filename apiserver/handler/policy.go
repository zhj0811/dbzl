package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhj0811/dbzl/apiserver/common"
	"github.com/zhj0811/dbzl/apiserver/db"
)

func GetPolicies(c *gin.Context) {
	page := c.Query("page")
	count := c.Query("count")
	insured := c.Query("insured")
	list, totalCount, err := db.GetPolicies(page, count, insured)
	if err != nil {
		logger.Errorf("Get policy list failed %s", err.Error())
		Response(c, err, common.GetDBErr, nil)
		return
	}
	logger.Infof("Get policies %+v", list)
	res := ListInfo{Total: totalCount, List: list}
	Response(c, nil, common.Success, res)
	return
}
