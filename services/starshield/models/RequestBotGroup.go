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


type RequestBotGroup struct {

    /* 已通过验证的自动程序
运行正常的自动流量。用于支持搜索引擎和其他服务。
 (Optional) */
    VerifiedBot int `json:"verifiedBot"`

    /* 自动
自动程序分数为 1 的流量。可能是不需要的自动流量。
 (Optional) */
    Automated int `json:"automated"`

    /* 可能自动
自动程序分数为 2-29 的流量。可能是自动流量，但可能包含需要的流量。
 (Optional) */
    LikelyAutomated int `json:"likelyAutomated"`

    /* 可能人工
自动程序分数为 30-99 的流量。可能有人向您的服务请求资源。
 (Optional) */
    LikelyHuman int `json:"likelyHuman"`
}