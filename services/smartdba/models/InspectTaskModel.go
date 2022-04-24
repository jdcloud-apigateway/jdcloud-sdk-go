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


type InspectTaskModel struct {

    /* 诊断id (Optional) */
    InspectId string `json:"inspectId"`

    /* rds实例id (Optional) */
    InstanceId string `json:"instanceId"`

    /* rds实例名称 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 实例规格 (Optional) */
    InstanceClass string `json:"instanceClass"`

    /* 区域 (Optional) */
    Region string `json:"region"`

    /* 实例类型 (Optional) */
    InstanceType string `json:"instanceType"`

    /* 数据库版本 (Optional) */
    EngineVersion string `json:"engineVersion"`

    /* 巡检分值 (Optional) */
    Score int `json:"score"`

    /* 巡检范围的起始时间，如：2020-11-09T00:00:00Z (Optional) */
    BeginTime string `json:"beginTime"`

    /* 巡检范围的截至时间，如：2020-11-09T23:59:59Z (Optional) */
    EndTime string `json:"endTime"`

    /* 巡检开始时间，2021-11-09T00:19:00Z (Optional) */
    CreateTime string `json:"createTime"`

    /* 巡检完成时间，2021-11-09T00:19:30Z (Optional) */
    UpdateTime string `json:"updateTime"`

    /* cpu使用率 (Optional) */
    CupInfo string `json:"cupInfo"`

    /* 内存使用率 (Optional) */
    MemoryInfo string `json:"memoryInfo"`

    /* 连接使用率 (Optional) */
    ConnectInfo string `json:"connectInfo"`

    /* 慢SQL数量 (Optional) */
    SlowSqlInfo string `json:"slowSqlInfo"`

    /* 大表数量 (Optional) */
    BigTableInfo string `json:"bigTableInfo"`

    /* 是否死锁 (Optional) */
    DeadLockInfo string `json:"deadLockInfo"`
}