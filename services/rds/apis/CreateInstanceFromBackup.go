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
    rds "github.com/jdcloud-api/jdcloud-sdk-go/services/rds/models"
)

type CreateInstanceFromBackupRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 备份ID  */
    BackupId string `json:"backupId"`

    /* 标识是创建什么类型的实例，例如MySQL，SQL Server等,具体可参见文档[枚举参数定义](../Enum-Definitions/Enum-Definitions.md)<br>**注意：备份来源实例的engine和要创建的实例的engine必须一致**  */
    Engine string `json:"engine"`

    /* 新建实例规格  */
    InstanceSpec *rds.RestoredNewDBInstanceSpec `json:"instanceSpec"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param backupId: 备份ID (Required)
 * param engine: 标识是创建什么类型的实例，例如MySQL，SQL Server等,具体可参见文档[枚举参数定义](../Enum-Definitions/Enum-Definitions.md)<br>**注意：备份来源实例的engine和要创建的实例的engine必须一致** (Required)
 * param instanceSpec: 新建实例规格 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateInstanceFromBackupRequest(
    regionId string,
    backupId string,
    engine string,
    instanceSpec *rds.RestoredNewDBInstanceSpec,
) *CreateInstanceFromBackupRequest {

	return &CreateInstanceFromBackupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances:createInstanceFromBackup",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BackupId: backupId,
        Engine: engine,
        InstanceSpec: instanceSpec,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param backupId: 备份ID (Required)
 * param engine: 标识是创建什么类型的实例，例如MySQL，SQL Server等,具体可参见文档[枚举参数定义](../Enum-Definitions/Enum-Definitions.md)<br>**注意：备份来源实例的engine和要创建的实例的engine必须一致** (Required)
 * param instanceSpec: 新建实例规格 (Required)
 */
func NewCreateInstanceFromBackupRequestWithAllParams(
    regionId string,
    backupId string,
    engine string,
    instanceSpec *rds.RestoredNewDBInstanceSpec,
) *CreateInstanceFromBackupRequest {

    return &CreateInstanceFromBackupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances:createInstanceFromBackup",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BackupId: backupId,
        Engine: engine,
        InstanceSpec: instanceSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateInstanceFromBackupRequestWithoutParam() *CreateInstanceFromBackupRequest {

    return &CreateInstanceFromBackupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances:createInstanceFromBackup",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *CreateInstanceFromBackupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param backupId: 备份ID(Required) */
func (r *CreateInstanceFromBackupRequest) SetBackupId(backupId string) {
    r.BackupId = backupId
}

/* param engine: 标识是创建什么类型的实例，例如MySQL，SQL Server等,具体可参见文档[枚举参数定义](../Enum-Definitions/Enum-Definitions.md)<br>**注意：备份来源实例的engine和要创建的实例的engine必须一致**(Required) */
func (r *CreateInstanceFromBackupRequest) SetEngine(engine string) {
    r.Engine = engine
}

/* param instanceSpec: 新建实例规格(Required) */
func (r *CreateInstanceFromBackupRequest) SetInstanceSpec(instanceSpec *rds.RestoredNewDBInstanceSpec) {
    r.InstanceSpec = instanceSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateInstanceFromBackupRequest) GetRegionId() string {
    return r.RegionId
}

type CreateInstanceFromBackupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateInstanceFromBackupResult `json:"result"`
}

type CreateInstanceFromBackupResult struct {
    InstanceId string `json:"instanceId"`
    OrderId string `json:"orderId"`
}