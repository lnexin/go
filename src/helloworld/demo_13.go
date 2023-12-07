package main

import (
	"encoding/json"
	"fmt"
)

type Res struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func main() {
	var res Res
	res.Code = 200
	res.Msg = "have a error message"

	toJson(&res)

	//值传递
	setData(res)
	toJson(&res)

	// 引用传递
	setData2(&res)
	toJson(&res)
}

func setData(res Res) {
	res.Code = 403
	res.Msg = "unauthiroze"
}

func setData2(res *Res) {
	res.Code = 301
	res.Msg = "forward"
}

func toJson(res *Res) {
	jtext, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshel error", errs)
	}

	fmt.Println("format json is:", string(jtext))

}
