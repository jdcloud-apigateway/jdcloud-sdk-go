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


type NetworkInterfaceSpec struct {

    /* 子网ID  */
    SubnetId string `json:"subnetId"`

    /* 可用区，用户的默认可用区，该参数无效，不建议使用 (Optional) */
    Az *string `json:"az"`

    /* 网卡名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。 (Optional) */
    NetworkInterfaceName *string `json:"networkInterfaceName"`

    /* 网卡主IP，如果不指定，会自动从子网中分配 (Optional) */
    PrimaryIpAddress *string `json:"primaryIpAddress"`

    /* SecondaryIp列表 (Optional) */
    SecondaryIpAddresses []string `json:"secondaryIpAddresses"`

    /* 自动分配的SecondaryIp数量 (Optional) */
    SecondaryIpCount *int `json:"secondaryIpCount"`

    /* 要绑定的安全组ID列表，最多指定5个安全组 (Optional) */
    SecurityGroups []string `json:"securityGroups"`

    /* 源和目标IP地址校验，取值为0或者1,默认为1 (Optional) */
    SanityCheck *int `json:"sanityCheck"`

    /* 描述,​ 允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description *string `json:"description"`

    /* 自动分配的ipv6地址数量 (Optional) */
    Ipv6AddressCount *int `json:"ipv6AddressCount"`

    /* 指定分配的ipv6地址，不能与ipv6AddressCount同时指定 (Optional) */
    Ipv6Address *string `json:"ipv6Address"`
}
