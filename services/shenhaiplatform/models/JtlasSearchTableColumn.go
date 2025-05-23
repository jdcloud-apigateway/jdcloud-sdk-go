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


type JtlasSearchTableColumn struct {

    /* 字段id (Optional) */
    Id string `json:"id"`

    /* 字段名称 (Optional) */
    Name string `json:"name"`

    /* 字段描述 (Optional) */
    Comment string `json:"comment"`

    /* 字段类型 (Optional) */
    Type string `json:"type"`

    /* 是否是分区字段 (Optional) */
    IsPartition bool `json:"isPartition"`

    /* 字段长度 (Optional) */
    Length int `json:"length"`

    /* 字段位置 (Optional) */
    Position int `json:"position"`
}
