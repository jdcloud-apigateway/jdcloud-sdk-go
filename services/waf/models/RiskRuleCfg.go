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


type RiskRuleCfg struct {

    /* 规则id (Optional) */
    Id int `json:"id"`

    /* WAF实例id (Optional) */
    WafInstanceId string `json:"wafInstanceId"`

    /* 域名 (Optional) */
    Domain string `json:"domain"`

    /* 规则名称 (Optional) */
    Name string `json:"name"`

    /* uri 以/开头 (Optional) */
    Uri string `json:"uri"`

    /* 动作 支持notice / verify@captcha (Optional) */
    Action string `json:"action"`

    /* 0-使用中 1-禁用 (Optional) */
    Disable int `json:"disable"`

    /* 更新时间，s (Optional) */
    UpdateTime int `json:"updateTime"`
}
