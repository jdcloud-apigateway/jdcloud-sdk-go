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


type DBInstanceInternal struct {

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 实例状态.RUNNING：运行, ERROR：错误 ,BUILDING：创建中, DELETING：删除中, RESTORING：恢复中, RESIZING：变配中 (Optional) */
    InstanceStatus string `json:"instanceStatus"`
}