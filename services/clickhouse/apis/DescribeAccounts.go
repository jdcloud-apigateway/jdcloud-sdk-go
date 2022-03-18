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
    clickhouse "github.com/jdcloud-api/jdcloud-sdk-go/services/clickhouse/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeAccountsRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页; (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页显示的数据条数，默认为10，取值范围：[10,100]，用于查询列表的接口 (Optional) */
    PageSize *int `json:"pageSize"`

    /* accountStatus, 支持operator选项：eq
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAccountsRequest(
    regionId string,
    instanceId string,
) *DescribeAccountsRequest {

	return &DescribeAccountsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/accounts",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页; (Optional)
 * param pageSize: 每页显示的数据条数，默认为10，取值范围：[10,100]，用于查询列表的接口 (Optional)
 * param filters: accountStatus, 支持operator选项：eq
 (Optional)
 */
func NewDescribeAccountsRequestWithAllParams(
    regionId string,
    instanceId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeAccountsRequest {

    return &DescribeAccountsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/accounts",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAccountsRequestWithoutParam() *DescribeAccountsRequest {

    return &DescribeAccountsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/accounts",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *DescribeAccountsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例ID(Required) */
func (r *DescribeAccountsRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页;(Optional) */
func (r *DescribeAccountsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 每页显示的数据条数，默认为10，取值范围：[10,100]，用于查询列表的接口(Optional) */
func (r *DescribeAccountsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: accountStatus, 支持operator选项：eq
(Optional) */
func (r *DescribeAccountsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAccountsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAccountsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAccountsResult `json:"result"`
}

type DescribeAccountsResult struct {
    Accounts []clickhouse.Account `json:"accounts"`
    TotalCount int `json:"totalCount"`
}