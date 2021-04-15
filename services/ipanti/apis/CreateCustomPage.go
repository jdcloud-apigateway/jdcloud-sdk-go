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
    ipanti "github.com/jdcloud-api/jdcloud-sdk-go/services/ipanti/models"
)

type CreateCustomPageRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`

    /* 实例 ID  */
    InstanceId string `json:"instanceId"`

    /* 添加自定义页面请求参数  */
    CustomPageSpec *ipanti.CustomPageSpec `json:"customPageSpec"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 实例 ID (Required)
 * param customPageSpec: 添加自定义页面请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateCustomPageRequest(
    regionId string,
    instanceId string,
    customPageSpec *ipanti.CustomPageSpec,
) *CreateCustomPageRequest {

	return &CreateCustomPageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/customPages",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        CustomPageSpec: customPageSpec,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 实例 ID (Required)
 * param customPageSpec: 添加自定义页面请求参数 (Required)
 */
func NewCreateCustomPageRequestWithAllParams(
    regionId string,
    instanceId string,
    customPageSpec *ipanti.CustomPageSpec,
) *CreateCustomPageRequest {

    return &CreateCustomPageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/customPages",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        CustomPageSpec: customPageSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateCustomPageRequestWithoutParam() *CreateCustomPageRequest {

    return &CreateCustomPageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/customPages",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *CreateCustomPageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例 ID(Required) */
func (r *CreateCustomPageRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param customPageSpec: 添加自定义页面请求参数(Required) */
func (r *CreateCustomPageRequest) SetCustomPageSpec(customPageSpec *ipanti.CustomPageSpec) {
    r.CustomPageSpec = customPageSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateCustomPageRequest) GetRegionId() string {
    return r.RegionId
}

type CreateCustomPageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateCustomPageResult `json:"result"`
}

type CreateCustomPageResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}