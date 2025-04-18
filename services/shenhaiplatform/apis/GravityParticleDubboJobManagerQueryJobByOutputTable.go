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

type GravityParticleDubboJobManagerQueryJobByOutputTableRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 数据库名称 (Optional) */
    DatabaseName *string `json:"databaseName"`

    /* 表名称 (Optional) */
    TableName *string `json:"tableName"`

    /* 工作空间名称 (Optional) */
    WorkspaceCode *string `json:"workspaceCode"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGravityParticleDubboJobManagerQueryJobByOutputTableRequest(
    regionId string,
    appName string,
) *GravityParticleDubboJobManagerQueryJobByOutputTableRequest {

	return &GravityParticleDubboJobManagerQueryJobByOutputTableRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerQueryJobByOutputTable",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param databaseName: 数据库名称 (Optional)
 * param tableName: 表名称 (Optional)
 * param workspaceCode: 工作空间名称 (Optional)
 */
func NewGravityParticleDubboJobManagerQueryJobByOutputTableRequestWithAllParams(
    regionId string,
    appName string,
    databaseName *string,
    tableName *string,
    workspaceCode *string,
) *GravityParticleDubboJobManagerQueryJobByOutputTableRequest {

    return &GravityParticleDubboJobManagerQueryJobByOutputTableRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerQueryJobByOutputTable",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        DatabaseName: databaseName,
        TableName: tableName,
        WorkspaceCode: workspaceCode,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGravityParticleDubboJobManagerQueryJobByOutputTableRequestWithoutParam() *GravityParticleDubboJobManagerQueryJobByOutputTableRequest {

    return &GravityParticleDubboJobManagerQueryJobByOutputTableRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerQueryJobByOutputTable",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GravityParticleDubboJobManagerQueryJobByOutputTableRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *GravityParticleDubboJobManagerQueryJobByOutputTableRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param databaseName: 数据库名称(Optional) */
func (r *GravityParticleDubboJobManagerQueryJobByOutputTableRequest) SetDatabaseName(databaseName string) {
    r.DatabaseName = &databaseName
}
/* param tableName: 表名称(Optional) */
func (r *GravityParticleDubboJobManagerQueryJobByOutputTableRequest) SetTableName(tableName string) {
    r.TableName = &tableName
}
/* param workspaceCode: 工作空间名称(Optional) */
func (r *GravityParticleDubboJobManagerQueryJobByOutputTableRequest) SetWorkspaceCode(workspaceCode string) {
    r.WorkspaceCode = &workspaceCode
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GravityParticleDubboJobManagerQueryJobByOutputTableRequest) GetRegionId() string {
    return r.RegionId
}

type GravityParticleDubboJobManagerQueryJobByOutputTableResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GravityParticleDubboJobManagerQueryJobByOutputTableResult `json:"result"`
}

type GravityParticleDubboJobManagerQueryJobByOutputTableResult struct {
    Success int `json:"success"`
    Code string `json:"code"`
    Msg string `json:"msg"`
    Result []string `json:"result"`
}