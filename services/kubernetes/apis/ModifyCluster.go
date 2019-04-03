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

type ModifyClusterRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 集群 ID  */
    ClusterId string `json:"clusterId"`

    /* 集群名称 (Optional) */
    Name *string `json:"name"`

    /* 集群 name 和 description 必须要指定一个 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: 地域 ID (Required)
 * param clusterId: 集群 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyClusterRequest(
    regionId string,
    clusterId string,
) *ModifyClusterRequest {

	return &ModifyClusterRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/clusters/{clusterId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ClusterId: clusterId,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param clusterId: 集群 ID (Required)
 * param name: 集群名称 (Optional)
 * param description: 集群 name 和 description 必须要指定一个 (Optional)
 */
func NewModifyClusterRequestWithAllParams(
    regionId string,
    clusterId string,
    name *string,
    description *string,
) *ModifyClusterRequest {

    return &ModifyClusterRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/clusters/{clusterId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClusterId: clusterId,
        Name: name,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyClusterRequestWithoutParam() *ModifyClusterRequest {

    return &ModifyClusterRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/clusters/{clusterId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *ModifyClusterRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param clusterId: 集群 ID(Required) */
func (r *ModifyClusterRequest) SetClusterId(clusterId string) {
    r.ClusterId = clusterId
}

/* param name: 集群名称(Optional) */
func (r *ModifyClusterRequest) SetName(name string) {
    r.Name = &name
}

/* param description: 集群 name 和 description 必须要指定一个(Optional) */
func (r *ModifyClusterRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyClusterRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyClusterResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyClusterResult `json:"result"`
}

type ModifyClusterResult struct {
}