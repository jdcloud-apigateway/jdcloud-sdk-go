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

type DescribeStoragesRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例。  */
    InstanceGid string `json:"instanceGid"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceGid: RDS 实例ID，唯一标识一个RDS实例。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeStoragesRequest(
    regionId string,
    instanceGid string,
) *DescribeStoragesRequest {

	return &DescribeStoragesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instance/{instanceGid}/describeStorages",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceGid: instanceGid,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceGid: RDS 实例ID，唯一标识一个RDS实例。 (Required)
 */
func NewDescribeStoragesRequestWithAllParams(
    regionId string,
    instanceGid string,
) *DescribeStoragesRequest {

    return &DescribeStoragesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/describeStorages",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceGid: instanceGid,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeStoragesRequestWithoutParam() *DescribeStoragesRequest {

    return &DescribeStoragesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/describeStorages",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DescribeStoragesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceGid: RDS 实例ID，唯一标识一个RDS实例。(Required) */
func (r *DescribeStoragesRequest) SetInstanceGid(instanceGid string) {
    r.InstanceGid = instanceGid
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeStoragesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeStoragesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeStoragesResult `json:"result"`
}

type DescribeStoragesResult struct {
    TotalDiskSize string `json:"totalDiskSize"`
    UsedDiskSize string `json:"usedDiskSize"`
    TotalIncreaseSize string `json:"totalIncreaseSize"`
    AverageIncrease string `json:"averageIncrease"`
    FreeDiskSize string `json:"freeDiskSize"`
    PredictAvailableDays string `json:"predictAvailableDays"`
}