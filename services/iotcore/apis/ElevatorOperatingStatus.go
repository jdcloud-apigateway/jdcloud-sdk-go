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

type ElevatorOperatingStatusRequest struct {

    core.JDCloudRequest

    /* Hub实例Id  */
    InstanceId string `json:"instanceId"`

    /* 区域Id  */
    RegionId string `json:"regionId"`

    /* 电梯连接码 (Optional) */
    Identifier *string `json:"identifier"`
}

/*
 * param instanceId: Hub实例Id (Required)
 * param regionId: 区域Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewElevatorOperatingStatusRequest(
    instanceId string,
    regionId string,
) *ElevatorOperatingStatusRequest {

	return &ElevatorOperatingStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/loongrayinstances/{instanceId}/registerElevatorOperatingStatus",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        InstanceId: instanceId,
        RegionId: regionId,
	}
}

/*
 * param instanceId: Hub实例Id (Required)
 * param regionId: 区域Id (Required)
 * param identifier: 电梯连接码 (Optional)
 */
func NewElevatorOperatingStatusRequestWithAllParams(
    instanceId string,
    regionId string,
    identifier *string,
) *ElevatorOperatingStatusRequest {

    return &ElevatorOperatingStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loongrayinstances/{instanceId}/registerElevatorOperatingStatus",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        InstanceId: instanceId,
        RegionId: regionId,
        Identifier: identifier,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewElevatorOperatingStatusRequestWithoutParam() *ElevatorOperatingStatusRequest {

    return &ElevatorOperatingStatusRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loongrayinstances/{instanceId}/registerElevatorOperatingStatus",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param instanceId: Hub实例Id(Required) */
func (r *ElevatorOperatingStatusRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param regionId: 区域Id(Required) */
func (r *ElevatorOperatingStatusRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param identifier: 电梯连接码(Optional) */
func (r *ElevatorOperatingStatusRequest) SetIdentifier(identifier string) {
    r.Identifier = &identifier
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ElevatorOperatingStatusRequest) GetRegionId() string {
    return r.RegionId
}

type ElevatorOperatingStatusResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ElevatorOperatingStatusResult `json:"result"`
}

type ElevatorOperatingStatusResult struct {
    Data []iotcore.RegisterValue `json:"data"`
    TotalCount int `json:"totalCount"`
}