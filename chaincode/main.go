package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode is a simple chaincode implementation
type SimpleChaincode struct{}

// Init method
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// 链码初始化或升级时的处理逻辑
}

// Invoke method
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// 链码运行中被调用或查询时的处理逻辑
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode : %v", err)
	}
}
