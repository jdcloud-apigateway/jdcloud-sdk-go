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
    smartdba "github.com/jdcloud-api/jdcloud-sdk-go/services/smartdba/models"
)

type CreateClusterRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 待接入实例gid列表  */
    GidList []string `json:"gidList"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param gidList: 待接入实例gid列表 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateClusterRequest(
    regionId string,
    gidList []string,
) *CreateClusterRequest {

	return &CreateClusterRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/newcutin",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        GidList: gidList,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param gidList: 待接入实例gid列表 (Required)
 */
func NewCreateClusterRequestWithAllParams(
    regionId string,
    gidList []string,
) *CreateClusterRequest {

    return &CreateClusterRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/newcutin",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        GidList: gidList,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateClusterRequestWithoutParam() *CreateClusterRequest {

    return &CreateClusterRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/newcutin",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *CreateClusterRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param gidList: 待接入实例gid列表(Required) */
func (r *CreateClusterRequest) SetGidList(gidList []string) {
    r.GidList = gidList
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateClusterRequest) GetRegionId() string {
    return r.RegionId
}

type CreateClusterResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateClusterResult `json:"result"`
}

type CreateClusterResult struct {
    TotalCount int `json:"totalCount"`
    Data []smartdba.InstancesInfo `json:"data"`
}