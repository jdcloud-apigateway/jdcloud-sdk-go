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

type DescribeLoadBalancersBySecurityGroupRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* securityGroup ID  */
    SecurityGroupId string `json:"securityGroupId"`

    /* 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: Region ID (Required)
 * param securityGroupId: securityGroup ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLoadBalancersBySecurityGroupRequest(
    regionId string,
    securityGroupId string,
) *DescribeLoadBalancersBySecurityGroupRequest {

	return &DescribeLoadBalancersBySecurityGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/loadBalancers/{securityGroupId}:describeLoadBalancersBySecurityGroup",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SecurityGroupId: securityGroupId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param securityGroupId: securityGroup ID (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 */
func NewDescribeLoadBalancersBySecurityGroupRequestWithAllParams(
    regionId string,
    securityGroupId string,
    pageNumber *int,
    pageSize *int,
) *DescribeLoadBalancersBySecurityGroupRequest {

    return &DescribeLoadBalancersBySecurityGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loadBalancers/{securityGroupId}:describeLoadBalancersBySecurityGroup",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SecurityGroupId: securityGroupId,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLoadBalancersBySecurityGroupRequestWithoutParam() *DescribeLoadBalancersBySecurityGroupRequest {

    return &DescribeLoadBalancersBySecurityGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loadBalancers/{securityGroupId}:describeLoadBalancersBySecurityGroup",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeLoadBalancersBySecurityGroupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param securityGroupId: securityGroup ID(Required) */
func (r *DescribeLoadBalancersBySecurityGroupRequest) SetSecurityGroupId(securityGroupId string) {
    r.SecurityGroupId = securityGroupId
}
/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页(Optional) */
func (r *DescribeLoadBalancersBySecurityGroupRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *DescribeLoadBalancersBySecurityGroupRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLoadBalancersBySecurityGroupRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeLoadBalancersBySecurityGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLoadBalancersBySecurityGroupResult `json:"result"`
}

type DescribeLoadBalancersBySecurityGroupResult struct {
    TotalCount int `json:"totalCount"`
    LoadBalancers []lb.LoadBalancer `json:"loadBalancers"`
}