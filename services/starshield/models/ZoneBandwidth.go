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


type ZoneBandwidth struct {

    /* 域名 (Optional) */
    Zonename string `json:"zonename"`

    /* 所属实例名称 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 所属PIN (Optional) */
    Pin string `json:"pin"`

    /* 套餐类型 (Optional) */
    PackType string `json:"packType"`

    /* 计费状态 (Optional) */
    ChargeState string `json:"chargeState"`

    /* 实例到期时间 (Optional) */
    InstanceExpireTime string `json:"instanceExpireTime"`

    /* 查询时间段内的峰值带宽，单位bit per second (Optional) */
    Bps int `json:"bps"`

    /* 查询时间段内的业务峰值带宽，单位bit per second (Optional) */
    NormalBps int `json:"normalBps"`

    /* 查询时间段内的攻击峰值带宽，单位bit per second (Optional) */
    MitigationBps int `json:"mitigationBps"`

    /* 套餐总流量 (Optional) */
    TotalFlowStr string `json:"totalFlowStr"`

    /* 套餐已使用流量 (Optional) */
    UsedFlowStr string `json:"usedFlowStr"`

    /* 套餐剩余流量 (Optional) */
    RemainingFlowStr string `json:"remainingFlowStr"`

    /* 实例使用状态，正常, 超量, 弹性账单 (Optional) */
    InstanceUsedStatus string `json:"instanceUsedStatus"`

    /* 备案状态：false 未备案、true 已备案 (Optional) */
    IcpStatus bool `json:"icpStatus"`
}
