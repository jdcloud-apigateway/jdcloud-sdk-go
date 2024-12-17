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

type MobileStatusRequest struct {

    core.JDCloudRequest

    /* 手机号码  */
    Mobile string `json:"mobile"`
}

/*
 * param mobile: 手机号码 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewMobileStatusRequest(
    mobile string,
) *MobileStatusRequest {

	return &MobileStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/mobile:status",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Mobile: mobile,
	}
}

/*
 * param mobile: 手机号码 (Required)
 */
func NewMobileStatusRequestWithAllParams(
    mobile string,
) *MobileStatusRequest {

    return &MobileStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/mobile:status",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Mobile: mobile,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewMobileStatusRequestWithoutParam() *MobileStatusRequest {

    return &MobileStatusRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/mobile:status",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param mobile: 手机号码(Required) */
func (r *MobileStatusRequest) SetMobile(mobile string) {
    r.Mobile = mobile
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r MobileStatusRequest) GetRegionId() string {
    return ""
}

type MobileStatusResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result MobileStatusResult `json:"result"`
}

type MobileStatusResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
    ChargeFlag string `json:"chargeFlag"`
    Area string `json:"area"`
    Operator string `json:"operator"`
    Status string `json:"status"`
    MnpStatus string `json:"mnpStatus"`
}