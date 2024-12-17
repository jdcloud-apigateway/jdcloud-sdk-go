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

type ModifyInstanceReleaseTimeRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 需要更改定时删除信息的云主机id列表。
 (Optional) */
    InstanceIds []string `json:"instanceIds"`

    /* ｜ 为云主机设置或修改定时删除时间，如果为空，若云主机存在定时删除任务，则取消其定时删除任务。 支持的时间范围为当前时间+1小时至当前时间+10年。示例："2025-01-01 00:00:00" (Optional) */
    AutoReleaseTime *string `json:"autoReleaseTime"`
}

/*
 * param regionId: 地域ID。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyInstanceReleaseTimeRequest(
    regionId string,
) *ModifyInstanceReleaseTimeRequest {

	return &ModifyInstanceReleaseTimeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances:modifyInstanceReleaseTime",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param instanceIds: 需要更改定时删除信息的云主机id列表。
 (Optional)
 * param autoReleaseTime: ｜ 为云主机设置或修改定时删除时间，如果为空，若云主机存在定时删除任务，则取消其定时删除任务。 支持的时间范围为当前时间+1小时至当前时间+10年。示例："2025-01-01 00:00:00" (Optional)
 */
func NewModifyInstanceReleaseTimeRequestWithAllParams(
    regionId string,
    instanceIds []string,
    autoReleaseTime *string,
) *ModifyInstanceReleaseTimeRequest {

    return &ModifyInstanceReleaseTimeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances:modifyInstanceReleaseTime",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceIds: instanceIds,
        AutoReleaseTime: autoReleaseTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyInstanceReleaseTimeRequestWithoutParam() *ModifyInstanceReleaseTimeRequest {

    return &ModifyInstanceReleaseTimeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances:modifyInstanceReleaseTime",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *ModifyInstanceReleaseTimeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceIds: 需要更改定时删除信息的云主机id列表。
(Optional) */
func (r *ModifyInstanceReleaseTimeRequest) SetInstanceIds(instanceIds []string) {
    r.InstanceIds = instanceIds
}
/* param autoReleaseTime: ｜ 为云主机设置或修改定时删除时间，如果为空，若云主机存在定时删除任务，则取消其定时删除任务。 支持的时间范围为当前时间+1小时至当前时间+10年。示例："2025-01-01 00:00:00"(Optional) */
func (r *ModifyInstanceReleaseTimeRequest) SetAutoReleaseTime(autoReleaseTime string) {
    r.AutoReleaseTime = &autoReleaseTime
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyInstanceReleaseTimeRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyInstanceReleaseTimeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyInstanceReleaseTimeResult `json:"result"`
}

type ModifyInstanceReleaseTimeResult struct {
}