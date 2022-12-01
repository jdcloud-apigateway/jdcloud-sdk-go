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


type DmsOnlineTask struct {

    /* 主键id。 (Optional) */
    Id int `json:"id"`

    /* 数据源id。 (Optional) */
    DataSourceId int `json:"dataSourceId"`

    /* 数据库名称。 (Optional) */
    DbName string `json:"dbName"`

    /* 任务类型，DDL， DML。 (Optional) */
    TaskType string `json:"taskType"`

    /* 计划时间。 (Optional) */
    PlanTime string `json:"planTime"`

    /* 是否为事务。 (Optional) */
    Transaction bool `json:"transaction"`

    /* 是否为并行执行。 (Optional) */
    Parallel bool `json:"parallel"`

    /* 是否忽略错误。 (Optional) */
    IgnoreError bool `json:"ignoreError"`

    /* 运行状态。CREATE("CREATE", 1),RUNNING("RUNNING", 2), SUCCESS("SUCCESS", 3), FAILED("FAILED", 4), SUSPEND("SUSPEND", 5); (Optional) */
    RunStatus string `json:"runStatus"`

    /* 创建日期。 (Optional) */
    CreatedDate string `json:"createdDate"`

    /* 结束日期。 (Optional) */
    FinishDate string `json:"finishDate"`

    /* 创建用户。 (Optional) */
    CreateUser string `json:"createUser"`

    /* 状态0有效，1无效。 (Optional) */
    Status string `json:"status"`

    /* 执行消息。 (Optional) */
    Message string `json:"message"`

    /* 运行ip。 (Optional) */
    RunIp string `json:"runIp"`
}
