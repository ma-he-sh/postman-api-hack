package actions

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"
)

var hostfile = ".devcli"

type CodeList struct {
	Code  string `json:"code"`
	Fname string `json:"Fname"`
}

type SectionListData struct {
	Hash  string     `json:"hash"`
	Codes []CodeList `json:"codes"`
}

type SectionStatusData struct {
	Status string `json:"status"`
}

type SectionVersionData struct {
	Hash  string `json:"hash"`
	Stamp string `json:"stamp"`
}

type CodeData struct {
	Fname   string `json:"fname"`
	Code    string `json:"code"`
	Content string `json:"content"`
}

type ListPayload struct {
	Data    SectionListData `json:"data"`
	Resp    string          `json:"resp"`
	Server  string          `json:"server"`
	Type    string          `json:"type"`
	Version string          `json:"version"`
}

type StatusPayload struct {
	Data    SectionStatusData `json:"data"`
	Resp    string            `json:"resp"`
	Server  string            `json:"server"`
	Type    string            `json:"type"`
	Version string            `json:"version"`
}

type VersionPayload struct {
	Data    SectionVersionData `json:"data"`
	Resp    string             `json:"resp"`
	Server  string             `json:"server"`
	Type    string             `json:"type"`
	Version string             `json:"version"`
}

type FetchPayload struct {
	Data    CodeData `json:"data"`
	Resp    string   `json:"resp"`
	Server  string   `json:"server"`
	Type    string   `json:"type"`
	Version string   `json:"version"`
}

func getHost() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	hostfile = usr.HomeDir + "/" + hostfile

	if _, err := os.Stat(hostfile); os.IsNotExist(err) {
		log.Println(`
			>> .devcli file not found
			>> Please create a file ~/.devcli and set the hostname
		`)
	}

	content, err := ioutil.ReadFile(hostfile)
	if err != nil {
		log.Fatal(err)
	}

	host := string(content)
	return strings.Trim(host, "\n")
}

// List available gitignore files
func RequestList() ListPayload {
	host := getHost()
	resp, err := http.Get(host + "/api/list")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var result ListPayload

	json.NewDecoder(resp.Body).Decode(&result)

	return result
}

// Get Status
func RequestStatus() StatusPayload {
	host := getHost()
	resp, err := http.Get(host + "/api/status")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var result StatusPayload

	json.NewDecoder(resp.Body).Decode(&result)

	return result
}

// Get Version
func RequestVersion() VersionPayload {
	host := getHost()
	resp, err := http.Get(host + "/api/version")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var result VersionPayload

	json.NewDecoder(resp.Body).Decode(&result)

	return result
}

// Get Content
func RequestCodes(codes []string) FetchPayload {
	requestBody, err := json.Marshal(map[string]interface{}{
		"codes": codes,
	})

	log.Println(requestBody)

	if err != nil {
		log.Fatal(err)
	}

	host := getHost()
	resp, err := http.Post(host+"/api/fetch", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	log.Println(resp.Body)

	var result FetchPayload

	json.NewDecoder(resp.Body).Decode(&result)

	return result
}
