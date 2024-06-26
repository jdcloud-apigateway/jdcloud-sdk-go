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


type UnhealthReason struct {

    /* CONNECT_FAILED:LB与后端server建立连接失败（HTTP/TCP）;CONNECT_TIMEOUT: LB与后端server建立连接超时（HTTP/TCP）;CONNECT_INTERRUPT: LB与后端服务器连接中断（TCP/HTTP）; RECEIVE_RESPONSE_FAILED: LB接受后端服务器响应失败（HTTP/TCP/ICMP）；RECEIVE_RESPONSE_TIMEOUT: LB接受后端服务器响应超时（HTTP/TCP/ICMP）；REQUEST_FAILED: LB向后端server发送请求失败（HTTP/TCP/ICMP）；REQUEST_TIMEOUT: LB向后端server发送请求超时（HTTP/TCP/ICMP）;HTTP_CODE_MISMATCH: LB从后端收到的响应码与预期正常的不一致（HTTP）;RESPONSE_FORMAT_ERROR: LB从后端接收的响应格式错误（HTTP/TCP） (Optional) */
    Reason string `json:"reason"`

    /* 原因描述 (Optional) */
    Detail string `json:"detail"`
}
