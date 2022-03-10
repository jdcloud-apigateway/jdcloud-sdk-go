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


type AgentAttributes struct {

    /* agent Id (Optional) */
    AgentId string `json:"agentId"`

    /* 主机名称 (Optional) */
    HostName string `json:"hostName"`

    /* 主机IP (Optional) */
    Ip string `json:"ip"`

    /* agent 的状态 (Optional) */
    AgentStatus string `json:"agentStatus"`

    /* 该Agent管理的数据源的具体信息，同一个Agent可以支持多种不同的数据源类型 (Optional) */
    DataSource []DataSourceEntry `json:"dataSource"`
}
