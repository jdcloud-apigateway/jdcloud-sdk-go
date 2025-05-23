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


type GpjmListJobInstanceReq struct {

    /* 页面大小  */
    PageSize int `json:"pageSize"`

    /* 页码  */
    PageNum int `json:"pageNum"`

    /* 运行状态 (Optional) */
    StatusList []string `json:"statusList"`

    /* 作业上上线下线状态 (Optional) */
    Enable string `json:"enable"`

    /* 运行周期 (Optional) */
    Cycle string `json:"cycle"`

    /* 处理类型 (Optional) */
    ProcessType string `json:"processType"`

    /* 负责人 (Optional) */
    Manager string `json:"manager"`

    /* 作业名 (Optional) */
    JobName string `json:"jobName"`

    /* 任务类型 (Optional) */
    JobType string `json:"jobType"`

    /* 任务创建日期-结束 (Optional) */
    CreateTimeBefore string `json:"createTimeBefore"`

    /* 任务创建日期-开始 (Optional) */
    CreateTimeAfter string `json:"createTimeAfter"`

    /* 运行开始时间范围 (Optional) */
    StartTimeBefore string `json:"startTimeBefore"`

    /* 运行开始时间范围 (Optional) */
    StartTimeAfter string `json:"startTimeAfter"`

    /* 运行结束时间范围 (Optional) */
    EndTimeAfter string `json:"endTimeAfter"`

    /* 运行结束时间范围 (Optional) */
    EndTimeBefore string `json:"endTimeBefore"`

    /* 数据日期 (Optional) */
    TxDate string `json:"txDate"`

    /*  (Optional) */
    JobChildType string `json:"jobChildType"`

    /* 客户作业名 (Optional) */
    CstJobName string `json:"cstJobName"`
}
