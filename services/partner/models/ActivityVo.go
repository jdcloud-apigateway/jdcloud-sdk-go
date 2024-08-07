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


type ActivityVo struct {

    /* 活动id (Optional) */
    Id int `json:"id"`

    /* 活动uuid (Optional) */
    Uuid string `json:"uuid"`

    /* 活动名称 (Optional) */
    ActivityName string `json:"activityName"`

    /* 活动日期 (Optional) */
    ActivityTime string `json:"activityTime"`

    /* 活动地点 (Optional) */
    Address string `json:"address"`

    /* 活动规模 (Optional) */
    Scale int `json:"scale"`

    /* 活动形式 (Optional) */
    Form int `json:"form"`

    /* 活动形式 (Optional) */
    FormStr string `json:"formStr"`

    /* 线索数 (Optional) */
    ClueNum int `json:"clueNum"`
}
