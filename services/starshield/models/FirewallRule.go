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


type FirewallRule struct {

    /*  (Optional) */
    Products []string `json:"products"`

    /* 规则的优先级，允许控制处理顺序。一个较小的数字表示高优先级。如果不提供，任何有优先权的规则将在没有优先权的规则之前排序。 (Optional) */
    Priority int `json:"priority"`

    /* 此防火墙规则当前是否已暂停。 (Optional) */
    Paused bool `json:"paused"`

    /* 短引用标记，用于快速选择相关规则。 (Optional) */
    Ref string `json:"ref"`

    /*  (Optional) */
    Action_parameters Action_parameters `json:"action_parameters"`

    /* 应用于匹配请求的行动。注意，行动 "log "只适用于企业客户。 (Optional) */
    Action string `json:"action"`

    /*  (Optional) */
    Filter Filter `json:"filter"`

    /* 防火墙规则标识符 (Optional) */
    Id string `json:"id"`

    /* 对规则的描述，以帮助识别它。 (Optional) */
    Description string `json:"description"`
}
