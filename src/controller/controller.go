package controller

import (
	"encoding/json"
)

type Message struct {
	Message string `json:"message"`
}

func Hello() string {
	hello := Message{ Message: "Hello World" }
	helloJson, _ := json.Marshal(hello)
	return string(helloJson)
}
