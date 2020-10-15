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

type QueryStatisticsDataGroupSumRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    EndTime *string `json:"endTime"`

    /* 需要查询的域名, 必须为用户pin下有权限的域名 (Optional) */
    Domain *string `json:"domain"`

    /* 待查询的子域名 (Optional) */
    SubDomain *string `json:"subDomain"`

    /* 需要查询的字段 (Optional) */
    Fields *string `json:"fields"`

    /*  (Optional) */
    Area *string `json:"area"`

    /*  (Optional) */
    Isp *string `json:"isp"`

    /*  (Optional) */
    Origin *string `json:"origin"`

    /* 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据 (Optional) */
    Period *string `json:"period"`

    /* 分组依据 (Optional) */
    GroupBy *string `json:"groupBy"`

    /* true 代表查询境外数据，默认false查询境内数据 (Optional) */
    Abroad *bool `json:"abroad"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryStatisticsDataGroupSumRequest(
) *QueryStatisticsDataGroupSumRequest {

	return &QueryStatisticsDataGroupSumRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/vodStatistics:groupSum",
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
 * param subDomain: 待查询的子域名 (Optional)
 * param fields: 需要查询的字段 (Optional)
 * param area:  (Optional)
 * param isp:  (Optional)
 * param origin:  (Optional)
 * param period: 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据 (Optional)
 * param groupBy: 分组依据 (Optional)
 * param abroad: true 代表查询境外数据，默认false查询境内数据 (Optional)
 */
func NewQueryStatisticsDataGroupSumRequestWithAllParams(
    startTime *string,
    endTime *string,
    domain *string,
    subDomain *string,
    fields *string,
    area *string,
    isp *string,
    origin *string,
    period *string,
    groupBy *string,
    abroad *bool,
) *QueryStatisticsDataGroupSumRequest {

    return &QueryStatisticsDataGroupSumRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/vodStatistics:groupSum",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Domain: domain,
        SubDomain: subDomain,
        Fields: fields,
        Area: area,
        Isp: isp,
        Origin: origin,
        Period: period,
        GroupBy: groupBy,
        Abroad: abroad,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryStatisticsDataGroupSumRequestWithoutParam() *QueryStatisticsDataGroupSumRequest {

    return &QueryStatisticsDataGroupSumRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/vodStatistics:groupSum",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param domain: 需要查询的域名, 必须为用户pin下有权限的域名(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetDomain(domain string) {
    r.Domain = &domain
}

/* param subDomain: 待查询的子域名(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetSubDomain(subDomain string) {
    r.SubDomain = &subDomain
}

/* param fields: 需要查询的字段(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetFields(fields string) {
    r.Fields = &fields
}

/* param area: (Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetArea(area string) {
    r.Area = &area
}

/* param isp: (Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetIsp(isp string) {
    r.Isp = &isp
}

/* param origin: (Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetOrigin(origin string) {
    r.Origin = &origin
}

/* param period: 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetPeriod(period string) {
    r.Period = &period
}

/* param groupBy: 分组依据(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetGroupBy(groupBy string) {
    r.GroupBy = &groupBy
}

/* param abroad: true 代表查询境外数据，默认false查询境内数据(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetAbroad(abroad bool) {
    r.Abroad = &abroad
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryStatisticsDataGroupSumRequest) GetRegionId() string {
    return ""
}

type QueryStatisticsDataGroupSumResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryStatisticsDataGroupSumResult `json:"result"`
}

type QueryStatisticsDataGroupSumResult struct {
    StartTime string `json:"startTime"`
    EndTime string `json:"endTime"`
    Domain string `json:"domain"`
    Statistics []cdn.StatisticsGroupSumDataItem `json:"statistics"`
}