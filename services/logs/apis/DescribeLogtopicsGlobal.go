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
    logs "github.com/jdcloud-api/jdcloud-sdk-go/services/logs/models"
)

type DescribeLogtopicsGlobalRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 过滤条件，key，Values， 合法的key：logtopicName， logtopicUID， logsetName， logsetUID (Optional) */
    Filters []logs.Filter `json:"filters"`

    /* 过滤条件，key，Values (Optional) */
    Tags []logs.TagFilter `json:"tags"`

    /* 日志主题采集的日志类型 (Optional) */
    AppName *string `json:"appName"`
}

/*
 * param regionId: 地域 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLogtopicsGlobalRequest(
    regionId string,
) *DescribeLogtopicsGlobalRequest {

	return &DescribeLogtopicsGlobalRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/logtopics",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param filters: 过滤条件，key，Values， 合法的key：logtopicName， logtopicUID， logsetName， logsetUID (Optional)
 * param tags: 过滤条件，key，Values (Optional)
 * param appName: 日志主题采集的日志类型 (Optional)
 */
func NewDescribeLogtopicsGlobalRequestWithAllParams(
    regionId string,
    filters []logs.Filter,
    tags []logs.TagFilter,
    appName *string,
) *DescribeLogtopicsGlobalRequest {

    return &DescribeLogtopicsGlobalRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logtopics",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Filters: filters,
        Tags: tags,
        AppName: appName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLogtopicsGlobalRequestWithoutParam() *DescribeLogtopicsGlobalRequest {

    return &DescribeLogtopicsGlobalRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logtopics",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeLogtopicsGlobalRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param filters: 过滤条件，key，Values， 合法的key：logtopicName， logtopicUID， logsetName， logsetUID(Optional) */
func (r *DescribeLogtopicsGlobalRequest) SetFilters(filters []logs.Filter) {
    r.Filters = filters
}
/* param tags: 过滤条件，key，Values(Optional) */
func (r *DescribeLogtopicsGlobalRequest) SetTags(tags []logs.TagFilter) {
    r.Tags = tags
}
/* param appName: 日志主题采集的日志类型(Optional) */
func (r *DescribeLogtopicsGlobalRequest) SetAppName(appName string) {
    r.AppName = &appName
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLogtopicsGlobalRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeLogtopicsGlobalResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLogtopicsGlobalResult `json:"result"`
}

type DescribeLogtopicsGlobalResult struct {
    Data []logs.LogtopicBaseEnd `json:"data"`
}