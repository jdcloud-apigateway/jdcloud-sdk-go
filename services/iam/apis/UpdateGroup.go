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
    iam "github.com/jdcloud-api/jdcloud-sdk-go/services/iam/models"
)

type UpdateGroupRequest struct {

    core.JDCloudRequest

    /* 用户组名称  */
    GroupName string `json:"groupName"`

    /*   */
    UpdateGroupInfo *iam.UpdateGroupInfo `json:"updateGroupInfo"`
}

/*
 * param groupName: 用户组名称 (Required)
 * param updateGroupInfo:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateGroupRequest(
    groupName string,
    updateGroupInfo *iam.UpdateGroupInfo,
) *UpdateGroupRequest {

	return &UpdateGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/group/{groupName}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        GroupName: groupName,
        UpdateGroupInfo: updateGroupInfo,
	}
}

/*
 * param groupName: 用户组名称 (Required)
 * param updateGroupInfo:  (Required)
 */
func NewUpdateGroupRequestWithAllParams(
    groupName string,
    updateGroupInfo *iam.UpdateGroupInfo,
) *UpdateGroupRequest {

    return &UpdateGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/group/{groupName}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        GroupName: groupName,
        UpdateGroupInfo: updateGroupInfo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateGroupRequestWithoutParam() *UpdateGroupRequest {

    return &UpdateGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/group/{groupName}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param groupName: 用户组名称(Required) */
func (r *UpdateGroupRequest) SetGroupName(groupName string) {
    r.GroupName = groupName
}
/* param updateGroupInfo: (Required) */
func (r *UpdateGroupRequest) SetUpdateGroupInfo(updateGroupInfo *iam.UpdateGroupInfo) {
    r.UpdateGroupInfo = updateGroupInfo
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateGroupRequest) GetRegionId() string {
    return ""
}

type UpdateGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateGroupResult `json:"result"`
}

type UpdateGroupResult struct {
}