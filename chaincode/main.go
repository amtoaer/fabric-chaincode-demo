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
	fmt.Println("init chaincode...")
	_, args := stub.GetFunctionAndParameters()
	if len(args) != 2 {
		return shim.Error("Error,len(args) must be 2")
	}
	fmt.Println("try to save data...")
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error("Error while saving data...")
	}
	fmt.Println("success!")
	return shim.Success(nil)
}

// Invoke method
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// 链码运行中被调用或查询时的处理逻辑
	function, args := stub.GetFunctionAndParameters()

	if function == "query" {
		return query(stub, args)
	}
	return shim.Error("Error,invaild operation...")
}

func query(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Error,len(args) must be 1")
	}
	result, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Error when querying data...")
	}
	if result == nil {
		return shim.Error("Error,no data found...")
	}
	return shim.Success(result)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode : %v", err)
	}
}
