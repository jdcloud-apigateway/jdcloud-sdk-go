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

type ModifyFileSystemAttributeRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 文件系统ID  */
    FileSystemId string `json:"fileSystemId"`

    /* 文件系统名称(参数规则：不可为空，只支持中文、数字、大小写字母、英文下划线“_”及中划线“-”，且不能超过32字符) (Optional) */
    Name *string `json:"name"`

    /* 文件系统描述(参数规则：不能超过256字符) (Optional) */
    Description *string `json:"description"`

}

/*
 * param regionId: 地域ID (Required)
 * param fileSystemId: 文件系统ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyFileSystemAttributeRequest(
    regionId string,
    fileSystemId string,
) *ModifyFileSystemAttributeRequest {

	return &ModifyFileSystemAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/fileSystems/{fileSystemId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        FileSystemId: fileSystemId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param fileSystemId: 文件系统ID (Required)
 * param name: 文件系统名称(参数规则：不可为空，只支持中文、数字、大小写字母、英文下划线“_”及中划线“-”，且不能超过32字符) (Optional)
 * param description: 文件系统描述(参数规则：不能超过256字符) (Optional)
 */
func NewModifyFileSystemAttributeRequestWithAllParams(
    regionId string,
    fileSystemId string,
    name *string,
    description *string,
) *ModifyFileSystemAttributeRequest {

    return &ModifyFileSystemAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/fileSystems/{fileSystemId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        FileSystemId: fileSystemId,
        Name: name,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyFileSystemAttributeRequestWithoutParam() *ModifyFileSystemAttributeRequest {

    return &ModifyFileSystemAttributeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/fileSystems/{fileSystemId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ModifyFileSystemAttributeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param fileSystemId: 文件系统ID(Required) */
func (r *ModifyFileSystemAttributeRequest) SetFileSystemId(fileSystemId string) {
    r.FileSystemId = fileSystemId
}

/* param name: 文件系统名称(参数规则：不可为空，只支持中文、数字、大小写字母、英文下划线“_”及中划线“-”，且不能超过32字符)(Optional) */
func (r *ModifyFileSystemAttributeRequest) SetName(name string) {
    r.Name = &name
}

/* param description: 文件系统描述(参数规则：不能超过256字符)(Optional) */
func (r *ModifyFileSystemAttributeRequest) SetDescription(description string) {
    r.Description = &description
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyFileSystemAttributeRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyFileSystemAttributeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyFileSystemAttributeResult `json:"result"`
}

type ModifyFileSystemAttributeResult struct {
}