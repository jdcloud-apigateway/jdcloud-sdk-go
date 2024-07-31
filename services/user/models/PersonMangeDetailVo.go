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


type PersonMangeDetailVo struct {

    /* pin (Optional) */
    Pin string `json:"pin"`

    /* 个人姓名 (Optional) */
    Name string `json:"name"`

    /* 身份证号 (Optional) */
    Cnumber string `json:"cnumber"`

    /* 新老实名 (Optional) */
    OldOrNewUser int `json:"oldOrNewUser"`

    /* 认证时间 (Optional) */
    AuthTime string `json:"authTime"`

    /* 认证方式 (Optional) */
    AuthType int `json:"authType"`

    /* 认证状态 (Optional) */
    PersonStatus int `json:"personStatus"`

    /* 认证渠道 (Optional) */
    AuthChannel string `json:"authChannel"`
}
