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
    invoice "github.com/jdcloud-api/jdcloud-sdk-go/services/invoice/models"
)

type DescribeAddressesRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 用户pin(运营后台必传) (Optional) */
    Pin *string `json:"pin"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAddressesRequest(
    regionId string,
) *DescribeAddressesRequest {

	return &DescribeAddressesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/addresses",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 * param pin: 用户pin(运营后台必传) (Optional)
 */
func NewDescribeAddressesRequestWithAllParams(
    regionId string,
    pin *string,
) *DescribeAddressesRequest {

    return &DescribeAddressesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/addresses",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        Pin: pin,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAddressesRequestWithoutParam() *DescribeAddressesRequest {

    return &DescribeAddressesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/addresses",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: (Required) */
func (r *DescribeAddressesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param pin: 用户pin(运营后台必传)(Optional) */
func (r *DescribeAddressesRequest) SetPin(pin string) {
    r.Pin = &pin
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAddressesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAddressesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAddressesResult `json:"result"`
}

type DescribeAddressesResult struct {
    Addresses []invoice.PostAddress `json:"addresses"`
}