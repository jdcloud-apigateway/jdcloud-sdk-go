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


type ProtocolProp struct {

    /* 属性名称 (Optional) */
    PropName string `json:"propName"`

    /* 属性描述 (Optional) */
    PropDesc string `json:"propDesc"`

    /* 属性值长度限制 (Optional) */
    InputLength string `json:"inputLength"`

    /* 属性默认值 (Optional) */
    InputDefault string `json:"inputDefault"`

    /* 是否必填,0-非唯一，1-唯一 (Optional) */
    Unique int `json:"unique"`

    /* 是否必填,0-非必填，1-必填 (Optional) */
    Required int `json:"required"`
}