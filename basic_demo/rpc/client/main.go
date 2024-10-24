package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"time"
)

type RespData struct {
	Data int `json:"data"`
}

func Add(a, b int) RespData {
	req := HttpRequest.NewRequest()
	req.SetTimeout(10 * time.Second)
	res, _ := req.Get(fmt.Sprintf("http://127.0.0.1:8000/add?a=%d&b=%d", a, b))
	body, _ := res.Body()
	var resp RespData
	_ = json.Unmarshal(body, &resp)
	return resp
}

func main() {
	fmt.Println(Add(3, 4))
}
