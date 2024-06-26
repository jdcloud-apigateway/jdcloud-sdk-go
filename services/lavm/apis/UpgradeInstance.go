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

type UpgradeInstanceRequest struct {

    core.JDCloudRequest

    /* 轻量应用云主机的实例ID
  */
    InstanceId string `json:"instanceId"`

    /* regionId
  */
    RegionId string `json:"regionId"`

    /* 轻量应用云主机方案ID
  */
    PlanId string `json:"planId"`
}

/*
 * param instanceId: 轻量应用云主机的实例ID
 (Required)
 * param regionId: regionId
 (Required)
 * param planId: 轻量应用云主机方案ID
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpgradeInstanceRequest(
    instanceId string,
    regionId string,
    planId string,
) *UpgradeInstanceRequest {

	return &UpgradeInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:upgradeInstance",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        InstanceId: instanceId,
        RegionId: regionId,
        PlanId: planId,
	}
}

/*
 * param instanceId: 轻量应用云主机的实例ID
 (Required)
 * param regionId: regionId
 (Required)
 * param planId: 轻量应用云主机方案ID
 (Required)
 */
func NewUpgradeInstanceRequestWithAllParams(
    instanceId string,
    regionId string,
    planId string,
) *UpgradeInstanceRequest {

    return &UpgradeInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:upgradeInstance",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        InstanceId: instanceId,
        RegionId: regionId,
        PlanId: planId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpgradeInstanceRequestWithoutParam() *UpgradeInstanceRequest {

    return &UpgradeInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:upgradeInstance",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param instanceId: 轻量应用云主机的实例ID
(Required) */
func (r *UpgradeInstanceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param regionId: regionId
(Required) */
func (r *UpgradeInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param planId: 轻量应用云主机方案ID
(Required) */
func (r *UpgradeInstanceRequest) SetPlanId(planId string) {
    r.PlanId = planId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpgradeInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type UpgradeInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpgradeInstanceResult `json:"result"`
}

type UpgradeInstanceResult struct {
}