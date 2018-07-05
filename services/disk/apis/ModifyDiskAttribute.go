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

type ModifyDiskAttributeRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 云硬盘ID  */
    DiskId string `json:"diskId"`

    /* 云硬盘名称 (Optional) */
    Name *string `json:"name"`

    /* 云硬盘描述，name和description必须要指定一个 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: 地域ID 
 * param diskId: 云硬盘ID 
 * param name: 云硬盘名称 (Optional)
 * param description: 云硬盘描述，name和description必须要指定一个 (Optional)
 */
func NewModifyDiskAttributeRequest(
    regionId string,
    diskId string,
) *ModifyDiskAttributeRequest {

	return &ModifyDiskAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/disks/{diskId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DiskId: diskId,
	}
}

func (r *ModifyDiskAttributeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *ModifyDiskAttributeRequest) SetDiskId(diskId string) {
    r.DiskId = diskId
}

func (r *ModifyDiskAttributeRequest) SetName(name string) {
    r.Name = &name
}

func (r *ModifyDiskAttributeRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyDiskAttributeRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyDiskAttributeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyDiskAttributeResult `json:"result"`
}

type ModifyDiskAttributeResult struct {
}