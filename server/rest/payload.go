package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Payload struct {
	Stamp     string  `json:"stamp"`
	Hash      string  `json:"hash"`
	Languages []Codes `json:"lang"`
}

type Codes struct {
	Fname   string `json:"fname"`
	Code    string `json:"code"`
	Content string `json:"content"`
}

type ListOfCodes struct {
	Hash  string
	Codes []map[string]interface{}
}

// check file exists
func fileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// get payload file path
func GetPayloadPath() string {
	return RESTDataPath() + "/" + RESTDataFileName()
}

// get payload
func GetPayload() Payload {
	var payload Payload

	payloadFileDir := GetPayloadPath()

	if fileExists(payloadFileDir) {
		jsonFile, err := os.Open(payloadFileDir)
		if err != nil {
			fmt.Println(err)
			return payload
		}

		defer jsonFile.Close()

		byteVal, _ := ioutil.ReadAll(jsonFile)

		json.Unmarshal(byteVal, &payload)
		return payload
	}

	return payload
}

// get a list of codes
func GetListOfLang() ListOfCodes {
	var langList ListOfCodes
	payload := GetPayload()

	langList.Hash = payload.Hash
	for i := 0; i < len(payload.Languages); i++ {
		ss := map[string]interface{}{
			"code":  payload.Languages[i].Code,
			"fname": payload.Languages[i].Fname,
		}
		langList.Codes = append(langList.Codes, ss)
	}

	return langList
}
