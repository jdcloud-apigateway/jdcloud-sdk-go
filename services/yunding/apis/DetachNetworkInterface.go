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

type DetachNetworkInterfaceRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 云主机ID  */
    InstanceId string `json:"instanceId"`

    /* 弹性网卡ID  */
    NetworkInterfaceId string `json:"networkInterfaceId"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: 云主机ID (Required)
 * param networkInterfaceId: 弹性网卡ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDetachNetworkInterfaceRequest(
    regionId string,
    instanceId string,
    networkInterfaceId string,
) *DetachNetworkInterfaceRequest {

	return &DetachNetworkInterfaceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ydVmInstances/{instanceId}:detachNetworkInterface",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        NetworkInterfaceId: networkInterfaceId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: 云主机ID (Required)
 * param networkInterfaceId: 弹性网卡ID (Required)
 */
func NewDetachNetworkInterfaceRequestWithAllParams(
    regionId string,
    instanceId string,
    networkInterfaceId string,
) *DetachNetworkInterfaceRequest {

    return &DetachNetworkInterfaceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ydVmInstances/{instanceId}:detachNetworkInterface",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        NetworkInterfaceId: networkInterfaceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDetachNetworkInterfaceRequestWithoutParam() *DetachNetworkInterfaceRequest {

    return &DetachNetworkInterfaceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ydVmInstances/{instanceId}:detachNetworkInterface",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DetachNetworkInterfaceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: 云主机ID(Required) */
func (r *DetachNetworkInterfaceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param networkInterfaceId: 弹性网卡ID(Required) */
func (r *DetachNetworkInterfaceRequest) SetNetworkInterfaceId(networkInterfaceId string) {
    r.NetworkInterfaceId = networkInterfaceId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DetachNetworkInterfaceRequest) GetRegionId() string {
    return r.RegionId
}

type DetachNetworkInterfaceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DetachNetworkInterfaceResult `json:"result"`
}

type DetachNetworkInterfaceResult struct {
}