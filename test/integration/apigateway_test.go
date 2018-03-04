package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestApiGatewayIntegration(t *testing.T) {
	resp, err := http.Get("https://helloworld.dev.danielspeixoto.com/hello")
	if err != nil {
		t.Fatal(err)
	}
	// Closes Body after this function finishes
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Error("Status Code: " + string(resp.StatusCode))
	}
	if string(body) != "Hello World" {
		t.Error("Response was = " + string(body))
	}
	
}
