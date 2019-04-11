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


type Userview struct {

    /* 自定义线路ID (Optional) */
    ViewId int `json:"viewId"`

    /* 自定义线路名称, 最多64个字符 (Optional) */
    ViewName string `json:"viewName"`

    /* 主域名ID (Optional) */
    DomainId int `json:"domainId"`

    /* 域名 (Optional) */
    DomainName string `json:"domainName"`

    /* 是否删除，0:没有删除，1:已删除 (Optional) */
    IsDelete int `json:"isDelete"`

    /* 创建者 (Optional) */
    Creator string `json:"creator"`

    /* 创建时间，格式Unix timestamp，时间单位：秒 (Optional) */
    CreateTime int `json:"createTime"`

    /* 更新者 (Optional) */
    Updator string `json:"updator"`

    /* 更新时间，格式Unix timestamp，时间单位：秒 (Optional) */
    UpdateTime int `json:"updateTime"`
}
