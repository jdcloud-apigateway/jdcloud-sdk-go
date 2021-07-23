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


type AccountingRule struct {

    /* 站点  */
    Site int `json:"site"`

    /* 产品线 (Optional) */
    AppCode string `json:"appCode"`

    /* 出账对象类型 1：通用 2：用户  */
    TargetType int `json:"targetType"`

    /* 产品 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* pin (Optional) */
    Pin string `json:"pin"`

    /* 出账类型  1：实时出账 2：定期出账  */
    OutAccountType int `json:"outAccountType"`

    /* 出账周期：限制范围 1-28 (Optional) */
    OutAccountDay int `json:"outAccountDay"`

    /* 定期出账 时间表达式 (Optional) */
    TimeCron string `json:"timeCron"`
}
