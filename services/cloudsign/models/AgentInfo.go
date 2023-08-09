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


type AgentInfo struct {

    /* 授权信息ID (Optional) */
    Id string `json:"id"`

    /* 被授权人 (Optional) */
    AgentName string `json:"agentName"`

    /* 被授权人手机号 (Optional) */
    AgentPhone string `json:"agentPhone"`

    /* 被授权人身份证号 (Optional) */
    AgentIdCardNum string `json:"agentIdCardNum"`

    /* 授权书名称 (Optional) */
    AgentFileName string `json:"agentFileName"`

    /* 授权时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 备注 (Optional) */
    Note string `json:"note"`
}