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
    iotcore "github.com/jdcloud-api/jdcloud-sdk-go/services/iotcore/models"
)

type DiscribeThingModelRequest struct {

    core.JDCloudRequest

    /* 区域id  */
    RegionId string `json:"regionId"`

    /* 实例Id  */
    InstanceId string `json:"instanceId"`

    /* 物模型ID编号  */
    ThingModelId string `json:"thingModelId"`

    /* 版本号。如果为空，则返回最新版本 (Optional) */
    ThingModelVersion *string `json:"thingModelVersion"`
}

/*
 * param regionId: 区域id (Required)
 * param instanceId: 实例Id (Required)
 * param thingModelId: 物模型ID编号 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDiscribeThingModelRequest(
    regionId string,
    instanceId string,
    thingModelId string,
) *DiscribeThingModelRequest {

	return &DiscribeThingModelRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/coreinstances/{instanceId}/thingModel:discribeThingModel",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        ThingModelId: thingModelId,
	}
}

/*
 * param regionId: 区域id (Required)
 * param instanceId: 实例Id (Required)
 * param thingModelId: 物模型ID编号 (Required)
 * param thingModelVersion: 版本号。如果为空，则返回最新版本 (Optional)
 */
func NewDiscribeThingModelRequestWithAllParams(
    regionId string,
    instanceId string,
    thingModelId string,
    thingModelVersion *string,
) *DiscribeThingModelRequest {

    return &DiscribeThingModelRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/coreinstances/{instanceId}/thingModel:discribeThingModel",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        ThingModelId: thingModelId,
        ThingModelVersion: thingModelVersion,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDiscribeThingModelRequestWithoutParam() *DiscribeThingModelRequest {

    return &DiscribeThingModelRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/coreinstances/{instanceId}/thingModel:discribeThingModel",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 区域id(Required) */
func (r *DiscribeThingModelRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例Id(Required) */
func (r *DiscribeThingModelRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param thingModelId: 物模型ID编号(Required) */
func (r *DiscribeThingModelRequest) SetThingModelId(thingModelId string) {
    r.ThingModelId = thingModelId
}

/* param thingModelVersion: 版本号。如果为空，则返回最新版本(Optional) */
func (r *DiscribeThingModelRequest) SetThingModelVersion(thingModelVersion string) {
    r.ThingModelVersion = &thingModelVersion
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DiscribeThingModelRequest) GetRegionId() string {
    return r.RegionId
}

type DiscribeThingModelResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DiscribeThingModelResult `json:"result"`
}

type DiscribeThingModelResult struct {
    ThingModelRespTO iotcore.ThingModelRespTO `json:"thingModelRespTO"`
}