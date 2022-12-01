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
    yunding "github.com/jdcloud-api/jdcloud-sdk-go/services/yunding/models"
)

type DescribeYdRdsInstancesRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 应用appKey;  */
    AppKey string `json:"appKey"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param appKey: 应用appKey; (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeYdRdsInstancesRequest(
    regionId string,
    appKey string,
) *DescribeYdRdsInstancesRequest {

	return &DescribeYdRdsInstancesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/rdsInstances",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppKey: appKey,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param appKey: 应用appKey; (Required)
 */
func NewDescribeYdRdsInstancesRequestWithAllParams(
    regionId string,
    appKey string,
) *DescribeYdRdsInstancesRequest {

    return &DescribeYdRdsInstancesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rdsInstances",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppKey: appKey,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeYdRdsInstancesRequestWithoutParam() *DescribeYdRdsInstancesRequest {

    return &DescribeYdRdsInstancesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rdsInstances",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DescribeYdRdsInstancesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appKey: 应用appKey;(Required) */
func (r *DescribeYdRdsInstancesRequest) SetAppKey(appKey string) {
    r.AppKey = appKey
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeYdRdsInstancesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeYdRdsInstancesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeYdRdsInstancesResult `json:"result"`
}

type DescribeYdRdsInstancesResult struct {
    Clusters []yunding.RdsInstance `json:"clusters"`
    TotalCount int `json:"totalCount"`
}