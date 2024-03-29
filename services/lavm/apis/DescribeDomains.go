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
    lavm "github.com/jdcloud-api/jdcloud-sdk-go/services/lavm/models"
)

type DescribeDomainsRequest struct {

    core.JDCloudRequest

    /* regionId
  */
    RegionId string `json:"regionId"`

    /* 域名名称。支持模糊搜索，最多20个，例如: `[\"name-1\", \"name-2\"]`, json array 字串。
 (Optional) */
    DomainNames *string `json:"domainNames"`

    /* 页码；默认为1。 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；<br>默认为20；取值范围[10, 100]。 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: regionId
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDomainsRequest(
    regionId string,
) *DescribeDomainsRequest {

	return &DescribeDomainsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instanceDomains",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: regionId
 (Required)
 * param domainNames: 域名名称。支持模糊搜索，最多20个，例如: `[\"name-1\", \"name-2\"]`, json array 字串。
 (Optional)
 * param pageNumber: 页码；默认为1。 (Optional)
 * param pageSize: 分页大小；<br>默认为20；取值范围[10, 100]。 (Optional)
 */
func NewDescribeDomainsRequestWithAllParams(
    regionId string,
    domainNames *string,
    pageNumber *int,
    pageSize *int,
) *DescribeDomainsRequest {

    return &DescribeDomainsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instanceDomains",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DomainNames: domainNames,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDomainsRequestWithoutParam() *DescribeDomainsRequest {

    return &DescribeDomainsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instanceDomains",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: regionId
(Required) */
func (r *DescribeDomainsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param domainNames: 域名名称。支持模糊搜索，最多20个，例如: `[\"name-1\", \"name-2\"]`, json array 字串。
(Optional) */
func (r *DescribeDomainsRequest) SetDomainNames(domainNames string) {
    r.DomainNames = &domainNames
}
/* param pageNumber: 页码；默认为1。(Optional) */
func (r *DescribeDomainsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 分页大小；<br>默认为20；取值范围[10, 100]。(Optional) */
func (r *DescribeDomainsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDomainsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDomainsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDomainsResult `json:"result"`
}

type DescribeDomainsResult struct {
    Domains []lavm.InstanceDomain `json:"domains"`
    TotalCount int `json:"totalCount"`
}