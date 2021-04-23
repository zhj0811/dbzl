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

func QueryPolicyByNumber(c *gin.Context) {
	number := c.Param("id")
	res, errCode, err := queryPolicyByNumber(number)
	if err != nil {
		logger.Errorf("Fabric query policy %s failed %s", number, err.Error())
		Response(c, err, errCode, nil)
		return
	}
	logger.Infof("Fabric query policy %s res %+v", number, res)
	Response(c, nil, common.Success, res)
	return
}

func queryPolicyByNumber(number string) (*PolicyFullInfo, int, error) {
	policy, err := db.GetPolicyByNumber(number)
	if err != nil {
		return nil, common.GetPolicyErr, err
	}
	if (policy.TxID != "") && (policy.BlockHeight == 0) {
		policy.BlockHeight, err = sdk.GetBlockHeightByTxID(policy.TxID)
		if err != nil {
			return nil, common.QueryErr, err
		}
		db.DB.Model(&policy).Update("block_height", policy.BlockHeight)
	}
	services, err := db.GetServicesByNumber(number)
	if err != nil {
		return nil, common.GetServiceErr, err
	}
	for _, service := range services {
		if (service.TxID != "") && (service.BlockHeight == 0) {
			service.BlockHeight, err = sdk.GetBlockHeightByTxID(service.TxID)
			if err != nil {
				return nil, common.QueryErr, err
			}
			db.DB.Model(&service).Update("block_height", service.BlockHeight)
		}
	}
	res := &PolicyFullInfo{
		Policy:   *policy,
		Services: services,
	}
	return res, common.Success, nil
}
