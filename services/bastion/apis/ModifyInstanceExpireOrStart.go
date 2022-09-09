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

type ModifyInstanceExpireOrStartRequest struct {

    core.JDCloudRequest

    /* regionId  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 实例到期或者恢复实例服务 1:恢复实例服务 13:设置服务到期状态  */
    Status int `json:"status"`
}

/*
 * param regionId: regionId (Required)
 * param instanceId: 实例ID (Required)
 * param status: 实例到期或者恢复实例服务 1:恢复实例服务 13:设置服务到期状态 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyInstanceExpireOrStartRequest(
    regionId string,
    instanceId string,
    status int,
) *ModifyInstanceExpireOrStartRequest {

	return &ModifyInstanceExpireOrStartRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceExpireOrStart",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        Status: status,
	}
}

/*
 * param regionId: regionId (Required)
 * param instanceId: 实例ID (Required)
 * param status: 实例到期或者恢复实例服务 1:恢复实例服务 13:设置服务到期状态 (Required)
 */
func NewModifyInstanceExpireOrStartRequestWithAllParams(
    regionId string,
    instanceId string,
    status int,
) *ModifyInstanceExpireOrStartRequest {

    return &ModifyInstanceExpireOrStartRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceExpireOrStart",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        Status: status,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyInstanceExpireOrStartRequestWithoutParam() *ModifyInstanceExpireOrStartRequest {

    return &ModifyInstanceExpireOrStartRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceExpireOrStart",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: regionId(Required) */
func (r *ModifyInstanceExpireOrStartRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: 实例ID(Required) */
func (r *ModifyInstanceExpireOrStartRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param status: 实例到期或者恢复实例服务 1:恢复实例服务 13:设置服务到期状态(Required) */
func (r *ModifyInstanceExpireOrStartRequest) SetStatus(status int) {
    r.Status = status
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyInstanceExpireOrStartRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyInstanceExpireOrStartResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyInstanceExpireOrStartResult `json:"result"`
}

type ModifyInstanceExpireOrStartResult struct {
    Result bool `json:"result"`
}