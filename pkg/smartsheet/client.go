/*
 * Copyright 2020 wfleming@grumpysysadm.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package smartsheet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	// Smartsheet API endpoint
	apiEndpoint = "https://api.smartsheet.com/2.0"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	AuthToken   string
	APIEndpoint string
	HTTPClient  http.Client
}

type ErrorObject struct {
	ErrorCode string `json:"errorCode"`
	RefId     string `json:"refId"`
	Message   string `json:"message"`
}

type ResultObject struct {
	FailedItems []BulkItemFailure `json:"failedItems"` //Array of BulkItemFailure objects which represents the items that failed to be added or updated. See Partial Success for more information. Applicable only for bulk operations that support partial success
	Message     string            `json:"message"`     // Message that indicates the outcome of the request. (One of SUCCESS or PARTIAL_SUCCESS)
	ResultCode  int               `json:"resultCode"`  //0 (zero) if successful, 3 for partial success of a bulk operation.
	Version     int               `json:"version"`     // New version of the Sheet. Applicable only for operations which update sheet data
	Result      interface{}       `json:"result"`
}

type BulkItemFailure struct {
	RowId int         `json:"rowId"` // The id of the Row that failed. Applicable only to bulk row operations
	Error ErrorObject `json:"error"` // The error caused by the failed item
	Index int         `json:"index"` // The index of the failed item in the bulk request array
}

type ClientOptions struct {
	endpoint string
	token    string
	client   http.Client
}

func NewSmartsheetClient(options *ClientOptions) *Client {
	if options.token == "" {
		options.token = os.Getenv("SMARTSHEET_ACCESS_TOKEN")
	}
	return &Client{
		AuthToken:   options.token,
		APIEndpoint: options.endpoint,
		HTTPClient: http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) delete(path string) (*http.Response, error) {
	return c.do("DELETE", path, nil, nil)
}

func (c *Client) put(path string, payload interface{}, headers *map[string]string) (*http.Response, error) {
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		return c.do("PUT", path, bytes.NewBuffer(data), headers)
	}
	return c.do("PUT", path, nil, headers)
}

func (c *Client) post(path string, payload interface{}, headers *map[string]string) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return c.do("POST", path, bytes.NewBuffer(data), headers)
}

func (c *Client) get(path string) (*http.Response, error) {
	return c.do("GET", path, nil, nil)
}

func (c *Client) do(method, path string, body io.Reader, headers *map[string]string) (*http.Response, error) {
	endpoint := c.APIEndpoint + path
	req, _ := http.NewRequest(method, endpoint, body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AuthToken))
	req.Header.Set("Content-Type", "application/json")
	if headers != nil {
		for k, v := range *headers {
			req.Header.Set(k, v)
		}
	}
	resp, err := c.HTTPClient.Do(req)
	return c.checkResponse(resp, err)
}

func (c *Client) decodeJSON(resp *http.Response, payload interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(payload)
}

func (c *Client) checkResponse(resp *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return resp, fmt.Errorf("error calling the API endpoint: %v", err)
	}
	if 199 >= resp.StatusCode || 300 <= resp.StatusCode {
		var eo *ErrorObject
		var getErr error
		if eo, getErr = c.getErrorFromResponse(resp); getErr != nil {
			return resp, fmt.Errorf("response did not contain formatted error: %s. HTTP response code: %v. Raw response: %+v", getErr, resp.StatusCode, resp)
		}
		return resp, fmt.Errorf("failed call API endpoint. HTTP response code: %v. Error: %v", resp.StatusCode, eo)
	}
	return resp, nil
}

func (c *Client) getErrorFromResponse(resp *http.Response) (*ErrorObject, error) {
	var result ErrorObject
	if err := c.decodeJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", err)
	}
	return &result, nil
}
