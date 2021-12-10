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
    starshield "github.com/jdcloud-api/jdcloud-sdk-go/services/starshield/models"
)

type PurgeFilesByURLRequest struct {

    core.JDCloudRequest

    /*   */
    Identifier string `json:"identifier"`

    /* 应从缓存中删除的URL数组 (Optional) */
    Files []string `json:"files"`
}

/*
 * param identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewPurgeFilesByURLRequest(
    identifier string,
) *PurgeFilesByURLRequest {

	return &PurgeFilesByURLRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{identifier}/purge_cache__purgeFilesByURL",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Identifier: identifier,
	}
}

/*
 * param identifier:  (Required)
 * param files: 应从缓存中删除的URL数组 (Optional)
 */
func NewPurgeFilesByURLRequestWithAllParams(
    identifier string,
    files []string,
) *PurgeFilesByURLRequest {

    return &PurgeFilesByURLRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{identifier}/purge_cache__purgeFilesByURL",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Identifier: identifier,
        Files: files,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewPurgeFilesByURLRequestWithoutParam() *PurgeFilesByURLRequest {

    return &PurgeFilesByURLRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{identifier}/purge_cache__purgeFilesByURL",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param identifier: (Required) */
func (r *PurgeFilesByURLRequest) SetIdentifier(identifier string) {
    r.Identifier = identifier
}

/* param files: 应从缓存中删除的URL数组(Optional) */
func (r *PurgeFilesByURLRequest) SetFiles(files []string) {
    r.Files = files
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r PurgeFilesByURLRequest) GetRegionId() string {
    return ""
}

type PurgeFilesByURLResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result PurgeFilesByURLResult `json:"result"`
}

type PurgeFilesByURLResult struct {
    Data starshield.Zone `json:"data"`
}