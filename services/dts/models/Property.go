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


type Property struct {

    /* 组件类型：el-select、el-input、el-radio-group、el-checkbox-group (Optional) */
    ComponentType *string `json:"componentType"`

    /* 参数名称 (Optional) */
    Variable *string `json:"variable"`

    /* 默认值 (Optional) */
    DefaultValue *interface{} `json:"defaultValue"`

    /* 数据类型：int、string、array、object (Optional) */
    DataType *string `json:"dataType"`

    /* 单位 (Optional) */
    Unit *string `json:"unit"`

    /* 标签名称 (Optional) */
    Label *string `json:"label"`

    /* 该参数的可选值 (Optional) */
    Options []Option `json:"options"`

    /* 该参数描述 (Optional) */
    Description *string `json:"description"`

    /* 该参数的可见性 (Optional) */
    Visibility []Visibility `json:"visibility"`

    /* 其他配置参数 (Optional) */
    AdditionalProperties *interface{} `json:"additionalProperties"`
}
