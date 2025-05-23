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


type GpjsSchedQueuePendingDto struct {

    /* 作业名称 (Optional) */
    JobName string `json:"jobName"`

    /* 数据日期 (Optional) */
    TxDate string `json:"txDate"`

    /* sessionid (Optional) */
    Sessionid int `json:"sessionid"`

    /* 作业优先级 (Optional) */
    Priority int `json:"priority"`

    /* 超时时间 (Optional) */
    Timeout int `json:"timeout"`

    /* 进入队列的类型 (Optional) */
    EqType string `json:"eqType"`

    /* 命名空间，已经废弃 (Optional) */
    NsName string `json:"nsName"`

    /* 服务器资源消耗，已经废弃 (Optional) */
    ServerResourVal int `json:"serverResourVal"`

    /* 运行参数 (Optional) */
    RunParams string `json:"runParams"`

    /* 调度时间 (Optional) */
    SchedTime string `json:"schedTime"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 更新时间 (Optional) */
    UpdateTime string `json:"updateTime"`
}
