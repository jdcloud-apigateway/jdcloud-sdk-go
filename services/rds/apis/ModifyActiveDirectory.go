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

type ModifyActiveDirectoryRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 目录服务的资源ID<br>- 加入目录服务：要加入的目录服务的资源ID<br>- 修改目录服务：新目录服务的资源ID<br>- 移除目录服务：传入字符串“none”，不区分大小写  */
    AdResourceId string `json:"adResourceId"`

    /* 修改后，是否强制重启实例，使修改生效。<br> - true 修改后立即重启<br>- false：默认值，修改后不重启，需用户自行重启  */
    ForceRestart string `json:"forceRestart"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param adResourceId: 目录服务的资源ID<br>- 加入目录服务：要加入的目录服务的资源ID<br>- 修改目录服务：新目录服务的资源ID<br>- 移除目录服务：传入字符串“none”，不区分大小写 (Required)
 * param forceRestart: 修改后，是否强制重启实例，使修改生效。<br> - true 修改后立即重启<br>- false：默认值，修改后不重启，需用户自行重启 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyActiveDirectoryRequest(
    regionId string,
    instanceId string,
    adResourceId string,
    forceRestart string,
) *ModifyActiveDirectoryRequest {

	return &ModifyActiveDirectoryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:modifyActiveDirectory",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        AdResourceId: adResourceId,
        ForceRestart: forceRestart,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param adResourceId: 目录服务的资源ID<br>- 加入目录服务：要加入的目录服务的资源ID<br>- 修改目录服务：新目录服务的资源ID<br>- 移除目录服务：传入字符串“none”，不区分大小写 (Required)
 * param forceRestart: 修改后，是否强制重启实例，使修改生效。<br> - true 修改后立即重启<br>- false：默认值，修改后不重启，需用户自行重启 (Required)
 */
func NewModifyActiveDirectoryRequestWithAllParams(
    regionId string,
    instanceId string,
    adResourceId string,
    forceRestart string,
) *ModifyActiveDirectoryRequest {

    return &ModifyActiveDirectoryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyActiveDirectory",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        AdResourceId: adResourceId,
        ForceRestart: forceRestart,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyActiveDirectoryRequestWithoutParam() *ModifyActiveDirectoryRequest {

    return &ModifyActiveDirectoryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyActiveDirectory",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *ModifyActiveDirectoryRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *ModifyActiveDirectoryRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param adResourceId: 目录服务的资源ID<br>- 加入目录服务：要加入的目录服务的资源ID<br>- 修改目录服务：新目录服务的资源ID<br>- 移除目录服务：传入字符串“none”，不区分大小写(Required) */
func (r *ModifyActiveDirectoryRequest) SetAdResourceId(adResourceId string) {
    r.AdResourceId = adResourceId
}
/* param forceRestart: 修改后，是否强制重启实例，使修改生效。<br> - true 修改后立即重启<br>- false：默认值，修改后不重启，需用户自行重启(Required) */
func (r *ModifyActiveDirectoryRequest) SetForceRestart(forceRestart string) {
    r.ForceRestart = forceRestart
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyActiveDirectoryRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyActiveDirectoryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyActiveDirectoryResult `json:"result"`
}

type ModifyActiveDirectoryResult struct {
}