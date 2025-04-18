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


type GpjmListJobReq struct {

    /* 页面大小  */
    PageSize int `json:"pageSize"`

    /* 页码  */
    PageNum int `json:"pageNum"`

    /* 最后一次运行状态 (Optional) */
    LastJobStatus []string `json:"lastJobStatus"`

    /* 处理类型 (Optional) */
    ExcludeJobTypes []string `json:"excludeJobTypes"`

    /*  (Optional) */
    MarketJobsAll string `json:"marketJobsAll"`

    /* 负责人 (Optional) */
    Manager string `json:"manager"`

    /* 任务类型 (Optional) */
    QueryMode string `json:"queryMode"`

    /* 作业上上线下线状态 (Optional) */
    Enable string `json:"enable"`

    /* 作业名 (Optional) */
    JobName string `json:"jobName"`

    /*  (Optional) */
    CurrentStatus string `json:"currentStatus"`

    /* 处理类型 (Optional) */
    ProcessType string `json:"processType"`

    /* 数据日期 (Optional) */
    LastTxdate string `json:"lastTxdate"`

    /* 运行周期 (Optional) */
    Cycle string `json:"cycle"`

    /* 任务执行时间开始 (Optional) */
    JobStartRunTime string `json:"jobStartRunTime"`

    /* 任务执行时间- 结束 (Optional) */
    JobEndRunTime string `json:"jobEndRunTime"`

    /* 任务创建日期-结束 (Optional) */
    CreateTimeBefore string `json:"createTimeBefore"`

    /* 任务创建日期-开始 (Optional) */
    CreateTimeAfter string `json:"createTimeAfter"`

    /* 任务修改日期-结束 (Optional) */
    UpdateTimeBefore string `json:"updateTimeBefore"`

    /* 任务修改日期-开始 (Optional) */
    UpdateTimeAfter string `json:"updateTimeAfter"`

    /* 任务类型 (Optional) */
    JobType string `json:"jobType"`

    /*  (Optional) */
    MarkClassifyCode string `json:"markClassifyCode"`

    /*  (Optional) */
    MarkName string `json:"markName"`

    /*  (Optional) */
    JobChildType string `json:"jobChildType"`

    /* 客户作业名 (Optional) */
    CstJobName string `json:"cstJobName"`
}
