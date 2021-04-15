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


type JRTCAuthInfo struct {

    /* appId (Optional) */
    AppId string `json:"appId"`

    /* appKey (Optional) */
    AppKey string `json:"appKey"`

    /* 用户id (Optional) */
    UserId string `json:"userId"`

    /* 会议号 (Optional) */
    RoomId int64 `json:"roomId"`

    /* 随机令牌 (Optional) */
    Nonce string `json:"nonce"`

    /* 时间戳-毫秒 (Optional) */
    Timestamp int64 `json:"timestamp"`

    /* token (Optional) */
    Token string `json:"token"`

    /* 是否可用（true-可用,false-不可用） (Optional) */
    Available bool `json:"available"`
}