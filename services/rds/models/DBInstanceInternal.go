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


type DBInstanceInternal struct {

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称，具体规则可参见帮助中心文档:[名称及密码限制](../../../documentation/Database-and-Cache-Service/RDS/Introduction/Restrictions/SQLServer-Restrictions.md) (Optional) */
    InstanceName string `json:"instanceName"`

    /* 实例引擎类型，如MySQL或SQL Server等，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    Engine string `json:"engine"`

    /* 实例类别，例如主实例，只读实例等，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    InstanceType string `json:"instanceType"`

    /* 实例状态，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    InstanceStatus string `json:"instanceStatus"`

    /* 实例的网络状态；0表示实例的网络处于中断状态，1表示实例的网络处于正常状态 (Optional) */
    AccessibleStatus int `json:"accessibleStatus"`

    /* 物理机IP (Optional) */
    HostIp string `json:"hostIp"`

    /* 实例ip (Optional) */
    UserIp string `json:"userIp"`

    /* 云主机ID (Optional) */
    VmId string `json:"vmId"`

    /* 实例内网域名 (Optional) */
    InternalDomainName string `json:"internalDomainName"`

    /* 应用访问端口 (Optional) */
    InstancePort string `json:"instancePort"`
}
