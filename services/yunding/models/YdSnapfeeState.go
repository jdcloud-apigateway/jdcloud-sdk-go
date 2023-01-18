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


type YdSnapfeeState struct {

    /* 云鼎pin (Optional) */
    YdPin string `json:"ydPin"`

    /* 云鼎资源id（产研侧） (Optional) */
    ResourceId string `json:"resourceId"`

    /* 状态 1：正常，3：欠费停服，4：欠费删除，13：到期停服/退款删除，14：到期删除，24：运营删除，34:用户删除 (Optional) */
    State int `json:"state"`
}