package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"log"
	"time"
)

type RequestBody struct {
	JsonRPC string   `json:"jsonrpc"`
	Id      int      `json:"id"`
	Params  []string `json:"params"`
	Method  string   `json:"method"`
}

func main() {
	reqBody := RequestBody{
		JsonRPC: "2.0",
		Id:      0,
		Params:  []string{"showyquasar"},
		Method:  "HelloServiceRPC.Hello",
	}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal("Failed to marshal JSON:", err)
	}
	resp, _ := HttpRequest.SetTimeout(10*time.Second).Post("http://localhost:1234/jsonrpc", bytes.NewBuffer(jsonData))
	fmt.Println(*resp)
}
