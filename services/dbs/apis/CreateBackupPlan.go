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
    dbs "github.com/jdcloud-api/jdcloud-sdk-go/services/dbs/models"
)

type CreateBackupPlanRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》]  */
    RegionId string `json:"regionId"`

    /* 备份计划名称，支持中文、数字、大小写字母、英文下划线“_”、减号“-”，且不少于2字符不超过64字符  */
    Name string `json:"name"`

    /* DBS服务包类型是枚举值， dbs.common.package 表示基础服务包，不含备份流量  */
    ServicePackage string `json:"servicePackage"`

    /* 购买规格  */
    ChargeSpec *dbs.ChargeSpec `json:"chargeSpec"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param name: 备份计划名称，支持中文、数字、大小写字母、英文下划线“_”、减号“-”，且不少于2字符不超过64字符 (Required)
 * param servicePackage: DBS服务包类型是枚举值， dbs.common.package 表示基础服务包，不含备份流量 (Required)
 * param chargeSpec: 购买规格 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateBackupPlanRequest(
    regionId string,
    name string,
    servicePackage string,
    chargeSpec *dbs.ChargeSpec,
) *CreateBackupPlanRequest {

	return &CreateBackupPlanRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/backupPlans",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        Name: name,
        ServicePackage: servicePackage,
        ChargeSpec: chargeSpec,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param name: 备份计划名称，支持中文、数字、大小写字母、英文下划线“_”、减号“-”，且不少于2字符不超过64字符 (Required)
 * param servicePackage: DBS服务包类型是枚举值， dbs.common.package 表示基础服务包，不含备份流量 (Required)
 * param chargeSpec: 购买规格 (Required)
 */
func NewCreateBackupPlanRequestWithAllParams(
    regionId string,
    name string,
    servicePackage string,
    chargeSpec *dbs.ChargeSpec,
) *CreateBackupPlanRequest {

    return &CreateBackupPlanRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupPlans",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        Name: name,
        ServicePackage: servicePackage,
        ChargeSpec: chargeSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateBackupPlanRequestWithoutParam() *CreateBackupPlanRequest {

    return &CreateBackupPlanRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupPlans",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](Required) */
func (r *CreateBackupPlanRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param name: 备份计划名称，支持中文、数字、大小写字母、英文下划线“_”、减号“-”，且不少于2字符不超过64字符(Required) */
func (r *CreateBackupPlanRequest) SetName(name string) {
    r.Name = name
}

/* param servicePackage: DBS服务包类型是枚举值， dbs.common.package 表示基础服务包，不含备份流量(Required) */
func (r *CreateBackupPlanRequest) SetServicePackage(servicePackage string) {
    r.ServicePackage = servicePackage
}

/* param chargeSpec: 购买规格(Required) */
func (r *CreateBackupPlanRequest) SetChargeSpec(chargeSpec *dbs.ChargeSpec) {
    r.ChargeSpec = chargeSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateBackupPlanRequest) GetRegionId() string {
    return r.RegionId
}

type CreateBackupPlanResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateBackupPlanResult `json:"result"`
}

type CreateBackupPlanResult struct {
    BuyId string `json:"buyId"`
    OrderNubmer string `json:"orderNubmer"`
    BackupPlanId string `json:"backupPlanId"`
}