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

type GravityParticleDubboJobManagerVerifyScriptFileRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 脚本文件字节数组 (Optional) */
    ScriptByte []string `json:"scriptByte"`

    /* 脚本名称 (Optional) */
    ScriptName *string `json:"scriptName"`

    /* 任务名称 (Optional) */
    JobName *string `json:"jobName"`

    /* 脚本类型 (Optional) */
    ScriptType *string `json:"scriptType"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGravityParticleDubboJobManagerVerifyScriptFileRequest(
    regionId string,
    appName string,
) *GravityParticleDubboJobManagerVerifyScriptFileRequest {

	return &GravityParticleDubboJobManagerVerifyScriptFileRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerVerifyScriptFile",
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
 * param scriptByte: 脚本文件字节数组 (Optional)
 * param scriptName: 脚本名称 (Optional)
 * param jobName: 任务名称 (Optional)
 * param scriptType: 脚本类型 (Optional)
 */
func NewGravityParticleDubboJobManagerVerifyScriptFileRequestWithAllParams(
    regionId string,
    appName string,
    scriptByte []string,
    scriptName *string,
    jobName *string,
    scriptType *string,
) *GravityParticleDubboJobManagerVerifyScriptFileRequest {

    return &GravityParticleDubboJobManagerVerifyScriptFileRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerVerifyScriptFile",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        ScriptByte: scriptByte,
        ScriptName: scriptName,
        JobName: jobName,
        ScriptType: scriptType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGravityParticleDubboJobManagerVerifyScriptFileRequestWithoutParam() *GravityParticleDubboJobManagerVerifyScriptFileRequest {

    return &GravityParticleDubboJobManagerVerifyScriptFileRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerVerifyScriptFile",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GravityParticleDubboJobManagerVerifyScriptFileRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *GravityParticleDubboJobManagerVerifyScriptFileRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param scriptByte: 脚本文件字节数组(Optional) */
func (r *GravityParticleDubboJobManagerVerifyScriptFileRequest) SetScriptByte(scriptByte []string) {
    r.ScriptByte = scriptByte
}
/* param scriptName: 脚本名称(Optional) */
func (r *GravityParticleDubboJobManagerVerifyScriptFileRequest) SetScriptName(scriptName string) {
    r.ScriptName = &scriptName
}
/* param jobName: 任务名称(Optional) */
func (r *GravityParticleDubboJobManagerVerifyScriptFileRequest) SetJobName(jobName string) {
    r.JobName = &jobName
}
/* param scriptType: 脚本类型(Optional) */
func (r *GravityParticleDubboJobManagerVerifyScriptFileRequest) SetScriptType(scriptType string) {
    r.ScriptType = &scriptType
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GravityParticleDubboJobManagerVerifyScriptFileRequest) GetRegionId() string {
    return r.RegionId
}

type GravityParticleDubboJobManagerVerifyScriptFileResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GravityParticleDubboJobManagerVerifyScriptFileResult `json:"result"`
}

type GravityParticleDubboJobManagerVerifyScriptFileResult struct {
    Success int `json:"success"`
    Code string `json:"code"`
    Msg string `json:"msg"`
    Result bool `json:"result"`
}