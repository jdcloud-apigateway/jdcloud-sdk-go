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


type AuditRange struct {

    /* 正常数据是否入审，true-入审，false-不入审，缺省为 false (Optional) */
    Pass bool `json:"pass"`

    /* 疑似数据是否入审，true-入审，false-不入审，缺省为 false (Optional) */
    Review bool `json:"review"`

    /* 确认违规数据是否入审，true-入审，false-不入审，缺省为 false (Optional) */
    Block bool `json:"block"`
}