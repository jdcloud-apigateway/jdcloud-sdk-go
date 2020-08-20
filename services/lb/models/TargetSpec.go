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


type TargetSpec struct {

    /* Target所属实例的Id，只有type为vm或container时此项才需要 (Optional) */
    InstanceId *string `json:"instanceId"`

    /* Target所属的type，取值为vm、container或ip，默认为vm。vm和container类型对应服务器组的instance类型，ip类型对应服务器组的ip类型。 (Optional) */
    Type *string `json:"type"`

    /* Target提供服务的端口，取值范围：0-65535，其中0表示与backend的端口相同，默认为0。 <br>【dnlb】使用限制：dnlb同一TargetGroup下，同一实例/ip仅允许一个端口提供服务 (Optional) */
    Port *int `json:"port"`

    /* 该Target的权重，取值范围：0-100 ，默认为10。0表示不参与流量转发，仅alb支持权重为0的target (Optional) */
    Weight *int `json:"weight"`

    /* Target的ip地址。仅当type为ip时，此项必须配置。 (Optional) */
    IpAddress *string `json:"ipAddress"`
}
