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
    iam "github.com/jdcloud-api/jdcloud-sdk-go/services/iam/models"
)

type UpdatePolicyRequest struct {

    core.JDCloudRequest

    /* 策略名称  */
    PolicyName string `json:"policyName"`

    /* 策略文档信息  */
    UpdatePolicyInfo *iam.UpdatePolicyInfo `json:"updatePolicyInfo"`
}

/*
 * param policyName: 策略名称 (Required)
 * param updatePolicyInfo: 策略文档信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdatePolicyRequest(
    policyName string,
    updatePolicyInfo *iam.UpdatePolicyInfo,
) *UpdatePolicyRequest {

	return &UpdatePolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/policy/{policyName}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        PolicyName: policyName,
        UpdatePolicyInfo: updatePolicyInfo,
	}
}

/*
 * param policyName: 策略名称 (Required)
 * param updatePolicyInfo: 策略文档信息 (Required)
 */
func NewUpdatePolicyRequestWithAllParams(
    policyName string,
    updatePolicyInfo *iam.UpdatePolicyInfo,
) *UpdatePolicyRequest {

    return &UpdatePolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/policy/{policyName}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        PolicyName: policyName,
        UpdatePolicyInfo: updatePolicyInfo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdatePolicyRequestWithoutParam() *UpdatePolicyRequest {

    return &UpdatePolicyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/policy/{policyName}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param policyName: 策略名称(Required) */
func (r *UpdatePolicyRequest) SetPolicyName(policyName string) {
    r.PolicyName = policyName
}
/* param updatePolicyInfo: 策略文档信息(Required) */
func (r *UpdatePolicyRequest) SetUpdatePolicyInfo(updatePolicyInfo *iam.UpdatePolicyInfo) {
    r.UpdatePolicyInfo = updatePolicyInfo
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdatePolicyRequest) GetRegionId() string {
    return ""
}

type UpdatePolicyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdatePolicyResult `json:"result"`
}

type UpdatePolicyResult struct {
}