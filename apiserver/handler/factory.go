package handler

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhj0811/dbzl/apiserver/common"
	"github.com/zhj0811/dbzl/apiserver/db"
	"github.com/zhj0811/dbzl/apiserver/sdk"
	"github.com/zhj0811/dbzl/common/define"
)

func InvokePolicy(c *gin.Context) {
	id := c.Param("id")

	txid, errCode, err := invokePolicyByNumber(id)
	if err != nil {
		logger.Errorf("Fabric invoke policy %s failed %s", id, err.Error())
		Response(c, err, errCode, nil)
		return
	}
	logger.Infof("Fabric invoke policy %s res %s", id, txid)
	Response(c, nil, common.Success, txid)
	return
}

func invokePolicyByNumber(id string) (string, int, error) {
	var policy db.TPolicy
	err := db.DB.Model(&db.TPolicy{}).First(&policy, "number = ?", id).Error
	if err != nil {
		return "", common.GetDBErr, err
	}
	args, err := json.Marshal(policy.Policy)
	if err != nil {
		return "", common.MarshalJSONErr, err
	}
	req := []string{define.SavePolicy, string(args)}
	res, err := sdk.Invoke(req)
	if err != nil {
		return "", common.InvokeErr, err
	}
	db.DB.Model(&policy).Update("tx_id", res.TxID)
	return res.TxID, common.Success, nil
}

func InvokeService(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Errorf("Fabric invoke service %s failed %s", id, err.Error())
		Response(c, err, common.RequestFormatErr, nil)
		return
	}
	txId, errCode, err := invokeServiceByID(uid)
	if err != nil {
		logger.Errorf("Fabric invoke service %s failed %s", id, err.Error())
		Response(c, err, errCode, nil)
		return
	}
	logger.Infof("Fabric invoke service %s res %s", id, txId)
	Response(c, nil, common.Success, txId)
	return
}

func invokeServiceByID(id uint64) (string, int, error) {
	var service db.TService
	err := db.DB.Model(&db.TService{}).First(&service, "id = ?", id).Error
	if err != nil {
		return "", common.GetDBErr, err
	}
	args, err := json.Marshal(service.Service)
	if err != nil {
		return "", common.MarshalJSONErr, err
	}
	req := []string{define.SaveService, string(args)}
	res, err := sdk.Invoke(req)
	if err != nil {
		return "", common.InvokeErr, err
	}
	db.DB.Model(&service).Update("tx_id", res.TxID)
	return res.TxID, common.Success, nil
}
