package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/zhj0811/dbzl/apiserver/common"
)

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
