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
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeSecurityPoliciesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* securityPolicyIds - 安全策略ID，支持多个
securityPolicyNames - 安全策略名称，支持多个; 支持operator为like的模糊搜索，此时name只能传单个
securityPolicyType - 安全策略类型(SYSTEM/CUSTOM/ALL)，支持单个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSecurityPoliciesRequest(
    regionId string,
) *DescribeSecurityPoliciesRequest {

	return &DescribeSecurityPoliciesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/securityPolicies/",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 * param filters: securityPolicyIds - 安全策略ID，支持多个
securityPolicyNames - 安全策略名称，支持多个; 支持operator为like的模糊搜索，此时name只能传单个
securityPolicyType - 安全策略类型(SYSTEM/CUSTOM/ALL)，支持单个
 (Optional)
 */
func NewDescribeSecurityPoliciesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeSecurityPoliciesRequest {

    return &DescribeSecurityPoliciesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/securityPolicies/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSecurityPoliciesRequestWithoutParam() *DescribeSecurityPoliciesRequest {

    return &DescribeSecurityPoliciesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/securityPolicies/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeSecurityPoliciesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页(Optional) */
func (r *DescribeSecurityPoliciesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *DescribeSecurityPoliciesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param filters: securityPolicyIds - 安全策略ID，支持多个
securityPolicyNames - 安全策略名称，支持多个; 支持operator为like的模糊搜索，此时name只能传单个
securityPolicyType - 安全策略类型(SYSTEM/CUSTOM/ALL)，支持单个
(Optional) */
func (r *DescribeSecurityPoliciesRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSecurityPoliciesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSecurityPoliciesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSecurityPoliciesResult `json:"result"`
}

type DescribeSecurityPoliciesResult struct {
    SecurityPolicies []lb.SecurityPolicy `json:"securityPolicies"`
    TotalCount int `json:"totalCount"`
}