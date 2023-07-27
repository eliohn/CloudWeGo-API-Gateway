package idlManager

import (
	"encoding/json"
	"io"
	"net/http"
)

var idlQueryUrl = "http://localhost:8889/idl/query"

type IDLInfo struct {
	Name string `json:"name"`
	Idl  string `json:"idl"`
}

// GetIDLContent get the idl content according to service name
func GetIDLContent(svcName string) string {
	request, err := http.NewRequest("GET", idlQueryUrl+"?name="+svcName, nil)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic("Error: failed to query IDL---" + err.Error())
	}
	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic("Error: failed to read body---" + err.Error())
	}
	var resp IDLInfo
	err = json.Unmarshal(body, &resp)
	if err != nil {
		panic("Error: failed to transform byte body into json---" + err.Error())
	}
	return resp.Idl
}
