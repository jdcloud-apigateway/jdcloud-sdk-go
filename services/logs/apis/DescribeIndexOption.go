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
    logs "github.com/jdcloud-api/jdcloud-sdk-go/services/logs/models"
)

type DescribeIndexOptionRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 日志主题 UID  */
    LogtopicUID string `json:"logtopicUID"`
}

/*
 * param regionId: 地域 Id (Required)
 * param logtopicUID: 日志主题 UID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeIndexOptionRequest(
    regionId string,
    logtopicUID string,
) *DescribeIndexOptionRequest {

	return &DescribeIndexOptionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/logtopics/{logtopicUID}/indexoptions",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LogtopicUID: logtopicUID,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param logtopicUID: 日志主题 UID (Required)
 */
func NewDescribeIndexOptionRequestWithAllParams(
    regionId string,
    logtopicUID string,
) *DescribeIndexOptionRequest {

    return &DescribeIndexOptionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logtopics/{logtopicUID}/indexoptions",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LogtopicUID: logtopicUID,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeIndexOptionRequestWithoutParam() *DescribeIndexOptionRequest {

    return &DescribeIndexOptionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logtopics/{logtopicUID}/indexoptions",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeIndexOptionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param logtopicUID: 日志主题 UID(Required) */
func (r *DescribeIndexOptionRequest) SetLogtopicUID(logtopicUID string) {
    r.LogtopicUID = logtopicUID
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeIndexOptionRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeIndexOptionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeIndexOptionResult `json:"result"`
}

type DescribeIndexOptionResult struct {
    CaseSensitive bool `json:"caseSensitive"`
    Chn bool `json:"chn"`
    FieldIndexOptions []logs.FieldIndexOption `json:"fieldIndexOptions"`
    FulltextIndex bool `json:"fulltextIndex"`
    LogReduce bool `json:"logReduce"`
    MaxTextLen int64 `json:"maxTextLen"`
    Token []string `json:"token"`
    Uid string `json:"uid"`
}