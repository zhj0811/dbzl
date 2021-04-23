package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/zhj0811/dbzl/apiserver/common"
	"github.com/zhj0811/dbzl/apiserver/db"
)

type ListInfo struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

type PolicyFullInfo struct {
	Policy   db.TPolicy     `json:"policy"`
	Services []*db.TService `json:"services"`
}

func Response(c *gin.Context, err error, errCode int, data interface{}) {
	res := &common.ResponseInfo{
		ErrCode: errCode,
		Data:    data,
	}
	if err != nil {
		res.ErrMsg = err.Error()
	}
	ret, _ := json.Marshal(res)
	c.Writer.Write(ret)
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	return
}
