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

type UranusExtraRunnerResultRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 执行id  */
    RunId int `json:"runId"`

    /* 结果文件名称  */
    FileName string `json:"fileName"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param runId: 执行id (Required)
 * param fileName: 结果文件名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUranusExtraRunnerResultRequest(
    regionId string,
    appName string,
    runId int,
    fileName string,
) *UranusExtraRunnerResultRequest {

	return &UranusExtraRunnerResultRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/uranusExtraRunnerResult",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
        RunId: runId,
        FileName: fileName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param runId: 执行id (Required)
 * param fileName: 结果文件名称 (Required)
 */
func NewUranusExtraRunnerResultRequestWithAllParams(
    regionId string,
    appName string,
    runId int,
    fileName string,
) *UranusExtraRunnerResultRequest {

    return &UranusExtraRunnerResultRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusExtraRunnerResult",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        RunId: runId,
        FileName: fileName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUranusExtraRunnerResultRequestWithoutParam() *UranusExtraRunnerResultRequest {

    return &UranusExtraRunnerResultRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusExtraRunnerResult",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UranusExtraRunnerResultRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *UranusExtraRunnerResultRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param runId: 执行id(Required) */
func (r *UranusExtraRunnerResultRequest) SetRunId(runId int) {
    r.RunId = runId
}
/* param fileName: 结果文件名称(Required) */
func (r *UranusExtraRunnerResultRequest) SetFileName(fileName string) {
    r.FileName = fileName
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UranusExtraRunnerResultRequest) GetRegionId() string {
    return r.RegionId
}

type UranusExtraRunnerResultResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UranusExtraRunnerResultResult `json:"result"`
}

type UranusExtraRunnerResultResult struct {
    Code string `json:"code"`
    ErrorTitle string `json:"errorTitle"`
    ErrorMsg string `json:"errorMsg"`
    Result interface{} `json:"result"`
    SubCode string `json:"subCode"`
    Successed bool `json:"successed"`
}