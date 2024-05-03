package pkg

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

// ktUrl is the base url for the ktCloud API
const ktUrl = "https://resistance.go-kt.com"

// apiUrl is the url to JSON-RPC endpoint
const apiUrl = ktUrl + "/json-rpc"

// ApiResponse is the JSON-RPC response structure
type ApiResponse struct {
	JsonRPC string `json:"jsonrpc"`
	ID      uint   `json:"id"`
	Error   struct {
		Code    uint   `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
	Result map[string]interface{} `json:"result,omitempty"`
}

// CheckApiAlive checks if the API is alive by sending a GET request to the /ping endpoint
func CheckApiAlive() bool {
	client := KtCustomClient()
	response, err := client.Get(ktUrl + "/ping")
	if err != nil {
		return false
	}
	defer response.Body.Close()

	text, err := io.ReadAll(response.Body)
	if err != nil {
		return false
	}

	return response.StatusCode == 200 || string(text) == "Pong!"
}

// ApiRequest sends a JSON-RPC request to the API. Token can be rewritten in the params map
func ApiRequest(token string, method string, params map[string]interface{}) (*ApiResponse, error) {
	if params == nil {
		params = make(map[string]interface{})
	}

	if _, ok := params["token"]; !ok {
		params["token"] = token
	}

	params = map[string]interface{}{
		"method": method,
		"params": params,
	}

	jsonData := jsonToReader(params)
	if jsonData == nil {
		return nil, errors.New("failed to convert json to reader")
	}

	requestUrl, err := url.Parse(apiUrl)
	if err != nil {
		return nil, err
	}

	client := KtCustomClient()
	response, err := client.Do(&http.Request{
		Method: "POST",
		URL:    requestUrl,
		Header: http.Header{"Content-Type": []string{"application/json-rpc"}},
		Body:   jsonData,
	})
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseData := &ApiResponse{}
	err = json.NewDecoder(response.Body).Decode(responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
