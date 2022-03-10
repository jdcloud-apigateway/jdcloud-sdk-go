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


type NowSession struct {

    /* 会话id (Optional) */
    Id int `json:"id"`

    /* 会话用户 (Optional) */
    User string `json:"user"`

    /* 会话源端IP (Optional) */
    Host string `json:"host"`

    /* 数据库 (Optional) */
    Db string `json:"db"`

    /* session命令 (Optional) */
    Command string `json:"command"`

    /* 会话活跃时间 (Optional) */
    Time int `json:"time"`

    /* 会话状态 (Optional) */
    State string `json:"state"`

    /* 正在执行的sql (Optional) */
    Info string `json:"info"`
}
