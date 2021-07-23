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
    vod "github.com/jdcloud-api/jdcloud-sdk-go/services/vod/models"
)

type UpdateVeditProjectRequest struct {

    core.JDCloudRequest

    /* 视频剪辑工程ID  */
    ProjectId int `json:"projectId"`

    /* 工程名称 (Optional) */
    ProjectName *string `json:"projectName"`

    /* 工程描述 (Optional) */
    Description *string `json:"description"`

    /* 时间线信息 (Optional) */
    Timeline *vod.Timeline `json:"timeline"`
}

/*
 * param projectId: 视频剪辑工程ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateVeditProjectRequest(
    projectId int,
) *UpdateVeditProjectRequest {

	return &UpdateVeditProjectRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/veditProjects/{projectId}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        ProjectId: projectId,
	}
}

/*
 * param projectId: 视频剪辑工程ID (Required)
 * param projectName: 工程名称 (Optional)
 * param description: 工程描述 (Optional)
 * param timeline: 时间线信息 (Optional)
 */
func NewUpdateVeditProjectRequestWithAllParams(
    projectId int,
    projectName *string,
    description *string,
    timeline *vod.Timeline,
) *UpdateVeditProjectRequest {

    return &UpdateVeditProjectRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/veditProjects/{projectId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        ProjectId: projectId,
        ProjectName: projectName,
        Description: description,
        Timeline: timeline,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateVeditProjectRequestWithoutParam() *UpdateVeditProjectRequest {

    return &UpdateVeditProjectRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/veditProjects/{projectId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param projectId: 视频剪辑工程ID(Required) */
func (r *UpdateVeditProjectRequest) SetProjectId(projectId int) {
    r.ProjectId = projectId
}

/* param projectName: 工程名称(Optional) */
func (r *UpdateVeditProjectRequest) SetProjectName(projectName string) {
    r.ProjectName = &projectName
}

/* param description: 工程描述(Optional) */
func (r *UpdateVeditProjectRequest) SetDescription(description string) {
    r.Description = &description
}

/* param timeline: 时间线信息(Optional) */
func (r *UpdateVeditProjectRequest) SetTimeline(timeline *vod.Timeline) {
    r.Timeline = timeline
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateVeditProjectRequest) GetRegionId() string {
    return ""
}

type UpdateVeditProjectResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateVeditProjectResult `json:"result"`
}

type UpdateVeditProjectResult struct {
    ProjectId int64 `json:"projectId"`
    ProjectName string `json:"projectName"`
    Description string `json:"description"`
    Timeline vod.Timeline `json:"timeline"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}