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

type UpdateCollectInfoStatusRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 采集配置 UID  */
    CollectInfoUID string `json:"collectInfoUID"`

    /* 采集状态，0-禁用，1-启用  */
    Enabled bool `json:"enabled"`
}

/*
 * param regionId: 地域 Id (Required)
 * param collectInfoUID: 采集配置 UID (Required)
 * param enabled: 采集状态，0-禁用，1-启用 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateCollectInfoStatusRequest(
    regionId string,
    collectInfoUID string,
    enabled bool,
) *UpdateCollectInfoStatusRequest {

	return &UpdateCollectInfoStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/collectinfos/{collectInfoUID}:switch",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CollectInfoUID: collectInfoUID,
        Enabled: enabled,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param collectInfoUID: 采集配置 UID (Required)
 * param enabled: 采集状态，0-禁用，1-启用 (Required)
 */
func NewUpdateCollectInfoStatusRequestWithAllParams(
    regionId string,
    collectInfoUID string,
    enabled bool,
) *UpdateCollectInfoStatusRequest {

    return &UpdateCollectInfoStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/collectinfos/{collectInfoUID}:switch",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CollectInfoUID: collectInfoUID,
        Enabled: enabled,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateCollectInfoStatusRequestWithoutParam() *UpdateCollectInfoStatusRequest {

    return &UpdateCollectInfoStatusRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/collectinfos/{collectInfoUID}:switch",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *UpdateCollectInfoStatusRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param collectInfoUID: 采集配置 UID(Required) */
func (r *UpdateCollectInfoStatusRequest) SetCollectInfoUID(collectInfoUID string) {
    r.CollectInfoUID = collectInfoUID
}
/* param enabled: 采集状态，0-禁用，1-启用(Required) */
func (r *UpdateCollectInfoStatusRequest) SetEnabled(enabled bool) {
    r.Enabled = enabled
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateCollectInfoStatusRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateCollectInfoStatusResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateCollectInfoStatusResult `json:"result"`
}

type UpdateCollectInfoStatusResult struct {
}