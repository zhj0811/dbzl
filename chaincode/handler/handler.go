package handler

import (
	"encoding/json"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/zhj0811/dbzl/common/define"
)

var logger = flogging.MustGetLogger("handler")

func SavePolicy(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	logger.Infof("Enter .....%s", function)
	var policy define.Policy
	txId := stub.GetTxID()
	err := json.Unmarshal([]byte(args[0]), &policy)
	if err != nil {
		return nil, err
	}
	err = stub.PutState(txId, []byte(args[0]))
	if err != nil {
		return nil, err
	}
	//p := Policy{PId: txId}
	p, err := json.Marshal(&Policy{PId: txId})
	if err != nil {
		return nil, err
	}
	err = stub.PutState(policy.Number, p)
	if err != nil {
		return nil, err
	}
}

func SaveService(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	logger.Infof("Enter .....%s", function)
	var req define.Service
	txId := stub.GetTxID()
	err := json.Unmarshal([]byte(args[0]), &req)
	if err != nil {
		return nil, err
	}
	err = stub.PutState(txId, []byte(args[0]))
	if err != nil {
		return nil, err
	}
	//p := Policy{PId: txId}
	pBytes, err := stub.GetState(req.Number)
	if err != nil {
		return nil, err
	}
	p, err := json.Marshal(&Policy{PId: txId})
	if err != nil {
		return nil, err
	}
	err = stub.PutState(req.Number, p)
	if err != nil {
		return nil, err
	}
}
