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


type JobInfoVo struct {

    /* 任务id (Optional) */
    EtlId int `json:"etlId"`

    /* 任务名称 (Optional) */
    CstJobName string `json:"cstJobName"`

    /* 任务编码 (Optional) */
    EtlJob string `json:"etlJob"`

    /* 处理类型 (Optional) */
    ProcessType string `json:"processType"`

    /* 任务执行类型 (Optional) */
    ExeType string `json:"exeType"`

    /* 任务描述 (Optional) */
    Description string `json:"description"`

    /* 周期具体日期 (Optional) */
    Frequency string `json:"frequency"`

    /* 最后一次运行开始时间 (Optional) */
    LastStartTime string `json:"lastStartTime"`

    /* 最后一次运行结束时间 (Optional) */
    LastEndTime string `json:"lastEndTime"`

    /* 运行时间 (Optional) */
    RunTime string `json:"runTime"`

    /* 最后一次运行状态 (Optional) */
    LastJobStatus string `json:"lastJobStatus"`

    /* 最后一次执行日期 (Optional) */
    LastTxDate string `json:"lastTxDate"`

    /* 最后返回编码 (Optional) */
    LastReturnCode int `json:"lastReturnCode"`

    /* 当前任务状态信息 (Optional) */
    CurrentStatusMsg string `json:"currentStatusMsg"`

    /* 是否上线，1-上线、2-下线 (Optional) */
    Enable string `json:"enable"`

    /* 触发类型:dependency 依赖、time 时间、file 文件、manual 手工、once 一次性 (Optional) */
    TriggerType string `json:"triggerType"`

    /* 触发时间 (Optional) */
    TriggerTime string `json:"triggerTime"`

    /* 运行周期 ,F 分钟、H 小时、D 天、W 周、M 月、O 一次性运行、N 无周期 (Optional) */
    Cycle string `json:"cycle"`

    /* 周期具体日期 (Optional) */
    Sequence string `json:"sequence"`

    /* 创建人 (Optional) */
    CreateUser string `json:"createUser"`

    /* 负责人 (Optional) */
    UserName string `json:"userName"`

    /* 重试次数 (Optional) */
    RetryCount int `json:"retryCount"`

    /* 周期开始时间(适用小时分钟) (Optional) */
    SequenceStartTime string `json:"sequenceStartTime"`

    /* 周期结束时间(适用小时分钟) (Optional) */
    SequenceEndTime string `json:"sequenceEndTime"`

    /* 周期间隔(适用小时分钟，当周期为小时，含义为间隔小时数，当周期为分钟，含义为间隔分钟数) (Optional) */
    SequenceInterval int `json:"sequenceInterval"`

    /* 是否重试过 (Optional) */
    FlagRetry string `json:"flagRetry"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 更新时间 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 周期中文名 ,F 分钟、H 小时、D 天、W 周、M 月、O 一次性运行、N 无周期 (Optional) */
    CycleLabel string `json:"cycleLabel"`

    /* 触发方式中文名:dependency 依赖、time 时间、file 文件、manual 手工、once 一次性 (Optional) */
    TriggerTypeLabel string `json:"triggerTypeLabel"`

    /* 处理类型中文名（数据计算，数据抽取，ODS加工，数据推送，数据同步, 质量任务） (Optional) */
    ProcessTypeCn string `json:"processTypeCn"`

    /* 数据来源渠道，集成开发（IDE）、数据管道（PIPE） (Optional) */
    Channel string `json:"channel"`

    /* 任务类型 (Optional) */
    JobType string `json:"jobType"`

    /* 任务子类型 (Optional) */
    JobChildType string `json:"jobChildType"`

    /* 任务优先级，数字越小优先级越高 (Optional) */
    Priority int `json:"priority"`

    /* 任务优先级按照等级分，L1-L4，数字越小优先级越高 (Optional) */
    PriorityLevel string `json:"priorityLevel"`

    /* 租户编码 (Optional) */
    CompanyCode string `json:"companyCode"`

    /* 工作空间编码 (Optional) */
    WorkspaceCode string `json:"workspaceCode"`
}
