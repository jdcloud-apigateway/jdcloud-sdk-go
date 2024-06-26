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


type TargetHealth struct {

    /* Target所在的虚拟服务器组Id, 与agId不能并存 (Optional) */
    TargetGroupId string `json:"targetGroupId"`

    /* Target所在的高可用组Id，与targetGroupId不能并存 (Optional) */
    AgId string `json:"agId"`

    /* Target所属实例的Id（type为vm或container时显示） (Optional) */
    InstanceId string `json:"instanceId"`

    /* Target所属的type，取值为vm、container或ip,默认为vm (Optional) */
    Type string `json:"type"`

    /* 健康检查的port (Optional) */
    Port int `json:"port"`

    /* 该Target的权重，取值范围：0-100 ，默认为10。0表示不参与流量转发 (Optional) */
    Weight int `json:"weight"`

    /* 该Target的健康状态，取值为healthy、unhealthy (Optional) */
    Status string `json:"status"`

    /* Target的IP地址。当Target type为vm或container时，表示vm或container的私网IP；当Target type为ip时，表示注册Target时指定的IP地址 (Optional) */
    IpAddress string `json:"ipAddress"`

    /* 健康异常结构 (Optional) */
    UnhealthReason UnhealthReason `json:"unhealthReason"`
}
