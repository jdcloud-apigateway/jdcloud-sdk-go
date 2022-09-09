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
    tidb "github.com/jdcloud-api/jdcloud-sdk-go/services/tidb/models"
)

type DescribeInstanceClassesRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 存储类型,目前只支持本地SSD;  */
    StorageType string `json:"storageType"`
}

/*
 * param regionId: 地域代码 (Required)
 * param storageType: 存储类型,目前只支持本地SSD; (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInstanceClassesRequest(
    regionId string,
    storageType string,
) *DescribeInstanceClassesRequest {

	return &DescribeInstanceClassesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances:describeInstanceClasses",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        StorageType: storageType,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param storageType: 存储类型,目前只支持本地SSD; (Required)
 */
func NewDescribeInstanceClassesRequestWithAllParams(
    regionId string,
    storageType string,
) *DescribeInstanceClassesRequest {

    return &DescribeInstanceClassesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances:describeInstanceClasses",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        StorageType: storageType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInstanceClassesRequestWithoutParam() *DescribeInstanceClassesRequest {

    return &DescribeInstanceClassesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances:describeInstanceClasses",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *DescribeInstanceClassesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param storageType: 存储类型,目前只支持本地SSD;(Required) */
func (r *DescribeInstanceClassesRequest) SetStorageType(storageType string) {
    r.StorageType = storageType
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstanceClassesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInstanceClassesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInstanceClassesResult `json:"result"`
}

type DescribeInstanceClassesResult struct {
    TidbFlavors tidb.NodeFlavor `json:"tidbFlavors"`
    TikvFlavors tidb.NodeFlavor `json:"tikvFlavors"`
    PdFlavors tidb.NodeFlavor `json:"pdFlavors"`
    MonitorFlavors tidb.NodeFlavor `json:"monitorFlavors"`
    TiflashFlavors tidb.NodeFlavor `json:"tiflashFlavors"`
    TicdcFlavors tidb.NodeFlavor `json:"ticdcFlavors"`
}