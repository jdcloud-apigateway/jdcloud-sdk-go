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


type ModifySecurityGroupRules struct {

    /* 安全组规则的ID。  */
    RuleId string `json:"ruleId"`

    /* 规则限定协议。300:All; 6:TCP; 17:UDP; 1:ICMP (Optional) */
    Protocol *int `json:"protocol"`

    /* 访问控制策略：allow:允许，deny：拒绝 (Optional) */
    RuleAction *string `json:"ruleAction"`

    /* 规则匹配优先级，取值范围为[1,100]，优先级数字越小优先级越高 (Optional) */
    Priority *int `json:"priority"`

    /* 安全组规则的起始端口。取值范围：1-65535 (Optional) */
    FromPort *int `json:"fromPort"`

    /* 安全组规则的终端口。取值范围：1-65535 (Optional) */
    ToPort *int `json:"toPort"`

    /* 安全组规则前缀，取值范围：正确的CIDR (Optional) */
    AddressPrefix *string `json:"addressPrefix"`

    /* 安全组规则的描述，取值范围：0-256个UTF-8编码下的全部字符 (Optional) */
    Description *string `json:"description"`
}
