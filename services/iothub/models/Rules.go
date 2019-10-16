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


type Rules struct {

    /* 规则编号 (Optional) */
    RuleId string `json:"ruleId"`

    /* 0-正常规则，1-异常规则 (Optional) */
    RuleType int `json:"ruleType"`

    /* 用户填写的规则信息 (Optional) */
    RuleInfo string `json:"ruleInfo"`

    /* 用户规则映射的jcq信息 (Optional) */
    JcqInfo string `json:"jcqInfo"`
}