// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type ValidateSiteMonitorAddressRequest struct {

    core.JDCloudRequest

    /*  (Optional) */
    Address *string `json:"address"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewValidateSiteMonitorAddressRequest(
) *ValidateSiteMonitorAddressRequest {

	return &ValidateSiteMonitorAddressRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/siteMonitorAddress:validate",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
	}
}

/*
 * param address:  (Optional)
 */
func NewValidateSiteMonitorAddressRequestWithAllParams(
    address *string,
) *ValidateSiteMonitorAddressRequest {

    return &ValidateSiteMonitorAddressRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/siteMonitorAddress:validate",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        Address: address,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewValidateSiteMonitorAddressRequestWithoutParam() *ValidateSiteMonitorAddressRequest {

    return &ValidateSiteMonitorAddressRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/siteMonitorAddress:validate",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param address: (Optional) */
func (r *ValidateSiteMonitorAddressRequest) SetAddress(address string) {
    r.Address = &address
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ValidateSiteMonitorAddressRequest) GetRegionId() string {
    return ""
}

type ValidateSiteMonitorAddressResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ValidateSiteMonitorAddressResult `json:"result"`
}

type ValidateSiteMonitorAddressResult struct {
    Suc bool `json:"suc"`
}