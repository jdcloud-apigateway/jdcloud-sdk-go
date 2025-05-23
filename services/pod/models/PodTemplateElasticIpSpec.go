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


type PodTemplateElasticIpSpec struct {

    /* 弹性公网IP的限速（单位：MB）  */
    BandwidthMbps int `json:"bandwidthMbps"`

    /* IP服务商，取值为bgp或no_bgp，默认：bgp (Optional) */
    Provider *string `json:"provider"`

    /* 弹性公网IP计费模式。可选值：
`bandwidth`：按带宽计费。
`flow`：按流量计费。
若指定`chargeMode=bandwidth`则弹性公网IP计费类型同实例（包年包月或按配置）。  */
    ChargeMode string `json:"chargeMode"`
}
