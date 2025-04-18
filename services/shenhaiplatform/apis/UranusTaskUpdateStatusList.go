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

type UranusTaskUpdateStatusListRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 任务节点集合  */
    TaskCodes []string `json:"taskCodes"`

    /* 任务节点状态  */
    Status int `json:"status"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param taskCodes: 任务节点集合 (Required)
 * param status: 任务节点状态 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUranusTaskUpdateStatusListRequest(
    regionId string,
    appName string,
    taskCodes []string,
    status int,
) *UranusTaskUpdateStatusListRequest {

	return &UranusTaskUpdateStatusListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/uranusTaskUpdateStatusList",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
        TaskCodes: taskCodes,
        Status: status,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param taskCodes: 任务节点集合 (Required)
 * param status: 任务节点状态 (Required)
 */
func NewUranusTaskUpdateStatusListRequestWithAllParams(
    regionId string,
    appName string,
    taskCodes []string,
    status int,
) *UranusTaskUpdateStatusListRequest {

    return &UranusTaskUpdateStatusListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusTaskUpdateStatusList",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        TaskCodes: taskCodes,
        Status: status,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUranusTaskUpdateStatusListRequestWithoutParam() *UranusTaskUpdateStatusListRequest {

    return &UranusTaskUpdateStatusListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusTaskUpdateStatusList",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UranusTaskUpdateStatusListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *UranusTaskUpdateStatusListRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param taskCodes: 任务节点集合(Required) */
func (r *UranusTaskUpdateStatusListRequest) SetTaskCodes(taskCodes []string) {
    r.TaskCodes = taskCodes
}
/* param status: 任务节点状态(Required) */
func (r *UranusTaskUpdateStatusListRequest) SetStatus(status int) {
    r.Status = status
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UranusTaskUpdateStatusListRequest) GetRegionId() string {
    return r.RegionId
}

type UranusTaskUpdateStatusListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UranusTaskUpdateStatusListResult `json:"result"`
}

type UranusTaskUpdateStatusListResult struct {
    Code string `json:"code"`
    ErrorTitle string `json:"errorTitle"`
    ErrorMsg string `json:"errorMsg"`
    Result interface{} `json:"result"`
    SubCode string `json:"subCode"`
    Successed bool `json:"successed"`
}