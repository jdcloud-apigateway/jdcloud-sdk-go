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

type DeleteAgRequest struct {

    core.JDCloudRequest

    /* 地域  */
    RegionId string `json:"regionId"`

    /* 高可用组 ID  */
    AgId string `json:"agId"`

}

/*
 * param regionId: 地域 (Required)
 * param agId: 高可用组 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteAgRequest(
    regionId string,
    agId string,
) *DeleteAgRequest {

	return &DeleteAgRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/availabilityGroups/{agId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AgId: agId,
	}
}

/*
 * param regionId: 地域 (Required)
 * param agId: 高可用组 ID (Required)
 */
func NewDeleteAgRequestWithAllParams(
    regionId string,
    agId string,
) *DeleteAgRequest {

    return &DeleteAgRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/availabilityGroups/{agId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AgId: agId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteAgRequestWithoutParam() *DeleteAgRequest {

    return &DeleteAgRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/availabilityGroups/{agId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域(Required) */
func (r *DeleteAgRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param agId: 高可用组 ID(Required) */
func (r *DeleteAgRequest) SetAgId(agId string) {
    r.AgId = agId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteAgRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteAgResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteAgResult `json:"result"`
}

type DeleteAgResult struct {
}