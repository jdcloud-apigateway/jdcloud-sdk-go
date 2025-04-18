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


type GpdpSchedTriggerDependTree struct {

    /*  (Optional) */
    Id int `json:"id"`

    /*  (Optional) */
    JobName string `json:"jobName"`

    /*  (Optional) */
    DependName string `json:"dependName"`

    /* 作业所属系统，G gravity版本、A automation版本,废弃字段 (Optional) */
    JobBelong string `json:"jobBelong"`

    /* 是否可以上下游并行，1启用、0关闭 (Optional) */
    FlagParallel string `json:"flagParallel"`

    /* 是否强依赖，1启用、0关闭 (Optional) */
    Necessary string `json:"necessary"`

    /* 依赖偏移量，0当天，-1昨天 (Optional) */
    DayOffset string `json:"dayOffset"`

    /* 依赖作业所属系统，G gravity版本、A automation版本,废弃字段 (Optional) */
    DependJobBelong string `json:"dependJobBelong"`

    /* 是否启用，1启用、0关闭 (Optional) */
    Enable string `json:"enable"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 更新时间 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 客户作业名称 (Optional) */
    CstJobName string `json:"cstJobName"`

    /* 最后状态 (Optional) */
    LastStatus string `json:"lastStatus"`

    /* 最后数据日期 (Optional) */
    LastTxDate string `json:"lastTxDate"`

    /* 计划执行时间 (Optional) */
    PlanExeTime string `json:"planExeTime"`

    /* 作业表名 (Optional) */
    TableName string `json:"tableName"`

    /* manager作业负责人 (Optional) */
    Manager string `json:"manager"`

    /* 作业描述 (Optional) */
    JobDesc string `json:"jobDesc"`

    /* 显示信息 (Optional) */
    Display string `json:"display"`

    /* 最后开始时间 (Optional) */
    LastStartTime string `json:"lastStartTime"`

    /* 最后结束时间 (Optional) */
    LastEndTime string `json:"lastEndTime"`

    /* 是否含有父类 (Optional) */
    Parent bool `json:"parent"`

    /* 是否有子类 (Optional) */
    Child bool `json:"child"`

    /* 依赖父作业清单 (Optional) */
    DependList []GpdpSchedTriggerDependTree `json:"dependList"`
}
