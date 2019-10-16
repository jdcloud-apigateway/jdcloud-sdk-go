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


type StatusReportResp struct {

    /* 手机号 (Optional) */
    PhoneNum string `json:"phoneNum"`

    /* 发送短信的序列号 (Optional) */
    SequenceNumber string `json:"sequenceNumber"`

    /* 短信发送时间（yyyy-MM-dd HH:mm:ss) (Optional) */
    SendTime string `json:"sendTime"`

    /* 接收到回执的时间（yyyy-MM-dd HH:mm:ss) (Optional) */
    ReportTime string `json:"reportTime"`

    /* 发送状态 (Optional) */
    Status int `json:"status"`

    /* 错误码 (Optional) */
    Code string `json:"code"`

    /* 长短信拆分序号（短短信直接返回1) (Optional) */
    SplitNum int `json:"splitNum"`
}