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

type GravityParticleParticleJobMonitorJobStatusCountRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 处理类型 (Optional) */
    ProcessType *string `json:"processType"`

    /* 查询数据 (Optional) */
    Date *string `json:"date"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGravityParticleParticleJobMonitorJobStatusCountRequest(
    regionId string,
    appName string,
) *GravityParticleParticleJobMonitorJobStatusCountRequest {

	return &GravityParticleParticleJobMonitorJobStatusCountRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/gravityParticleParticleJobMonitorJobStatusCount",
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
 * param processType: 处理类型 (Optional)
 * param date: 查询数据 (Optional)
 */
func NewGravityParticleParticleJobMonitorJobStatusCountRequestWithAllParams(
    regionId string,
    appName string,
    processType *string,
    date *string,
) *GravityParticleParticleJobMonitorJobStatusCountRequest {

    return &GravityParticleParticleJobMonitorJobStatusCountRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleParticleJobMonitorJobStatusCount",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        ProcessType: processType,
        Date: date,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGravityParticleParticleJobMonitorJobStatusCountRequestWithoutParam() *GravityParticleParticleJobMonitorJobStatusCountRequest {

    return &GravityParticleParticleJobMonitorJobStatusCountRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleParticleJobMonitorJobStatusCount",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GravityParticleParticleJobMonitorJobStatusCountRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *GravityParticleParticleJobMonitorJobStatusCountRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param processType: 处理类型(Optional) */
func (r *GravityParticleParticleJobMonitorJobStatusCountRequest) SetProcessType(processType string) {
    r.ProcessType = &processType
}
/* param date: 查询数据(Optional) */
func (r *GravityParticleParticleJobMonitorJobStatusCountRequest) SetDate(date string) {
    r.Date = &date
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GravityParticleParticleJobMonitorJobStatusCountRequest) GetRegionId() string {
    return r.RegionId
}

type GravityParticleParticleJobMonitorJobStatusCountResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GravityParticleParticleJobMonitorJobStatusCountResult `json:"result"`
}

type GravityParticleParticleJobMonitorJobStatusCountResult struct {
    Success int `json:"success"`
    Result shenhaiplatform.GpmnWorkTableJobStatusDto `json:"result"`
    Code string `json:"code"`
    Msg string `json:"msg"`
    _REQ_SEQ_NO_ string `json:"_REQ_SEQ_NO_"`
}