package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func main() {
	var result Result
	result.Code = 200
	result.Message = "success"
	toJson(&result)

	// 序列化
	jsons, errs := json.Marshal(result)
	if errs != nil {
		fmt.Println("json marshl error:", errs)
	}
	fmt.Println("json data:", string(jsons))

	// 反序列化
	var res2 Result
	errs = json.Unmarshal(jsons, &res2)
	if errs != nil {
		fmt.Println("json unmarshal error:", errs)
	}
	fmt.Println("res2", res2)

	// 改变数据
	setData(&result)
	toJson(&result)
}

func setData(res *Result) {
	res.Code = 500
	res.Message = "fail"
}

func toJson(res *Result) {
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshl error:", errs)
	}
	fmt.Println("json data:", string(jsons))
}
