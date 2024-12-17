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


type OpInstancesRes struct {

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 套餐类型 (Optional) */
    PackType string `json:"packType"`

    /* 计费状态 (Optional) */
    ChargeState string `json:"chargeState"`

    /* 域名增量包数量 (Optional) */
    ZonePackNum int `json:"zonePackNum"`

    /* 流量包数量 (Optional) */
    TrafficExpansion int `json:"trafficExpansion"`

    /* 已使用流量（单位：Byte） (Optional) */
    FlowUsedCnt int `json:"flowUsedCnt"`

    /* 剩余流量(单位：Gb) (Optional) */
    FlowRemain float64 `json:"flowRemain"`

    /* 套餐总流量 (Optional) */
    TotalFlowStr string `json:"totalFlowStr"`

    /* 套餐已使用流量 (Optional) */
    UsedFlowStr string `json:"usedFlowStr"`

    /* 套餐剩余流量 (Optional) */
    RemainingFlowStr string `json:"remainingFlowStr"`

    /* 套餐模式(BASE->基础套餐 FLOW->流量套餐) (Optional) */
    PackMode string `json:"packMode"`

    /* 备注 (Optional) */
    Memo string `json:"memo"`

    /* 购买时间, UTC时间格式，例如2017-11-10T23:00:00Z (Optional) */
    CreateTime string `json:"createTime"`

    /* 到期时间, UTC时间格式，例如2017-11-10T23:00:00Z (Optional) */
    ExpireTime string `json:"expireTime"`

    /* pin (Optional) */
    Pin string `json:"pin"`

    /* 是否企业实名认证(false->未认证 ture->已认证) (Optional) */
    RealName bool `json:"realName"`

    /* 域名信息 (Optional) */
    Zones []OpInstanceZoneRes `json:"zones"`

    /* 域名总数 (Optional) */
    ZoneNum int `json:"zoneNum"`

    /* 域名已激活数 (Optional) */
    ZoneActiveNum int `json:"zoneActiveNum"`
}