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


type DeviceResultResp struct {

    /* Id (Optional) */
    Id int `json:"id"`

    /* 应用名称 (Optional) */
    AppName string `json:"appName"`

    /* 操作系统 (Optional) */
    System string `json:"system"`

    /* Eid (Optional) */
    Eid string `json:"eid"`

    /* 设备风险 (Optional) */
    RiskTag string `json:"riskTag"`

    /* 创建时间，毫秒级时间戳 (Optional) */
    CreateTime int `json:"createTime"`

    /* 最新采集时间，毫秒级时间戳 (Optional) */
    UpdateTime int `json:"updateTime"`

    /* 采集次数 (Optional) */
    Count int `json:"count"`
}
