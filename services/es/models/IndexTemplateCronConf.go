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


type IndexTemplateCronConf struct {

    /* 索引模板名称  */
    TemplateName string `json:"templateName"`

    /* 任务起始执行时间，遵循ISO8601标准，使用UTC时间，格式为：YYYY-MM-DDTHH:mm:ssZ  */
    BeginTime string `json:"beginTime"`

    /* 任务执行频率, day： 每天，week： 每周，month：每月  */
    Cycle string `json:"cycle"`

    /* 索引前缀  */
    IndexPrefix string `json:"indexPrefix"`

    /* 索引后缀格式, "yyyy-MM-dd", "yyyy.MM.dd", "yyyy_MM_dd", "yyyyMMdd", "yyyyww", "yyyy-MM", "yyyy.MM", "yyyy_MM", "yyyyMM"  */
    IndexSuffixFormat string `json:"indexSuffixFormat"`

    /* 提前创建索引天数  */
    AheadOfDay int `json:"aheadOfDay"`

    /* 是否开启自动删除索引  */
    IsAutoDelete bool `json:"isAutoDelete"`

    /* 索引保留天数, 开启自动删除索引后生效  */
    ReserveOfDay int `json:"reserveOfDay"`
}