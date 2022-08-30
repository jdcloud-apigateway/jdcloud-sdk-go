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
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type WafQueryAttackDetailsRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    EndTime *string `json:"endTime"`

    /* 需要查询的域名, 必须为用户pin下有权限的域名 (Optional) */
    Domain *string `json:"domain"`

    /* 排序字段 (Optional) */
    SortField *string `json:"sortField"`

    /* 排序规则：desc，asc (Optional) */
    SortRule *string `json:"sortRule"`

    /* 页码，从1开始 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页大小，默认20 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewWafQueryAttackDetailsRequest(
) *WafQueryAttackDetailsRequest {

	return &WafQueryAttackDetailsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/wafAttackDetails",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param domain: 需要查询的域名, 必须为用户pin下有权限的域名 (Optional)
 * param sortField: 排序字段 (Optional)
 * param sortRule: 排序规则：desc，asc (Optional)
 * param pageNumber: 页码，从1开始 (Optional)
 * param pageSize: 页大小，默认20 (Optional)
 */
func NewWafQueryAttackDetailsRequestWithAllParams(
    startTime *string,
    endTime *string,
    domain *string,
    sortField *string,
    sortRule *string,
    pageNumber *int,
    pageSize *int,
) *WafQueryAttackDetailsRequest {

    return &WafQueryAttackDetailsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/wafAttackDetails",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Domain: domain,
        SortField: sortField,
        SortRule: sortRule,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewWafQueryAttackDetailsRequestWithoutParam() *WafQueryAttackDetailsRequest {

    return &WafQueryAttackDetailsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/wafAttackDetails",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *WafQueryAttackDetailsRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}
/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *WafQueryAttackDetailsRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}
/* param domain: 需要查询的域名, 必须为用户pin下有权限的域名(Optional) */
func (r *WafQueryAttackDetailsRequest) SetDomain(domain string) {
    r.Domain = &domain
}
/* param sortField: 排序字段(Optional) */
func (r *WafQueryAttackDetailsRequest) SetSortField(sortField string) {
    r.SortField = &sortField
}
/* param sortRule: 排序规则：desc，asc(Optional) */
func (r *WafQueryAttackDetailsRequest) SetSortRule(sortRule string) {
    r.SortRule = &sortRule
}
/* param pageNumber: 页码，从1开始(Optional) */
func (r *WafQueryAttackDetailsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 页大小，默认20(Optional) */
func (r *WafQueryAttackDetailsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r WafQueryAttackDetailsRequest) GetRegionId() string {
    return ""
}

type WafQueryAttackDetailsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result WafQueryAttackDetailsResult `json:"result"`
}

type WafQueryAttackDetailsResult struct {
    Total string `json:"total"`
    AttackDetails []cdn.AttackDetail `json:"attackDetails"`
}