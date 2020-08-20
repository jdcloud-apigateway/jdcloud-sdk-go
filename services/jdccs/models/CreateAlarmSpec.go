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


type CreateAlarmSpec struct {

    /* idc机房实例ID (Optional) */
    Idc string `json:"idc"`

    /* 资源类型，bandwidth:带宽 (Optional) */
    ResourceType string `json:"resourceType"`

    /* 带宽实例ID (Optional) */
    ResourceId string `json:"resourceId"`

    /* 规则名称 (Optional) */
    Name string `json:"name"`

    /* 监控项，bandwidthTrafficIn:上行实时流量 bandwidthTrafficOut:下行实时流量 (Optional) */
    Metric string `json:"metric"`

    /* 统计周期（单位：分钟） (Optional) */
    Period int `json:"period"`

    /* 统计方法：平均值=avg、最大值=max、最小值=min (Optional) */
    StatisticMethod string `json:"statisticMethod"`

    /* 计算方式 >=、>、<、<=、=、！= (Optional) */
    Operator string `json:"operator"`

    /* 阈值 (Optional) */
    Threshold float64 `json:"threshold"`

    /* 连续多少次后报警 (Optional) */
    Times int `json:"times"`

    /* 通知周期 单位：小时 (Optional) */
    NoticePeriod int `json:"noticePeriod"`

    /* 规则状态 disabled:禁用 enabled:启用 (Optional) */
    Status string `json:"status"`

    /* 通知方式 all:全部 sms：短信 email:邮件 (Optional) */
    NoticeMethod string `json:"noticeMethod"`

    /* 通知对象 all:全部 persons：个人 groups:角色组 (Optional) */
    NoticeObj string `json:"noticeObj"`

    /* 通知对象用户ID,若多个用逗号分隔 (Optional) */
    UserId string `json:"userId"`
}
