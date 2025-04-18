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

type GravityParticleJobscheSetJobEnableFlagRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 作业名称 (Optional) */
    JobName *string `json:"jobName"`

    /* 是否可用，1 已上线，2已下线 (Optional) */
    EnableFlag *string `json:"enableFlag"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGravityParticleJobscheSetJobEnableFlagRequest(
    regionId string,
    appName string,
) *GravityParticleJobscheSetJobEnableFlagRequest {

	return &GravityParticleJobscheSetJobEnableFlagRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/gravityParticleJobscheSetJobEnableFlag",
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
 * param jobName: 作业名称 (Optional)
 * param enableFlag: 是否可用，1 已上线，2已下线 (Optional)
 */
func NewGravityParticleJobscheSetJobEnableFlagRequestWithAllParams(
    regionId string,
    appName string,
    jobName *string,
    enableFlag *string,
) *GravityParticleJobscheSetJobEnableFlagRequest {

    return &GravityParticleJobscheSetJobEnableFlagRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleJobscheSetJobEnableFlag",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        JobName: jobName,
        EnableFlag: enableFlag,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGravityParticleJobscheSetJobEnableFlagRequestWithoutParam() *GravityParticleJobscheSetJobEnableFlagRequest {

    return &GravityParticleJobscheSetJobEnableFlagRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleJobscheSetJobEnableFlag",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GravityParticleJobscheSetJobEnableFlagRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *GravityParticleJobscheSetJobEnableFlagRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param jobName: 作业名称(Optional) */
func (r *GravityParticleJobscheSetJobEnableFlagRequest) SetJobName(jobName string) {
    r.JobName = &jobName
}
/* param enableFlag: 是否可用，1 已上线，2已下线(Optional) */
func (r *GravityParticleJobscheSetJobEnableFlagRequest) SetEnableFlag(enableFlag string) {
    r.EnableFlag = &enableFlag
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GravityParticleJobscheSetJobEnableFlagRequest) GetRegionId() string {
    return r.RegionId
}

type GravityParticleJobscheSetJobEnableFlagResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GravityParticleJobscheSetJobEnableFlagResult `json:"result"`
}

type GravityParticleJobscheSetJobEnableFlagResult struct {
    Success int `json:"success"`
    Result string `json:"result"`
    Code string `json:"code"`
    Msg string `json:"msg"`
    _REQ_SEQ_NO_ string `json:"_REQ_SEQ_NO_"`
}