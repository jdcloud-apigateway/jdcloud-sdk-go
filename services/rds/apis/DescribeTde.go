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
    rds "github.com/jdcloud-api/jdcloud-sdk-go/services/rds/models"
)

type DescribeTdeRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTdeRequest(
    regionId string,
    instanceId string,
) *DescribeTdeRequest {

	return &DescribeTdeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/tde",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 */
func NewDescribeTdeRequestWithAllParams(
    regionId string,
    instanceId string,
) *DescribeTdeRequest {

    return &DescribeTdeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/tde",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTdeRequestWithoutParam() *DescribeTdeRequest {

    return &DescribeTdeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/tde",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DescribeTdeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *DescribeTdeRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTdeRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeTdeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTdeResult `json:"result"`
}

type DescribeTdeResult struct {
    Status string `json:"status"`
    StatusDetail rds.StatusDetail `json:"statusDetail"`
}