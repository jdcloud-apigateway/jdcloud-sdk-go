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


type ConsumerInfoDetail struct {

    /* 消费者IP地址 (Optional) */
    ConsumerIp string `json:"consumerIp"`

    /* 消费时间戳(millionSecond) (Optional) */
    TimeStamp int `json:"timeStamp"`

    /* 消费耗时(second) (Optional) */
    CostTime int `json:"costTime"`

    /* 第几次消费 (Optional) */
    ConsumeTimes int `json:"consumeTimes"`

    /* 消费状态[SUCCESS,FAILED_WITHOUT_RESULT,FAILED_WITH_RESULT] (Optional) */
    ConsumerStatus string `json:"consumerStatus"`
}
