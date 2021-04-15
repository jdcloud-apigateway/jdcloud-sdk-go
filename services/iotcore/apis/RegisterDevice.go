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

type RegisterDeviceRequest struct {

    core.JDCloudRequest

    /* 区域id  */
    RegionId string `json:"regionId"`

    /* 实例Id  */
    InstanceId string `json:"instanceId"`

    /* 物模型ID编号  */
    DeviceInfoVO *iotcore.DeviceInfoVO `json:"deviceInfoVO"`
}

/*
 * param regionId: 区域id (Required)
 * param instanceId: 实例Id (Required)
 * param deviceInfoVO: 物模型ID编号 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRegisterDeviceRequest(
    regionId string,
    instanceId string,
    deviceInfoVO *iotcore.DeviceInfoVO,
) *RegisterDeviceRequest {

	return &RegisterDeviceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/coreinstances/{instanceId}/device:register",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        DeviceInfoVO: deviceInfoVO,
	}
}

/*
 * param regionId: 区域id (Required)
 * param instanceId: 实例Id (Required)
 * param deviceInfoVO: 物模型ID编号 (Required)
 */
func NewRegisterDeviceRequestWithAllParams(
    regionId string,
    instanceId string,
    deviceInfoVO *iotcore.DeviceInfoVO,
) *RegisterDeviceRequest {

    return &RegisterDeviceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/coreinstances/{instanceId}/device:register",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        DeviceInfoVO: deviceInfoVO,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRegisterDeviceRequestWithoutParam() *RegisterDeviceRequest {

    return &RegisterDeviceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/coreinstances/{instanceId}/device:register",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 区域id(Required) */
func (r *RegisterDeviceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例Id(Required) */
func (r *RegisterDeviceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param deviceInfoVO: 物模型ID编号(Required) */
func (r *RegisterDeviceRequest) SetDeviceInfoVO(deviceInfoVO *iotcore.DeviceInfoVO) {
    r.DeviceInfoVO = deviceInfoVO
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RegisterDeviceRequest) GetRegionId() string {
    return r.RegionId
}

type RegisterDeviceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RegisterDeviceResult `json:"result"`
}

type RegisterDeviceResult struct {
    DeviceId string `json:"deviceId"`
}