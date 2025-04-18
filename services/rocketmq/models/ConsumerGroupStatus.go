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


type ConsumerGroupStatus struct {

    /* 消费组名称 (Optional) */
    Group string `json:"group"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 描述 (Optional) */
    Description string `json:"description"`

    /* 消费组是否开启死信队列 (Optional) */
    EnableDlq bool `json:"enableDlq"`

    /* 消费者消费失败最大重试次数 (Optional) */
    RetryMaxTimes int `json:"retryMaxTimes"`

    /* 消费者是否在线 (Optional) */
    Online bool `json:"online"`

    /* 消费者在线个数 (Optional) */
    Count int `json:"count"`

    /* 消费模式（集群/广播/null，如果返回null前端显示未知） (Optional) */
    MessageModel string `json:"messageModel"`

    /* 消费组总积压数量 (Optional) */
    DiffTotol int `json:"diffTotol"`

    /* 是否开启消费 (Optional) */
    EnableConsume bool `json:"enableConsume"`
}
