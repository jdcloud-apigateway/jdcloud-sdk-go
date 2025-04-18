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
    shenhaiplatform "github.com/jdcloud-api/jdcloud-sdk-go/services/shenhaiplatform/models"
)

type ManageHubStoragePartitionListRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 租户code  */
    CompanyCode string `json:"companyCode"`

    /* 数据库名称  */
    Database string `json:"database"`

    /* 表名称  */
    Table string `json:"table"`

    /* 分页大小  */
    PageSize int `json:"pageSize"`

    /* 页码  */
    PageNum int `json:"pageNum"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param companyCode: 租户code (Required)
 * param database: 数据库名称 (Required)
 * param table: 表名称 (Required)
 * param pageSize: 分页大小 (Required)
 * param pageNum: 页码 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewManageHubStoragePartitionListRequest(
    regionId string,
    appName string,
    companyCode string,
    database string,
    table string,
    pageSize int,
    pageNum int,
) *ManageHubStoragePartitionListRequest {

	return &ManageHubStoragePartitionListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/manageHubStoragePartitionList",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
        CompanyCode: companyCode,
        Database: database,
        Table: table,
        PageSize: pageSize,
        PageNum: pageNum,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param companyCode: 租户code (Required)
 * param database: 数据库名称 (Required)
 * param table: 表名称 (Required)
 * param pageSize: 分页大小 (Required)
 * param pageNum: 页码 (Required)
 */
func NewManageHubStoragePartitionListRequestWithAllParams(
    regionId string,
    appName string,
    companyCode string,
    database string,
    table string,
    pageSize int,
    pageNum int,
) *ManageHubStoragePartitionListRequest {

    return &ManageHubStoragePartitionListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/manageHubStoragePartitionList",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        CompanyCode: companyCode,
        Database: database,
        Table: table,
        PageSize: pageSize,
        PageNum: pageNum,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewManageHubStoragePartitionListRequestWithoutParam() *ManageHubStoragePartitionListRequest {

    return &ManageHubStoragePartitionListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/manageHubStoragePartitionList",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ManageHubStoragePartitionListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *ManageHubStoragePartitionListRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param companyCode: 租户code(Required) */
func (r *ManageHubStoragePartitionListRequest) SetCompanyCode(companyCode string) {
    r.CompanyCode = companyCode
}
/* param database: 数据库名称(Required) */
func (r *ManageHubStoragePartitionListRequest) SetDatabase(database string) {
    r.Database = database
}
/* param table: 表名称(Required) */
func (r *ManageHubStoragePartitionListRequest) SetTable(table string) {
    r.Table = table
}
/* param pageSize: 分页大小(Required) */
func (r *ManageHubStoragePartitionListRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}
/* param pageNum: 页码(Required) */
func (r *ManageHubStoragePartitionListRequest) SetPageNum(pageNum int) {
    r.PageNum = pageNum
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ManageHubStoragePartitionListRequest) GetRegionId() string {
    return r.RegionId
}

type ManageHubStoragePartitionListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ManageHubStoragePartitionListResult `json:"result"`
}

type ManageHubStoragePartitionListResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
    Result shenhaiplatform.PageVoJcwSpacePartUsedVo `json:"result"`
}