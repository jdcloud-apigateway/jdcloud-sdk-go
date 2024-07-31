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


type EditPersonsInGroupReqVo struct {

    /* 联系组ID (Optional) */
    GroupId int64 `json:"groupId"`

    /* 联系人ID（可多个，英文逗号隔开） (Optional) */
    PersonIds string `json:"personIds"`

    /* 是否是账号联系人：1-账号联系人 2-非账号联系人（可多个，英文逗号隔开，与personIds相对应） (Optional) */
    IsSelfs string `json:"isSelfs"`
}
