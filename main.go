package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url := "http://localhost:3000/api/token"
	/*
		values := map[string]string{"key": "ed57a0a7747b90df38f97f2f8e27a104249899a5639490c33821e4ca8b329e1f",
			"channel":   "mobile",
			"timestamp": "1234567890"}*/

	var values = []byte(`{"channel":"mobile","timestamp":"1234567890","key":"ed57a0a7747b90df38f97f2f8e27a104249899a5639490c33821e4ca8b329e1f"}`)

	jsonValue, _ := json.Marshal(values)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-type", "application/json")

	fmt.Println(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println(resp.Header)

	//jsonString := []byte(`{"key":"ed57a0a7747b90df38f97f2f8e27a104249899a5639490c33821e4ca8b329e1f"}`)
}
