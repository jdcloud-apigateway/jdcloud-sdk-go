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

type CreateLogtopicRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 日志集 UID  */
    LogsetUID string `json:"logsetUID"`

    /* 日志主题名称  */
    Name string `json:"name"`

    /* 日志集描述 (Optional) */
    Description *string `json:"description"`

    /* 保序 (Optional) */
    InOrder *bool `json:"inOrder"`

    /* 保存周期，只能是 7， 15， 30 (Optional) */
    LifeCycle *int `json:"lifeCycle"`

    /* 标签列表 (Optional) */
    Tags []logs.Tag `json:"tags"`
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUID: 日志集 UID (Required)
 * param name: 日志主题名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateLogtopicRequest(
    regionId string,
    logsetUID string,
    name string,
) *CreateLogtopicRequest {

	return &CreateLogtopicRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LogsetUID: logsetUID,
        Name: name,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUID: 日志集 UID (Required)
 * param name: 日志主题名称 (Required)
 * param description: 日志集描述 (Optional)
 * param inOrder: 保序 (Optional)
 * param lifeCycle: 保存周期，只能是 7， 15， 30 (Optional)
 * param tags: 标签列表 (Optional)
 */
func NewCreateLogtopicRequestWithAllParams(
    regionId string,
    logsetUID string,
    name string,
    description *string,
    inOrder *bool,
    lifeCycle *int,
    tags []logs.Tag,
) *CreateLogtopicRequest {

    return &CreateLogtopicRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LogsetUID: logsetUID,
        Name: name,
        Description: description,
        InOrder: inOrder,
        LifeCycle: lifeCycle,
        Tags: tags,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateLogtopicRequestWithoutParam() *CreateLogtopicRequest {

    return &CreateLogtopicRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *CreateLogtopicRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param logsetUID: 日志集 UID(Required) */
func (r *CreateLogtopicRequest) SetLogsetUID(logsetUID string) {
    r.LogsetUID = logsetUID
}
/* param name: 日志主题名称(Required) */
func (r *CreateLogtopicRequest) SetName(name string) {
    r.Name = name
}
/* param description: 日志集描述(Optional) */
func (r *CreateLogtopicRequest) SetDescription(description string) {
    r.Description = &description
}
/* param inOrder: 保序(Optional) */
func (r *CreateLogtopicRequest) SetInOrder(inOrder bool) {
    r.InOrder = &inOrder
}
/* param lifeCycle: 保存周期，只能是 7， 15， 30(Optional) */
func (r *CreateLogtopicRequest) SetLifeCycle(lifeCycle int) {
    r.LifeCycle = &lifeCycle
}
/* param tags: 标签列表(Optional) */
func (r *CreateLogtopicRequest) SetTags(tags []logs.Tag) {
    r.Tags = tags
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateLogtopicRequest) GetRegionId() string {
    return r.RegionId
}

type CreateLogtopicResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateLogtopicResult `json:"result"`
}

type CreateLogtopicResult struct {
    UID string `json:"uID"`
}