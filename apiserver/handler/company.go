package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/zhj0811/dbzl/apiserver/common"
	"github.com/zhj0811/dbzl/apiserver/sdk"
	"github.com/zhj0811/dbzl/common/define"
)

func UploadCompany(c *gin.Context) {
	req := &define.Company{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		logger.Errorf("Read request com failed %s", err.Error())
		Response(c, err, common.RequestFormatErr, nil)
		return
	}
	txId, errCode, err := invokeCompany(req)
	if err != nil {
		logger.Errorf("Invoke policy failed %s", err.Error())
		Response(c, err, errCode, nil)
		return
	}
	logger.Infof("Invoke policy %s success, tx id: %s", req.ID, txId)
	Response(c, nil, common.Success, txId)
	return
}

func invokeCompany(company *define.Company) (string, int, error) {
	args, err := json.Marshal(company)
	if err != nil {
		//logger.Errorf("Marshal policy failed %s", err.Error())
		//Response(c, err, common.RequestFormatErr, nil)
		return "", common.RequestFormatErr, err
	}
	req := []string{define.SaveCompany, string(args)}
	res, err := sdk.Invoke(req)
	if err != nil {
		return "", common.InvokeErr, err
	}
	return res.TxID, common.Success, nil
}

func QueryCompany(c *gin.Context) {
	id := c.Param("id")
	res, errCode, err := queryCompany(id)
	if err != nil {
		logger.Errorf("Fabric query company %s failed %s", id, err.Error())
		Response(c, err, errCode, nil)
		return
	}
	logger.Infof("Fabric query company success %+v", res)
	Response(c, nil, common.Success, res)
	return
}

func queryCompany(id string) (*define.CompanyInfo, int, error) {
	bytes, err := queryByKey(id)
	if err != nil {
		return nil, common.QueryErr, err
	}
	res := &define.CompanyInfo{}
	err = json.Unmarshal(bytes, res)
	if err != nil {
		return nil, common.UnmarshalJSONErr, err
	}
	filterTx, err := sdk.GetFilterTxByTxID(res.TxID)
	if err != nil {
		return nil, common.QueryErr, err
	}
	res.BlockHeight = filterTx.BlockNum
	res.Timestamp = filterTx.Timestamp
	return res, common.Success, nil
}
