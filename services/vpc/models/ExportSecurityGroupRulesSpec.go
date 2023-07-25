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


type ExportSecurityGroupRulesSpec struct {

    /* 导出文件类型, 目前支持excel和csv
导出字段，目前依次是"安全组规则ID", "方向", "优先级", "策略", "协议", "目标端口", "源/目的IP", "备注", "添加时间"  */
    FileType string `json:"fileType"`
}
