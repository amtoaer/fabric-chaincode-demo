package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type paymentChaincode struct{}

func checkError(errs ...error) bool {
	// 批量检查错误，返回是否出错的bool值
	var result = false
	for _, err := range errs {
		if err != nil {
			result = true
			break
		}
	}
	return result
}

func (t *paymentChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	// 需要4个参数，依次为
	// 账户1、账户1余额、账户2、账户2余额
	if len(args) != 4 {
		shim.Error("len(args) must be 4")
	}
	// 检查余额是否是整数
	_, err1 := strconv.Atoi(args[1])
	_, err2 := strconv.Atoi(args[3])
	if checkError(err1, err2) {
		return shim.Error("Error while parsing value...")
	}
	// 尝试写入状态
	err1 = stub.PutState(args[0], []byte(args[1]))
	err2 = stub.PutState(args[2], []byte(args[3]))
	// 检查写入错误
	if checkError(err1, err2) {
		return shim.Error("Error while writing data...")
	}
	// 输出成功信息并返回
	fmt.Println("init success!")
	return shim.Success([]byte("success"))
}

func (t *paymentChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// 解析得到函数名和参数
	function, args := stub.GetFunctionAndParameters()
	// switch选择调用的函数
	switch function {
	case "find":
		return find(stub, args)
	case "payment":
		return payment(stub, args)
	case "delete":
		return del(stub, args)
	default:
		return shim.Error("invalid function name...")
	}
}

func find(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// find的参数为账户名，因此长度必须为1
	if len(args) != 1 {
		return shim.Error("len(args) must be 1 when use find function")
	}
	// 获取余额并检查错误
	result, err := stub.GetState(args[0])
	if checkError(err) || result == nil {
		return shim.Error("Error while getting state...")
	}
	return shim.Success(result)
}

func payment(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// 转账，需要三个参数，分别是：
	// 账号1，账号2，金额
	if len(args) != 3 {
		return shim.Error("len(args) must be 3 when use payment function")
	}
	// 得到原始余额
	from, err1 := stub.GetState(args[0])
	to, err2 := stub.GetState(args[1])
	if checkError(err1, err2) {
		return shim.Error("Error while getting state...")
	}
	// 将原始余额转为整数
	fromBefore, _ := strconv.Atoi(string(from))
	toBefore, _ := strconv.Atoi(string(to))
	// 尝试将转账金额转为整数并检查错误
	payment, err := strconv.Atoi(args[2])
	if checkError(err) {
		return shim.Error("args[2] must be int")
	}
	// 付款账号余额减少，收款账号余额增加
	fromAfter := fromBefore - payment
	toAfter := toBefore + payment
	// 写入状态
	stub.PutState(args[0], []byte(strconv.Itoa(fromAfter)))
	stub.PutState(args[1], []byte(strconv.Itoa(toAfter)))
	return shim.Success([]byte("payment success!"))
}
func del(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// 删除账号只需要账号名这一个参数
	if len(args) != 1 {
		return shim.Error("len(args) must be 1 when use delete function")
	}
	// 尝试删除并检查结果
	err := stub.DelState(args[0])
	if checkError(err) {
		return shim.Error("Error while deleting state...")
	}
	return shim.Success([]byte("delete success!"))
}

func main() {
	err := shim.Start(new(paymentChaincode))
	if err != nil {
		fmt.Println("Error while starting chaincode...")
	}
}
