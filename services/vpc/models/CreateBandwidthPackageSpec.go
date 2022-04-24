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

package models

import charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type CreateBandwidthPackageSpec struct {

    /* 名称，只支持中文、数字、大小写字母、英文下划线“_”及中划线“-”，且长度不超过32个字符  */
    Name string `json:"name"`

    /* 描述，长度不超过256个字符 (Optional) */
    Description string `json:"description"`

    /* 共享带宽包带宽上限，取值范围200-5000，单位为Mbps，保底带宽 = 共享带宽包带宽上限 * 20%  */
    BandwidthMbps int `json:"bandwidthMbps"`

    /* 线路信息，默认bgp，目前只支持中心节点的BGP线路 (Optional) */
    Provider string `json:"provider"`

    /* 计费配置。支持包年包月、按配置、按用量计费模式 (Optional) */
    ChargeSpec charge.ChargeSpec `json:"chargeSpec"`

    /* 用户标签 (Optional) */
    UserTags []Tag `json:"userTags"`

    /* 资源所属资源组ID (Optional) */
    ResourceGroupId string `json:"resourceGroupId"`
}