package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

const apiKeyHeader = "X-ApiKey"

//Post posts the error entry data to raygun endpoint
func Post(endpoint string, req *PostRequest, apiKey string, silent bool, debug bool) error {
	if silent {
		enc, _ := json.MarshalIndent(req, "", "\t")
		fmt.Println(string(enc))
		return nil
	}
	httpClient := &http.Client{}
	json, err := json.Marshal(req)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to convert to JSON (%s): %#v", err.Error(), req)
		return errors.New(errMsg)
	}

	r, err := http.NewRequest("POST", endpoint+"/entries", bytes.NewBuffer(json))
	if err != nil {
		errMsg := fmt.Sprintf("Unable to create request (%s)", err.Error())
		return errors.New(errMsg)
	}
	r.Header.Add(apiKeyHeader, apiKey)
	resp, err := httpClient.Do(r)

	defer resp.Body.Close()

	if resp.StatusCode == 202 {
		if debug {
			log.Println("Successfully sent message to Raygun")
		}
		return nil
	}

	errMsg := fmt.Sprintf("Unexpected answer from Raygun %d", resp.StatusCode)
	if err != nil {
		errMsg = fmt.Sprintf("%s: %s", errMsg, err.Error())
	}

	return errors.New(errMsg)
}
