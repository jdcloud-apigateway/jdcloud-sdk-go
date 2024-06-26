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

type DeleteSecurityPolicyRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Security Policy Id  */
    SecurityPolicyId string `json:"securityPolicyId"`
}

/*
 * param regionId: Region ID (Required)
 * param securityPolicyId: Security Policy Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteSecurityPolicyRequest(
    regionId string,
    securityPolicyId string,
) *DeleteSecurityPolicyRequest {

	return &DeleteSecurityPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/securityPolicies/{securityPolicyId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SecurityPolicyId: securityPolicyId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param securityPolicyId: Security Policy Id (Required)
 */
func NewDeleteSecurityPolicyRequestWithAllParams(
    regionId string,
    securityPolicyId string,
) *DeleteSecurityPolicyRequest {

    return &DeleteSecurityPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/securityPolicies/{securityPolicyId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SecurityPolicyId: securityPolicyId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteSecurityPolicyRequestWithoutParam() *DeleteSecurityPolicyRequest {

    return &DeleteSecurityPolicyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/securityPolicies/{securityPolicyId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteSecurityPolicyRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param securityPolicyId: Security Policy Id(Required) */
func (r *DeleteSecurityPolicyRequest) SetSecurityPolicyId(securityPolicyId string) {
    r.SecurityPolicyId = securityPolicyId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteSecurityPolicyRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteSecurityPolicyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteSecurityPolicyResult `json:"result"`
}

type DeleteSecurityPolicyResult struct {
}