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

type UranusScriptSaveContentRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 文件Code  */
    FileCode string `json:"fileCode"`

    /* 文件内容  */
    Content string `json:"content"`

    /* 文件类型 后缀名  */
    FileType string `json:"fileType"`

    /* 创建时间 (Optional) */
    CreatedTime *string `json:"createdTime"`

    /* 锁状态  */
    LockStatus int `json:"lockStatus"`

    /* 获得锁的用户  */
    LockUser string `json:"lockUser"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param fileCode: 文件Code (Required)
 * param content: 文件内容 (Required)
 * param fileType: 文件类型 后缀名 (Required)
 * param lockStatus: 锁状态 (Required)
 * param lockUser: 获得锁的用户 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUranusScriptSaveContentRequest(
    regionId string,
    appName string,
    fileCode string,
    content string,
    fileType string,
    lockStatus int,
    lockUser string,
) *UranusScriptSaveContentRequest {

	return &UranusScriptSaveContentRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/uranusScriptSaveContent",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
        FileCode: fileCode,
        Content: content,
        FileType: fileType,
        LockStatus: lockStatus,
        LockUser: lockUser,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param fileCode: 文件Code (Required)
 * param content: 文件内容 (Required)
 * param fileType: 文件类型 后缀名 (Required)
 * param createdTime: 创建时间 (Optional)
 * param lockStatus: 锁状态 (Required)
 * param lockUser: 获得锁的用户 (Required)
 */
func NewUranusScriptSaveContentRequestWithAllParams(
    regionId string,
    appName string,
    fileCode string,
    content string,
    fileType string,
    createdTime *string,
    lockStatus int,
    lockUser string,
) *UranusScriptSaveContentRequest {

    return &UranusScriptSaveContentRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusScriptSaveContent",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        FileCode: fileCode,
        Content: content,
        FileType: fileType,
        CreatedTime: createdTime,
        LockStatus: lockStatus,
        LockUser: lockUser,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUranusScriptSaveContentRequestWithoutParam() *UranusScriptSaveContentRequest {

    return &UranusScriptSaveContentRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusScriptSaveContent",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UranusScriptSaveContentRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *UranusScriptSaveContentRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param fileCode: 文件Code(Required) */
func (r *UranusScriptSaveContentRequest) SetFileCode(fileCode string) {
    r.FileCode = fileCode
}
/* param content: 文件内容(Required) */
func (r *UranusScriptSaveContentRequest) SetContent(content string) {
    r.Content = content
}
/* param fileType: 文件类型 后缀名(Required) */
func (r *UranusScriptSaveContentRequest) SetFileType(fileType string) {
    r.FileType = fileType
}
/* param createdTime: 创建时间(Optional) */
func (r *UranusScriptSaveContentRequest) SetCreatedTime(createdTime string) {
    r.CreatedTime = &createdTime
}
/* param lockStatus: 锁状态(Required) */
func (r *UranusScriptSaveContentRequest) SetLockStatus(lockStatus int) {
    r.LockStatus = lockStatus
}
/* param lockUser: 获得锁的用户(Required) */
func (r *UranusScriptSaveContentRequest) SetLockUser(lockUser string) {
    r.LockUser = lockUser
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UranusScriptSaveContentRequest) GetRegionId() string {
    return r.RegionId
}

type UranusScriptSaveContentResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UranusScriptSaveContentResult `json:"result"`
}

type UranusScriptSaveContentResult struct {
    Code string `json:"code"`
    ErrorTitle string `json:"errorTitle"`
    ErrorMsg string `json:"errorMsg"`
    Result interface{} `json:"result"`
    SubCode string `json:"subCode"`
    Successed bool `json:"successed"`
}