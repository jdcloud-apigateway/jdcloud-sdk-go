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

type ModifyResourceRecordStatusRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 域名ID，请使用describeDomains接口获取。  */
    DomainId string `json:"domainId"`

    /* 解析记录ID，请使用describeResourceRecord接口获取。  */
    ResourceRecordId string `json:"resourceRecordId"`

    /* 要修改的状态，enable:启用 disable:停用  */
    Action string `json:"action"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用describeDomains接口获取。 (Required)
 * param resourceRecordId: 解析记录ID，请使用describeResourceRecord接口获取。 (Required)
 * param action: 要修改的状态，enable:启用 disable:停用 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyResourceRecordStatusRequest(
    regionId string,
    domainId string,
    resourceRecordId string,
    action string,
) *ModifyResourceRecordStatusRequest {

	return &ModifyResourceRecordStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domainId}/ResourceRecord/{resourceRecordId}/status",
			Method:  "PUT",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        DomainId: domainId,
        ResourceRecordId: resourceRecordId,
        Action: action,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用describeDomains接口获取。 (Required)
 * param resourceRecordId: 解析记录ID，请使用describeResourceRecord接口获取。 (Required)
 * param action: 要修改的状态，enable:启用 disable:停用 (Required)
 */
func NewModifyResourceRecordStatusRequestWithAllParams(
    regionId string,
    domainId string,
    resourceRecordId string,
    action string,
) *ModifyResourceRecordStatusRequest {

    return &ModifyResourceRecordStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/ResourceRecord/{resourceRecordId}/status",
            Method:  "PUT",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        DomainId: domainId,
        ResourceRecordId: resourceRecordId,
        Action: action,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyResourceRecordStatusRequestWithoutParam() *ModifyResourceRecordStatusRequest {

    return &ModifyResourceRecordStatusRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/ResourceRecord/{resourceRecordId}/status",
            Method:  "PUT",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *ModifyResourceRecordStatusRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 域名ID，请使用describeDomains接口获取。(Required) */
func (r *ModifyResourceRecordStatusRequest) SetDomainId(domainId string) {
    r.DomainId = domainId
}

/* param resourceRecordId: 解析记录ID，请使用describeResourceRecord接口获取。(Required) */
func (r *ModifyResourceRecordStatusRequest) SetResourceRecordId(resourceRecordId string) {
    r.ResourceRecordId = resourceRecordId
}

/* param action: 要修改的状态，enable:启用 disable:停用(Required) */
func (r *ModifyResourceRecordStatusRequest) SetAction(action string) {
    r.Action = action
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyResourceRecordStatusRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyResourceRecordStatusResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyResourceRecordStatusResult `json:"result"`
}

type ModifyResourceRecordStatusResult struct {
}