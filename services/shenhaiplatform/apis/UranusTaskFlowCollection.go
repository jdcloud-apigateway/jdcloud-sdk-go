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

type UranusTaskFlowCollectionRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 工作流Code  */
    FlowCode string `json:"flowCode"`

    /* 收藏类型不能为空 0 取消收藏 1 收藏 (Optional) */
    CollectionType *int `json:"collectionType"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param flowCode: 工作流Code (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUranusTaskFlowCollectionRequest(
    regionId string,
    appName string,
    flowCode string,
) *UranusTaskFlowCollectionRequest {

	return &UranusTaskFlowCollectionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/uranusTaskFlowCollection",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
        FlowCode: flowCode,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param flowCode: 工作流Code (Required)
 * param collectionType: 收藏类型不能为空 0 取消收藏 1 收藏 (Optional)
 */
func NewUranusTaskFlowCollectionRequestWithAllParams(
    regionId string,
    appName string,
    flowCode string,
    collectionType *int,
) *UranusTaskFlowCollectionRequest {

    return &UranusTaskFlowCollectionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusTaskFlowCollection",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        FlowCode: flowCode,
        CollectionType: collectionType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUranusTaskFlowCollectionRequestWithoutParam() *UranusTaskFlowCollectionRequest {

    return &UranusTaskFlowCollectionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusTaskFlowCollection",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UranusTaskFlowCollectionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *UranusTaskFlowCollectionRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param flowCode: 工作流Code(Required) */
func (r *UranusTaskFlowCollectionRequest) SetFlowCode(flowCode string) {
    r.FlowCode = flowCode
}
/* param collectionType: 收藏类型不能为空 0 取消收藏 1 收藏(Optional) */
func (r *UranusTaskFlowCollectionRequest) SetCollectionType(collectionType int) {
    r.CollectionType = &collectionType
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UranusTaskFlowCollectionRequest) GetRegionId() string {
    return r.RegionId
}

type UranusTaskFlowCollectionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UranusTaskFlowCollectionResult `json:"result"`
}

type UranusTaskFlowCollectionResult struct {
    Code string `json:"code"`
    ErrorTitle string `json:"errorTitle"`
    ErrorMsg string `json:"errorMsg"`
    Result interface{} `json:"result"`
    SubCode string `json:"subCode"`
    Successed bool `json:"successed"`
}