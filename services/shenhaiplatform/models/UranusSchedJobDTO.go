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


type UranusSchedJobDTO struct {

    /*  (Optional) */
    JobCode string `json:"jobCode"`

    /*  (Optional) */
    TaskInfo UranusTaskInfoDTO `json:"taskInfo"`

    /*  (Optional) */
    TaskCode string `json:"taskCode"`

    /*  (Optional) */
    FileCode string `json:"fileCode"`

    /*  (Optional) */
    Content string `json:"content"`

    /*  (Optional) */
    Manager string `json:"manager"`

    /*  (Optional) */
    ValidResult int `json:"validResult"`

    /*  (Optional) */
    ValidResultDesc string `json:"validResultDesc"`

    /*  (Optional) */
    PublicChangeType int `json:"publicChangeType"`

    /*  (Optional) */
    PublicChangeTypeDesc string `json:"publicChangeTypeDesc"`

    /*  (Optional) */
    Reason string `json:"reason"`

    /*  (Optional) */
    FileType string `json:"fileType"`

    /*  (Optional) */
    JobName string `json:"jobName"`

    /*  (Optional) */
    TriggerType string `json:"triggerType"`

    /*  (Optional) */
    Cycle string `json:"cycle"`

    /*  (Optional) */
    TaskParamList []UranusTaskParamDTO `json:"taskParamList"`

    /*  (Optional) */
    CreatedTime int64 `json:"createdTime"`

    /*  (Optional) */
    FlowCode string `json:"flowCode"`

    /*  (Optional) */
    LockStatus int `json:"lockStatus"`

    /*  (Optional) */
    LockUser string `json:"lockUser"`

    /*  (Optional) */
    TaskRelease int `json:"taskRelease"`
}
