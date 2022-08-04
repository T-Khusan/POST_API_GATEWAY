package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func MakeRequests(method, url string, reqBody []byte) (int, interface{}, error) {
	var response interface{}
	client := &http.Client{}

	reqReader := bytes.NewReader(reqBody)
	req, err := http.NewRequest(method, url, reqReader)
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Accept", "text/plain")

	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if resp == nil {
		return http.StatusInternalServerError, response, fmt.Errorf("error while connecting endpoint")
	}

	err = json.Unmarshal(body, &response)

	if resp.StatusCode > 200 {
		return resp.StatusCode, response, fmt.Errorf("connection error: response: %+v, status_code: %d", response, resp.StatusCode)
	}

	if err != nil && method != "POST" {
		fmt.Printf("ERROR WHILE MARSHALLING RESPONSE: %+v", err)
		return resp.StatusCode, string(body), err
	}
	return resp.StatusCode, response, nil
}
