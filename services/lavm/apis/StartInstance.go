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

type StartInstanceRequest struct {

    core.JDCloudRequest

    /* 轻量应用云主机的实例ID
  */
    InstanceId string `json:"instanceId"`

    /* regionId
  */
    RegionId string `json:"regionId"`
}

/*
 * param instanceId: 轻量应用云主机的实例ID
 (Required)
 * param regionId: regionId
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewStartInstanceRequest(
    instanceId string,
    regionId string,
) *StartInstanceRequest {

	return &StartInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:startInstance",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        InstanceId: instanceId,
        RegionId: regionId,
	}
}

/*
 * param instanceId: 轻量应用云主机的实例ID
 (Required)
 * param regionId: regionId
 (Required)
 */
func NewStartInstanceRequestWithAllParams(
    instanceId string,
    regionId string,
) *StartInstanceRequest {

    return &StartInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:startInstance",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        InstanceId: instanceId,
        RegionId: regionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewStartInstanceRequestWithoutParam() *StartInstanceRequest {

    return &StartInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:startInstance",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param instanceId: 轻量应用云主机的实例ID
(Required) */
func (r *StartInstanceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param regionId: regionId
(Required) */
func (r *StartInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r StartInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type StartInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result StartInstanceResult `json:"result"`
}

type StartInstanceResult struct {
}