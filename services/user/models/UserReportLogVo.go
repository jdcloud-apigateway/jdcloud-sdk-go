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


type UserReportLogVo struct {

    /* id (Optional) */
    Id int `json:"id"`

    /* pin (Optional) */
    Pin string `json:"pin"`

    /* 历史版本 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 操作人 (Optional) */
    Operator string `json:"operator"`

    /* 备注 (Optional) */
    Remark string `json:"remark"`
}