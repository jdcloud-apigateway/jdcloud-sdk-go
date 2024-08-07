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


type OaDetail struct {

    /* 任务名称 (Optional) */
    TaskName string `json:"taskName"`

    /* 审批人 (Optional) */
    SubmitRealName string `json:"submitRealName"`

    /* 任务状态 (Optional) */
    SubmitResult string `json:"submitResult"`

    /* 审批意见 (Optional) */
    SubmitComments string `json:"submitComments"`

    /* 任务开始时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 任务结束时间 (Optional) */
    EndTime string `json:"endTime"`
}
