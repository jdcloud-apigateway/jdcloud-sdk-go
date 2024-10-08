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


type RuleSet struct {

    /* 规则集的标识。 (Optional) */
    Id string `json:"id"`

    /* 规则集的名称，例如default。 (Optional) */
    Name string `json:"name"`

    /* 规则集的描述。 (Optional) */
    Description string `json:"description"`

    /* 规则集的类型，有效值zone。 (Optional) */
    Kind string `json:"kind"`

    /* 执行规则集的阶段，有效值http_ratelimit/http_request_late_transform/http_request_firewall_custom/http_request_firewall_managed。 (Optional) */
    Phase string `json:"phase"`

    /* 规则集最近修改时间。 (Optional) */
    Last_updated string `json:"last_updated"`

    /* 规则集中的所有规则。 (Optional) */
    Rules []Rule `json:"rules"`

    /* 规则集的版本。 (Optional) */
    Version string `json:"version"`

    /*  (Optional) */
    Source string `json:"source"`
}
