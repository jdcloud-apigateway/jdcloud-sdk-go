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
    waf "github.com/jdcloud-api/jdcloud-sdk-go/services/waf/models"
)

type GetEsLogDetailRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 实例id，代表要查询的WAF实例，为空时表示当前用户下的所有实例 (Optional) */
    WafInstanceId *string `json:"wafInstanceId"`

    /* 域名，为空时表示当前实例下的所有域名 (Optional) */
    Domain *string `json:"domain"`

    /* 来源ip，检索字段 (Optional) */
    Remote_addr *string `json:"remote_addr"`

    /* URI，检索字段 (Optional) */
    Document_uri *string `json:"document_uri"`

    /* url，检索字段 (Optional) */
    Url *string `json:"url"`

    /* 来源地域，检索字段 (Optional) */
    Anti_geo *string `json:"anti_geo"`

    /* 请求方法，检索字段 (Optional) */
    Request_method *string `json:"request_method"`

    /* 动作，检索字段，支持类型：""(为空时，默认查询全部动作类型)，"-"(放行)，"notice"(观察)，"forbidden"(拦截)，"redirect"(浏览器跳转)，"verify"(人机交互) (Optional) */
    Action *string `json:"action"`

    /* 状态码，检索字段 (Optional) */
    Status *string `json:"status"`

    /* 日志类型，检索字段，支持类型：""(为空时，默认查询全部日志类型)，"access"(访问日志)，"waf"(wafSDK)，"acl"(自定义规则)，"skip"(白名单)，"deny"(黑名单)，"cc"(CC攻击)，"webcache"(网页防篡改)，"css"(跨站脚本攻击)，"sqli"(SQL注入攻击)，""fileinc"(文件读取/包含攻击)，"cmding"(命令/代码执行攻击)，"sdd"(敏感文件探测)，"malscan"(恶意扫描攻击)，"bckack"(恶意/后门文件攻击)，"xmli"(XML注入攻击)，"dirt"(目录遍历攻击) (Optional) */
    LogType []string `json:"logType"`

    /* 日志Id，检索字段 (Optional) */
    LogId *string `json:"logId"`

    /* 开始时间戳，单位秒，时间间隔要求大于5分钟，小于30天。  */
    Start int `json:"start"`

    /* 结束时间戳，单位秒，时间间隔要求大于5分钟，小于30天。  */
    End int `json:"end"`

    /* 每页显示的个数，默认是10。  */
    PageSize int `json:"pageSize"`

    /* 页数，默认是1。  */
    PageIndex int `json:"pageIndex"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param start: 开始时间戳，单位秒，时间间隔要求大于5分钟，小于30天。 (Required)
 * param end: 结束时间戳，单位秒，时间间隔要求大于5分钟，小于30天。 (Required)
 * param pageSize: 每页显示的个数，默认是10。 (Required)
 * param pageIndex: 页数，默认是1。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetEsLogDetailRequest(
    regionId string,
    start int,
    end int,
    pageSize int,
    pageIndex int,
) *GetEsLogDetailRequest {

	return &GetEsLogDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/chart:getEsLog",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Start: start,
        End: end,
        PageSize: pageSize,
        PageIndex: pageIndex,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param wafInstanceId: 实例id，代表要查询的WAF实例，为空时表示当前用户下的所有实例 (Optional)
 * param domain: 域名，为空时表示当前实例下的所有域名 (Optional)
 * param remote_addr: 来源ip，检索字段 (Optional)
 * param document_uri: URI，检索字段 (Optional)
 * param url: url，检索字段 (Optional)
 * param anti_geo: 来源地域，检索字段 (Optional)
 * param request_method: 请求方法，检索字段 (Optional)
 * param action: 动作，检索字段，支持类型：""(为空时，默认查询全部动作类型)，"-"(放行)，"notice"(观察)，"forbidden"(拦截)，"redirect"(浏览器跳转)，"verify"(人机交互) (Optional)
 * param status: 状态码，检索字段 (Optional)
 * param logType: 日志类型，检索字段，支持类型：""(为空时，默认查询全部日志类型)，"access"(访问日志)，"waf"(wafSDK)，"acl"(自定义规则)，"skip"(白名单)，"deny"(黑名单)，"cc"(CC攻击)，"webcache"(网页防篡改)，"css"(跨站脚本攻击)，"sqli"(SQL注入攻击)，""fileinc"(文件读取/包含攻击)，"cmding"(命令/代码执行攻击)，"sdd"(敏感文件探测)，"malscan"(恶意扫描攻击)，"bckack"(恶意/后门文件攻击)，"xmli"(XML注入攻击)，"dirt"(目录遍历攻击) (Optional)
 * param logId: 日志Id，检索字段 (Optional)
 * param start: 开始时间戳，单位秒，时间间隔要求大于5分钟，小于30天。 (Required)
 * param end: 结束时间戳，单位秒，时间间隔要求大于5分钟，小于30天。 (Required)
 * param pageSize: 每页显示的个数，默认是10。 (Required)
 * param pageIndex: 页数，默认是1。 (Required)
 */
func NewGetEsLogDetailRequestWithAllParams(
    regionId string,
    wafInstanceId *string,
    domain *string,
    remote_addr *string,
    document_uri *string,
    url *string,
    anti_geo *string,
    request_method *string,
    action *string,
    status *string,
    logType []string,
    logId *string,
    start int,
    end int,
    pageSize int,
    pageIndex int,
) *GetEsLogDetailRequest {

    return &GetEsLogDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/chart:getEsLog",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        WafInstanceId: wafInstanceId,
        Domain: domain,
        Remote_addr: remote_addr,
        Document_uri: document_uri,
        Url: url,
        Anti_geo: anti_geo,
        Request_method: request_method,
        Action: action,
        Status: status,
        LogType: logType,
        LogId: logId,
        Start: start,
        End: end,
        PageSize: pageSize,
        PageIndex: pageIndex,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetEsLogDetailRequestWithoutParam() *GetEsLogDetailRequest {

    return &GetEsLogDetailRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/chart:getEsLog",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *GetEsLogDetailRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param wafInstanceId: 实例id，代表要查询的WAF实例，为空时表示当前用户下的所有实例(Optional) */
func (r *GetEsLogDetailRequest) SetWafInstanceId(wafInstanceId string) {
    r.WafInstanceId = &wafInstanceId
}

/* param domain: 域名，为空时表示当前实例下的所有域名(Optional) */
func (r *GetEsLogDetailRequest) SetDomain(domain string) {
    r.Domain = &domain
}

/* param remote_addr: 来源ip，检索字段(Optional) */
func (r *GetEsLogDetailRequest) SetRemote_addr(remote_addr string) {
    r.Remote_addr = &remote_addr
}

/* param document_uri: URI，检索字段(Optional) */
func (r *GetEsLogDetailRequest) SetDocument_uri(document_uri string) {
    r.Document_uri = &document_uri
}

/* param url: url，检索字段(Optional) */
func (r *GetEsLogDetailRequest) SetUrl(url string) {
    r.Url = &url
}

/* param anti_geo: 来源地域，检索字段(Optional) */
func (r *GetEsLogDetailRequest) SetAnti_geo(anti_geo string) {
    r.Anti_geo = &anti_geo
}

/* param request_method: 请求方法，检索字段(Optional) */
func (r *GetEsLogDetailRequest) SetRequest_method(request_method string) {
    r.Request_method = &request_method
}

/* param action: 动作，检索字段，支持类型：""(为空时，默认查询全部动作类型)，"-"(放行)，"notice"(观察)，"forbidden"(拦截)，"redirect"(浏览器跳转)，"verify"(人机交互)(Optional) */
func (r *GetEsLogDetailRequest) SetAction(action string) {
    r.Action = &action
}

/* param status: 状态码，检索字段(Optional) */
func (r *GetEsLogDetailRequest) SetStatus(status string) {
    r.Status = &status
}

/* param logType: 日志类型，检索字段，支持类型：""(为空时，默认查询全部日志类型)，"access"(访问日志)，"waf"(wafSDK)，"acl"(自定义规则)，"skip"(白名单)，"deny"(黑名单)，"cc"(CC攻击)，"webcache"(网页防篡改)，"css"(跨站脚本攻击)，"sqli"(SQL注入攻击)，""fileinc"(文件读取/包含攻击)，"cmding"(命令/代码执行攻击)，"sdd"(敏感文件探测)，"malscan"(恶意扫描攻击)，"bckack"(恶意/后门文件攻击)，"xmli"(XML注入攻击)，"dirt"(目录遍历攻击)(Optional) */
func (r *GetEsLogDetailRequest) SetLogType(logType []string) {
    r.LogType = logType
}

/* param logId: 日志Id，检索字段(Optional) */
func (r *GetEsLogDetailRequest) SetLogId(logId string) {
    r.LogId = &logId
}

/* param start: 开始时间戳，单位秒，时间间隔要求大于5分钟，小于30天。(Required) */
func (r *GetEsLogDetailRequest) SetStart(start int) {
    r.Start = start
}

/* param end: 结束时间戳，单位秒，时间间隔要求大于5分钟，小于30天。(Required) */
func (r *GetEsLogDetailRequest) SetEnd(end int) {
    r.End = end
}

/* param pageSize: 每页显示的个数，默认是10。(Required) */
func (r *GetEsLogDetailRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}

/* param pageIndex: 页数，默认是1。(Required) */
func (r *GetEsLogDetailRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = pageIndex
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetEsLogDetailRequest) GetRegionId() string {
    return r.RegionId
}

type GetEsLogDetailResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetEsLogDetailResult `json:"result"`
}

type GetEsLogDetailResult struct {
    PageSize int `json:"pageSize"`
    PageIndex int `json:"pageIndex"`
    Total int `json:"total"`
    Events []waf.EsLogEvent `json:"events"`
}