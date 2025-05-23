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
)

type JtlasPageSearchTablesRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 是否是过滤请求(必填)，如果是则返回aggregate结果  */
    FilterQuery bool `json:"filterQuery"`

    /* 搜索模式(必填)，准确和模糊检索  */
    SearchMode string `json:"searchMode"`

    /*  (Optional) */
    SearchKeyword *string `json:"searchKeyword"`

    /* 工作空间(必填)  */
    WorkspaceCode string `json:"workspaceCode"`

    /* dev或prod空字符串代表所有环境  */
    Env string `json:"env"`

    /* 检索范围(必填)  */
    SearchRange string `json:"searchRange"`

    /* 检索指定用户下的表的时候，需要传递(可选) (Optional) */
    PersonInCharge *string `json:"personInCharge"`

    /* 返回结果的排列方式(必填)，按照点击/创建时间升降/相关度  */
    Sort string `json:"sort"`

    /* 分页页码(必填) (Optional) */
    PageNum *int `json:"pageNum"`

    /* 分页大小(必填) (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param filterQuery: 是否是过滤请求(必填)，如果是则返回aggregate结果 (Required)
 * param searchMode: 搜索模式(必填)，准确和模糊检索 (Required)
 * param workspaceCode: 工作空间(必填) (Required)
 * param env: dev或prod空字符串代表所有环境 (Required)
 * param searchRange: 检索范围(必填) (Required)
 * param sort: 返回结果的排列方式(必填)，按照点击/创建时间升降/相关度 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewJtlasPageSearchTablesRequest(
    regionId string,
    appName string,
    filterQuery bool,
    searchMode string,
    workspaceCode string,
    env string,
    searchRange string,
    sort string,
) *JtlasPageSearchTablesRequest {

	return &JtlasPageSearchTablesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/jtlasPageSearchTables",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
        FilterQuery: filterQuery,
        SearchMode: searchMode,
        WorkspaceCode: workspaceCode,
        Env: env,
        SearchRange: searchRange,
        Sort: sort,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param filterQuery: 是否是过滤请求(必填)，如果是则返回aggregate结果 (Required)
 * param searchMode: 搜索模式(必填)，准确和模糊检索 (Required)
 * param searchKeyword:  (Optional)
 * param workspaceCode: 工作空间(必填) (Required)
 * param env: dev或prod空字符串代表所有环境 (Required)
 * param searchRange: 检索范围(必填) (Required)
 * param personInCharge: 检索指定用户下的表的时候，需要传递(可选) (Optional)
 * param sort: 返回结果的排列方式(必填)，按照点击/创建时间升降/相关度 (Required)
 * param pageNum: 分页页码(必填) (Optional)
 * param pageSize: 分页大小(必填) (Optional)
 */
func NewJtlasPageSearchTablesRequestWithAllParams(
    regionId string,
    appName string,
    filterQuery bool,
    searchMode string,
    searchKeyword *string,
    workspaceCode string,
    env string,
    searchRange string,
    personInCharge *string,
    sort string,
    pageNum *int,
    pageSize *int,
) *JtlasPageSearchTablesRequest {

    return &JtlasPageSearchTablesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/jtlasPageSearchTables",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        FilterQuery: filterQuery,
        SearchMode: searchMode,
        SearchKeyword: searchKeyword,
        WorkspaceCode: workspaceCode,
        Env: env,
        SearchRange: searchRange,
        PersonInCharge: personInCharge,
        Sort: sort,
        PageNum: pageNum,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewJtlasPageSearchTablesRequestWithoutParam() *JtlasPageSearchTablesRequest {

    return &JtlasPageSearchTablesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/jtlasPageSearchTables",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *JtlasPageSearchTablesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *JtlasPageSearchTablesRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param filterQuery: 是否是过滤请求(必填)，如果是则返回aggregate结果(Required) */
func (r *JtlasPageSearchTablesRequest) SetFilterQuery(filterQuery bool) {
    r.FilterQuery = filterQuery
}
/* param searchMode: 搜索模式(必填)，准确和模糊检索(Required) */
func (r *JtlasPageSearchTablesRequest) SetSearchMode(searchMode string) {
    r.SearchMode = searchMode
}
/* param searchKeyword: (Optional) */
func (r *JtlasPageSearchTablesRequest) SetSearchKeyword(searchKeyword string) {
    r.SearchKeyword = &searchKeyword
}
/* param workspaceCode: 工作空间(必填)(Required) */
func (r *JtlasPageSearchTablesRequest) SetWorkspaceCode(workspaceCode string) {
    r.WorkspaceCode = workspaceCode
}
/* param env: dev或prod空字符串代表所有环境(Required) */
func (r *JtlasPageSearchTablesRequest) SetEnv(env string) {
    r.Env = env
}
/* param searchRange: 检索范围(必填)(Required) */
func (r *JtlasPageSearchTablesRequest) SetSearchRange(searchRange string) {
    r.SearchRange = searchRange
}
/* param personInCharge: 检索指定用户下的表的时候，需要传递(可选)(Optional) */
func (r *JtlasPageSearchTablesRequest) SetPersonInCharge(personInCharge string) {
    r.PersonInCharge = &personInCharge
}
/* param sort: 返回结果的排列方式(必填)，按照点击/创建时间升降/相关度(Required) */
func (r *JtlasPageSearchTablesRequest) SetSort(sort string) {
    r.Sort = sort
}
/* param pageNum: 分页页码(必填)(Optional) */
func (r *JtlasPageSearchTablesRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}
/* param pageSize: 分页大小(必填)(Optional) */
func (r *JtlasPageSearchTablesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r JtlasPageSearchTablesRequest) GetRegionId() string {
    return r.RegionId
}

type JtlasPageSearchTablesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result JtlasPageSearchTablesResult `json:"result"`
}

type JtlasPageSearchTablesResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
    Result interface{} `json:"result"`
}