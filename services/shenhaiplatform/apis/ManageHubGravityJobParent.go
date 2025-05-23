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

type ManageHubGravityJobParentRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 分页-页码 (Optional) */
    PageNum *int `json:"pageNum"`

    /* 分页-每页数量 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 作业编码 (Optional) */
    JobName *string `json:"jobName"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewManageHubGravityJobParentRequest(
    regionId string,
    appName string,
) *ManageHubGravityJobParentRequest {

	return &ManageHubGravityJobParentRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/manageHubGravityJobParent",
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
 * param pageNum: 分页-页码 (Optional)
 * param pageSize: 分页-每页数量 (Optional)
 * param jobName: 作业编码 (Optional)
 */
func NewManageHubGravityJobParentRequestWithAllParams(
    regionId string,
    appName string,
    pageNum *int,
    pageSize *int,
    jobName *string,
) *ManageHubGravityJobParentRequest {

    return &ManageHubGravityJobParentRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/manageHubGravityJobParent",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        PageNum: pageNum,
        PageSize: pageSize,
        JobName: jobName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewManageHubGravityJobParentRequestWithoutParam() *ManageHubGravityJobParentRequest {

    return &ManageHubGravityJobParentRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/manageHubGravityJobParent",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ManageHubGravityJobParentRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *ManageHubGravityJobParentRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param pageNum: 分页-页码(Optional) */
func (r *ManageHubGravityJobParentRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}
/* param pageSize: 分页-每页数量(Optional) */
func (r *ManageHubGravityJobParentRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param jobName: 作业编码(Optional) */
func (r *ManageHubGravityJobParentRequest) SetJobName(jobName string) {
    r.JobName = &jobName
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ManageHubGravityJobParentRequest) GetRegionId() string {
    return r.RegionId
}

type ManageHubGravityJobParentResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ManageHubGravityJobParentResult `json:"result"`
}

type ManageHubGravityJobParentResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
    Result shenhaiplatform.PageVoJobRelationVo `json:"result"`
}