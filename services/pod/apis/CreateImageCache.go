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
    pod "github.com/jdcloud-api/jdcloud-sdk-go/services/pod/models"
)

type CreateImageCacheRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 镜像缓存名称 (Optional) */
    Name *string `json:"name"`

    /* 镜像缓存创建参数  */
    ImageCacheSpec *pod.ImageCacheSpec `json:"imageCacheSpec"`
}

/*
 * param regionId: Region ID (Required)
 * param imageCacheSpec: 镜像缓存创建参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateImageCacheRequest(
    regionId string,
    imageCacheSpec *pod.ImageCacheSpec,
) *CreateImageCacheRequest {

	return &CreateImageCacheRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/imageCache",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ImageCacheSpec: imageCacheSpec,
	}
}

/*
 * param regionId: Region ID (Required)
 * param name: 镜像缓存名称 (Optional)
 * param imageCacheSpec: 镜像缓存创建参数 (Required)
 */
func NewCreateImageCacheRequestWithAllParams(
    regionId string,
    name *string,
    imageCacheSpec *pod.ImageCacheSpec,
) *CreateImageCacheRequest {

    return &CreateImageCacheRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/imageCache",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Name: name,
        ImageCacheSpec: imageCacheSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateImageCacheRequestWithoutParam() *CreateImageCacheRequest {

    return &CreateImageCacheRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/imageCache",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateImageCacheRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param name: 镜像缓存名称(Optional) */
func (r *CreateImageCacheRequest) SetName(name string) {
    r.Name = &name
}
/* param imageCacheSpec: 镜像缓存创建参数(Required) */
func (r *CreateImageCacheRequest) SetImageCacheSpec(imageCacheSpec *pod.ImageCacheSpec) {
    r.ImageCacheSpec = imageCacheSpec
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateImageCacheRequest) GetRegionId() string {
    return r.RegionId
}

type CreateImageCacheResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateImageCacheResult `json:"result"`
}

type CreateImageCacheResult struct {
    ImageCacheId []string `json:"imageCacheId"`
}