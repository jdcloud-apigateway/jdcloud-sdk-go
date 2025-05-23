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

type GravityParticleOpenRerunRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 数据日期, 对于非小时/分钟周期任务，格式为 yyyy-MM-dd；对于小时/分钟周期任务，格式为 yyyy-MM-dd-HH-mm  */
    TxDate string `json:"txDate"`

    /* 作业所属工作空间编码，形如: ws_123456789012345678  */
    WorkspaceCode string `json:"workspaceCode"`

    /* 作业ID列表，最多允许传递20个作业  */
    JobIdList []string `json:"jobIdList"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param txDate: 数据日期, 对于非小时/分钟周期任务，格式为 yyyy-MM-dd；对于小时/分钟周期任务，格式为 yyyy-MM-dd-HH-mm (Required)
 * param workspaceCode: 作业所属工作空间编码，形如: ws_123456789012345678 (Required)
 * param jobIdList: 作业ID列表，最多允许传递20个作业 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGravityParticleOpenRerunRequest(
    regionId string,
    appName string,
    txDate string,
    workspaceCode string,
    jobIdList []string,
) *GravityParticleOpenRerunRequest {

	return &GravityParticleOpenRerunRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/gravityParticleOpenRerun",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
        TxDate: txDate,
        WorkspaceCode: workspaceCode,
        JobIdList: jobIdList,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param txDate: 数据日期, 对于非小时/分钟周期任务，格式为 yyyy-MM-dd；对于小时/分钟周期任务，格式为 yyyy-MM-dd-HH-mm (Required)
 * param workspaceCode: 作业所属工作空间编码，形如: ws_123456789012345678 (Required)
 * param jobIdList: 作业ID列表，最多允许传递20个作业 (Required)
 */
func NewGravityParticleOpenRerunRequestWithAllParams(
    regionId string,
    appName string,
    txDate string,
    workspaceCode string,
    jobIdList []string,
) *GravityParticleOpenRerunRequest {

    return &GravityParticleOpenRerunRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleOpenRerun",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        TxDate: txDate,
        WorkspaceCode: workspaceCode,
        JobIdList: jobIdList,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGravityParticleOpenRerunRequestWithoutParam() *GravityParticleOpenRerunRequest {

    return &GravityParticleOpenRerunRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleOpenRerun",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GravityParticleOpenRerunRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *GravityParticleOpenRerunRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param txDate: 数据日期, 对于非小时/分钟周期任务，格式为 yyyy-MM-dd；对于小时/分钟周期任务，格式为 yyyy-MM-dd-HH-mm(Required) */
func (r *GravityParticleOpenRerunRequest) SetTxDate(txDate string) {
    r.TxDate = txDate
}
/* param workspaceCode: 作业所属工作空间编码，形如: ws_123456789012345678(Required) */
func (r *GravityParticleOpenRerunRequest) SetWorkspaceCode(workspaceCode string) {
    r.WorkspaceCode = workspaceCode
}
/* param jobIdList: 作业ID列表，最多允许传递20个作业(Required) */
func (r *GravityParticleOpenRerunRequest) SetJobIdList(jobIdList []string) {
    r.JobIdList = jobIdList
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GravityParticleOpenRerunRequest) GetRegionId() string {
    return r.RegionId
}

type GravityParticleOpenRerunResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GravityParticleOpenRerunResult `json:"result"`
}

type GravityParticleOpenRerunResult struct {
    Success int `json:"success"`
    Result shenhaiplatform.JobRerunResult `json:"result"`
    Code string `json:"code"`
    Msg string `json:"msg"`
}