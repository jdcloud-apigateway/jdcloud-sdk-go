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


type Message struct {

    /* 消息id (Optional) */
    MsgId string `json:"msgId"`

    /* topic名称 (Optional) */
    Topic string `json:"topic"`

    /* 消息体 (Optional) */
    MessageBody string `json:"messageBody"`

    /* 消息tag (Optional) */
    Tags string `json:"tags"`

    /* 消息key (Optional) */
    Key string `json:"key"`

    /* 消息的存储时间 (Optional) */
    StoreTime string `json:"storeTime"`

    /* 消息的生产时间 (Optional) */
    BornTime string `json:"bornTime"`
}
