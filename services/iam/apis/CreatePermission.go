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

type CreatePermissionRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 权限信息  */
    CreatePermissionInfo *iam.CreatePermissionInfo `json:"createPermissionInfo"`
}

/*
 * param regionId: Region ID (Required)
 * param createPermissionInfo: 权限信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreatePermissionRequest(
    regionId string,
    createPermissionInfo *iam.CreatePermissionInfo,
) *CreatePermissionRequest {

	return &CreatePermissionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/permission",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CreatePermissionInfo: createPermissionInfo,
	}
}

/*
 * param regionId: Region ID (Required)
 * param createPermissionInfo: 权限信息 (Required)
 */
func NewCreatePermissionRequestWithAllParams(
    regionId string,
    createPermissionInfo *iam.CreatePermissionInfo,
) *CreatePermissionRequest {

    return &CreatePermissionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/permission",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CreatePermissionInfo: createPermissionInfo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreatePermissionRequestWithoutParam() *CreatePermissionRequest {

    return &CreatePermissionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/permission",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreatePermissionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param createPermissionInfo: 权限信息(Required) */
func (r *CreatePermissionRequest) SetCreatePermissionInfo(createPermissionInfo *iam.CreatePermissionInfo) {
    r.CreatePermissionInfo = createPermissionInfo
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreatePermissionRequest) GetRegionId() string {
    return r.RegionId
}

type CreatePermissionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreatePermissionResult `json:"result"`
}

type CreatePermissionResult struct {
}