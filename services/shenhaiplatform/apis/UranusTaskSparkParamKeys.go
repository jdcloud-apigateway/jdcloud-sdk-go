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

type UranusTaskSparkParamKeysRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUranusTaskSparkParamKeysRequest(
    regionId string,
    appName string,
) *UranusTaskSparkParamKeysRequest {

	return &UranusTaskSparkParamKeysRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/uranusTaskSparkParamKeys",
			Method:  "GET",
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
 */
func NewUranusTaskSparkParamKeysRequestWithAllParams(
    regionId string,
    appName string,
) *UranusTaskSparkParamKeysRequest {

    return &UranusTaskSparkParamKeysRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusTaskSparkParamKeys",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUranusTaskSparkParamKeysRequestWithoutParam() *UranusTaskSparkParamKeysRequest {

    return &UranusTaskSparkParamKeysRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusTaskSparkParamKeys",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UranusTaskSparkParamKeysRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *UranusTaskSparkParamKeysRequest) SetAppName(appName string) {
    r.AppName = appName
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UranusTaskSparkParamKeysRequest) GetRegionId() string {
    return r.RegionId
}

type UranusTaskSparkParamKeysResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UranusTaskSparkParamKeysResult `json:"result"`
}

type UranusTaskSparkParamKeysResult struct {
    Code string `json:"code"`
    ErrorTitle string `json:"errorTitle"`
    ErrorMsg string `json:"errorMsg"`
    Result []interface{} `json:"result"`
    SubCode string `json:"subCode"`
    Successed bool `json:"successed"`
}