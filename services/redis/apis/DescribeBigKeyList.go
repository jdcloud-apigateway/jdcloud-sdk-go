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
    redis "github.com/jdcloud-api/jdcloud-sdk-go/services/redis/models"
)

type DescribeBigKeyListRequest struct {

    core.JDCloudRequest

    /* 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2  */
    RegionId string `json:"regionId"`

    /* 缓存Redis实例ID，是访问实例的唯一标识  */
    CacheInstanceId string `json:"cacheInstanceId"`

    /* 格式:yyyy-MM-dd,表示查询某一天的大key分析列表  */
    Date string `json:"date"`
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param date: 格式:yyyy-MM-dd,表示查询某一天的大key分析列表 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeBigKeyListRequest(
    regionId string,
    cacheInstanceId string,
    date string,
) *DescribeBigKeyListRequest {

	return &DescribeBigKeyListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/bigKey",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        Date: date,
	}
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param date: 格式:yyyy-MM-dd,表示查询某一天的大key分析列表 (Required)
 */
func NewDescribeBigKeyListRequestWithAllParams(
    regionId string,
    cacheInstanceId string,
    date string,
) *DescribeBigKeyListRequest {

    return &DescribeBigKeyListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/bigKey",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        Date: date,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeBigKeyListRequestWithoutParam() *DescribeBigKeyListRequest {

    return &DescribeBigKeyListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/bigKey",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2(Required) */
func (r *DescribeBigKeyListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识(Required) */
func (r *DescribeBigKeyListRequest) SetCacheInstanceId(cacheInstanceId string) {
    r.CacheInstanceId = cacheInstanceId
}
/* param date: 格式:yyyy-MM-dd,表示查询某一天的大key分析列表(Required) */
func (r *DescribeBigKeyListRequest) SetDate(date string) {
    r.Date = date
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeBigKeyListRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeBigKeyListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeBigKeyListResult `json:"result"`
}

type DescribeBigKeyListResult struct {
    Analyses []redis.CacheAnalysis `json:"analyses"`
}