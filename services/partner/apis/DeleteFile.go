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

type DeleteFileRequest struct {

    core.JDCloudRequest

    /* 区域(如:cn-north-1)  */
    RegionId string `json:"regionId"`

    /* 合同附件唯一标识  */
    FileUuid string `json:"fileUuid"`
}

/*
 * param regionId: 区域(如:cn-north-1) (Required)
 * param fileUuid: 合同附件唯一标识 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteFileRequest(
    regionId string,
    fileUuid string,
) *DeleteFileRequest {

	return &DeleteFileRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/fileUuid/{fileUuid}/Contract:deleteFile",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        FileUuid: fileUuid,
	}
}

/*
 * param regionId: 区域(如:cn-north-1) (Required)
 * param fileUuid: 合同附件唯一标识 (Required)
 */
func NewDeleteFileRequestWithAllParams(
    regionId string,
    fileUuid string,
) *DeleteFileRequest {

    return &DeleteFileRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/fileUuid/{fileUuid}/Contract:deleteFile",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        FileUuid: fileUuid,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteFileRequestWithoutParam() *DeleteFileRequest {

    return &DeleteFileRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/fileUuid/{fileUuid}/Contract:deleteFile",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域(如:cn-north-1)(Required) */
func (r *DeleteFileRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param fileUuid: 合同附件唯一标识(Required) */
func (r *DeleteFileRequest) SetFileUuid(fileUuid string) {
    r.FileUuid = fileUuid
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteFileRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteFileResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteFileResult `json:"result"`
}

type DeleteFileResult struct {
    Result bool `json:"result"`
}