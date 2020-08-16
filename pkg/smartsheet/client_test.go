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
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewSmartsheetClientOptions(t *testing.T) {
	os.Setenv("SMARTSHEET_ACCESS_TOKEN", "123")
	options := ClientOptions{
		endpoint: "test-endpoint.com",
		token:    "123-token",
	}
	client := NewSmartsheetClient(&options)
	assert.Equal(t, client.AuthToken, options.token)
	assert.Equal(t, client.APIEndpoint, options.endpoint)

}

func TestNewSmartsheetClientOptionsEnv(t *testing.T) {
	os.Setenv("SMARTSHEET_ACCESS_TOKEN", "123")
	options := ClientOptions{}
	client := NewSmartsheetClient(&options)
	assert.Equal(t, client.AuthToken, "123")
}
