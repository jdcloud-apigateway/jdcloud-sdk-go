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

type DescribeDiagnoseTrendRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceGid string `json:"instanceGid"`

    /* 查询时间，格式为：2021-11-11T15:04:05Z  */
    EndTime string `json:"endTime"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceGid: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param endTime: 查询时间，格式为：2021-11-11T15:04:05Z (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDiagnoseTrendRequest(
    regionId string,
    instanceGid string,
    endTime string,
) *DescribeDiagnoseTrendRequest {

	return &DescribeDiagnoseTrendRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instance/{instanceGid}/diagnoseTrend",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceGid: instanceGid,
        EndTime: endTime,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceGid: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param endTime: 查询时间，格式为：2021-11-11T15:04:05Z (Required)
 */
func NewDescribeDiagnoseTrendRequestWithAllParams(
    regionId string,
    instanceGid string,
    endTime string,
) *DescribeDiagnoseTrendRequest {

    return &DescribeDiagnoseTrendRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/diagnoseTrend",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceGid: instanceGid,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDiagnoseTrendRequestWithoutParam() *DescribeDiagnoseTrendRequest {

    return &DescribeDiagnoseTrendRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/diagnoseTrend",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DescribeDiagnoseTrendRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceGid: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *DescribeDiagnoseTrendRequest) SetInstanceGid(instanceGid string) {
    r.InstanceGid = instanceGid
}

/* param endTime: 查询时间，格式为：2021-11-11T15:04:05Z(Required) */
func (r *DescribeDiagnoseTrendRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDiagnoseTrendRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDiagnoseTrendResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDiagnoseTrendResult `json:"result"`
}

type DescribeDiagnoseTrendResult struct {
    Data []smartdba.DataPoint `json:"data"`
}