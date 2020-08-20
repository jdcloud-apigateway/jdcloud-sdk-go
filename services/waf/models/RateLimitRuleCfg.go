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


type RateLimitRuleCfg struct {

    /* 序号id,更新时需要 (Optional) */
    Id int `json:"id"`

    /* 规则名称  */
    Name string `json:"name"`

    /* 限速host配置 (Optional) */
    Host MatchOpValCfg `json:"host"`

    /* 限速url配置 (Optional) */
    Url MatchOpValCfg `json:"url"`

    /* ip配置 (Optional) */
    Ip string `json:"ip"`

    /* 限速大小 (Optional) */
    Rate int `json:"rate"`

    /* forbidden redirect，缺省为forbidden (Optional) */
    MatchAction string `json:"matchAction"`

    /* 跳转地址，matchAction为redirect时必须为当前实例下的域名的url，forbidden时为自定义页面名称，缺省为default (Optional) */
    Redirection string `json:"redirection"`
}
