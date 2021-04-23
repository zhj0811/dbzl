package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhj0811/dbzl/apiserver/common"
	"github.com/zhj0811/dbzl/apiserver/db"
)

func GetServices(c *gin.Context) {
	page := c.Query("page")
	count := c.Query("count")
	insured := c.Query("insured")
	number := c.Query("number")
	list, totalCount, err := db.GetServices(page, count, number, insured)
	if err != nil {
		logger.Errorf("Get service list failed %s", err.Error())
		Response(c, err, common.GetDBErr, nil)
		return
	}
	logger.Infof("Get service %+v", list)
	res := ListInfo{Total: totalCount, List: list}
	Response(c, nil, common.Success, res)
	return
}
