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


type InstanceTrafficPackageUsage struct {

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 流量包当月周期内已使用流量 (Optional) */
    TrafficUsed int `json:"trafficUsed"`

    /* 流量包当月周期内的总流量 (Optional) */
    TrafficPackageTotal int `json:"trafficPackageTotal"`

    /* 流量包当月周期内的剩余流量 (Optional) */
    TrafficPackageRemaining int `json:"trafficPackageRemaining"`

    /* 流量包当月周期内超出流量包额度的流量 (Optional) */
    TrafficOverflow int `json:"trafficOverflow"`
}
