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


type GpdjmcLoopParam struct {

    /* 循环类型 (Optional) */
    LoopType *string `json:"loopType"`

    /* 循环间隔 (Optional) */
    LoopInterval *int `json:"loopInterval"`

    /* 循环次数（仅迭代类型需要） (Optional) */
    LoopTimes *int `json:"loopTimes"`

    /* 批次大小 (Optional) */
    BatchSize *int `json:"batchSize"`

    /* 循环列表参数（仅列表类型需要） (Optional) */
    LoopListParams []GpdjmcLoopListParam `json:"loopListParams"`
}
