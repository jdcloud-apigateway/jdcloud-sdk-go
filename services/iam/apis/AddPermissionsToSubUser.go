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

type AddPermissionsToSubUserRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 子用户用户名  */
    SubUser string `json:"subUser"`

    /* 权限信息  */
    AddPermissionsInfo *iam.AddPermissionsInfo `json:"addPermissionsInfo"`
}

/*
 * param regionId: Region ID (Required)
 * param subUser: 子用户用户名 (Required)
 * param addPermissionsInfo: 权限信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddPermissionsToSubUserRequest(
    regionId string,
    subUser string,
    addPermissionsInfo *iam.AddPermissionsInfo,
) *AddPermissionsToSubUserRequest {

	return &AddPermissionsToSubUserRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/subUser/{subUser}/permisssions",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SubUser: subUser,
        AddPermissionsInfo: addPermissionsInfo,
	}
}

/*
 * param regionId: Region ID (Required)
 * param subUser: 子用户用户名 (Required)
 * param addPermissionsInfo: 权限信息 (Required)
 */
func NewAddPermissionsToSubUserRequestWithAllParams(
    regionId string,
    subUser string,
    addPermissionsInfo *iam.AddPermissionsInfo,
) *AddPermissionsToSubUserRequest {

    return &AddPermissionsToSubUserRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subUser/{subUser}/permisssions",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SubUser: subUser,
        AddPermissionsInfo: addPermissionsInfo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddPermissionsToSubUserRequestWithoutParam() *AddPermissionsToSubUserRequest {

    return &AddPermissionsToSubUserRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subUser/{subUser}/permisssions",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *AddPermissionsToSubUserRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param subUser: 子用户用户名(Required) */
func (r *AddPermissionsToSubUserRequest) SetSubUser(subUser string) {
    r.SubUser = subUser
}
/* param addPermissionsInfo: 权限信息(Required) */
func (r *AddPermissionsToSubUserRequest) SetAddPermissionsInfo(addPermissionsInfo *iam.AddPermissionsInfo) {
    r.AddPermissionsInfo = addPermissionsInfo
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddPermissionsToSubUserRequest) GetRegionId() string {
    return r.RegionId
}

type AddPermissionsToSubUserResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddPermissionsToSubUserResult `json:"result"`
}

type AddPermissionsToSubUserResult struct {
}