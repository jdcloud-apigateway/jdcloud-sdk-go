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

type DescribePrivilegeRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 设置可见的引擎类型，如 MySQL 等  */
    Engine string `json:"engine"`

    /* RDS 实例ID，唯一标识一个RDS实例 (Optional) */
    InstanceId *string `json:"instanceId"`

    /* true表示展示高权限，默认false (Optional) */
    AllAdminPrivileges *bool `json:"allAdminPrivileges"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param engine: 设置可见的引擎类型，如 MySQL 等 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribePrivilegeRequest(
    regionId string,
    engine string,
) *DescribePrivilegeRequest {

	return &DescribePrivilegeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/common:describePrivilege",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Engine: engine,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param engine: 设置可见的引擎类型，如 MySQL 等 (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Optional)
 * param allAdminPrivileges: true表示展示高权限，默认false (Optional)
 */
func NewDescribePrivilegeRequestWithAllParams(
    regionId string,
    engine string,
    instanceId *string,
    allAdminPrivileges *bool,
) *DescribePrivilegeRequest {

    return &DescribePrivilegeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/common:describePrivilege",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Engine: engine,
        InstanceId: instanceId,
        AllAdminPrivileges: allAdminPrivileges,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribePrivilegeRequestWithoutParam() *DescribePrivilegeRequest {

    return &DescribePrivilegeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/common:describePrivilege",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DescribePrivilegeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param engine: 设置可见的引擎类型，如 MySQL 等(Required) */
func (r *DescribePrivilegeRequest) SetEngine(engine string) {
    r.Engine = engine
}
/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Optional) */
func (r *DescribePrivilegeRequest) SetInstanceId(instanceId string) {
    r.InstanceId = &instanceId
}
/* param allAdminPrivileges: true表示展示高权限，默认false(Optional) */
func (r *DescribePrivilegeRequest) SetAllAdminPrivileges(allAdminPrivileges bool) {
    r.AllAdminPrivileges = &allAdminPrivileges
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribePrivilegeRequest) GetRegionId() string {
    return r.RegionId
}

type DescribePrivilegeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribePrivilegeResult `json:"result"`
}

type DescribePrivilegeResult struct {
    GlobalPrivileges []string `json:"globalPrivileges"`
    DatabasePrivileges []string `json:"databasePrivileges"`
    TablePrivileges []string `json:"tablePrivileges"`
}