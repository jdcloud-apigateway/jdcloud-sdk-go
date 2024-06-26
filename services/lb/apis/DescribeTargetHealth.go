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
    lb "github.com/jdcloud-api/jdcloud-sdk-go/services/lb/models"
)

type DescribeTargetHealthRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Backend Id  */
    BackendId string `json:"backendId"`

    /* 页码, 默认为1,取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认返回全部，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: Region ID (Required)
 * param backendId: Backend Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTargetHealthRequest(
    regionId string,
    backendId string,
) *DescribeTargetHealthRequest {

	return &DescribeTargetHealthRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/backends/{backendId}/health",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BackendId: backendId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param backendId: Backend Id (Required)
 * param pageNumber: 页码, 默认为1,取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional)
 * param pageSize: 分页大小，默认返回全部，取值范围：[10,100] (Optional)
 */
func NewDescribeTargetHealthRequestWithAllParams(
    regionId string,
    backendId string,
    pageNumber *int,
    pageSize *int,
) *DescribeTargetHealthRequest {

    return &DescribeTargetHealthRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backends/{backendId}/health",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BackendId: backendId,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTargetHealthRequestWithoutParam() *DescribeTargetHealthRequest {

    return &DescribeTargetHealthRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backends/{backendId}/health",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeTargetHealthRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param backendId: Backend Id(Required) */
func (r *DescribeTargetHealthRequest) SetBackendId(backendId string) {
    r.BackendId = backendId
}
/* param pageNumber: 页码, 默认为1,取值范围：[1,∞), 页码超过总页数时, 显示最后一页(Optional) */
func (r *DescribeTargetHealthRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 分页大小，默认返回全部，取值范围：[10,100](Optional) */
func (r *DescribeTargetHealthRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTargetHealthRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeTargetHealthResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTargetHealthResult `json:"result"`
}

type DescribeTargetHealthResult struct {
    TargetHealths []lb.TargetHealth `json:"targetHealths"`
    TotalCount int `json:"totalCount"`
}